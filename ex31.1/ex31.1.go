package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type JsonDate time.Time

type Todo struct {
	ID         int      `json:"id,omitempty"`
	Name       string   `json:"name"`
	Completed  bool     `json:"completed,omitempty"`
	ExpireDate JsonDate `json:"date,omitempty"`
}

func (j *JsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if len(s) == 0 {
		return nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonDate(t)
	return nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

// Maybe a Format function for printing your date
func (j JsonDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

var todoMap map[int]Todo
var lastID int = 0

func MakeWebHandler() http.Handler {
	rd = render.New()
	todoMap = make(map[int]Todo)
	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")
	mux.HandleFunc("/todos", PostTodoHandler).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", RemoveTodoHandler).Methods("DELETE")
	mux.HandleFunc("/todos/{id:[0-9]+}", UpdateTodoHandler).Methods("POST")

	// test data
	todoMap[0] = Todo{ID: 0, Name: "aaa", Completed: false}
	lastID++
	return mux
}

type Todos []Todo

func (t Todos) Len() int {
	return len(t)
}

func (t Todos) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Todos) Less(i, j int) bool {
	return t[i].ID > t[j].ID
}

func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Todos, 0)
	for _, todo := range todoMap {
		list = append(list, todo)
	}

	sort.Sort(list)

	rd.JSON(w, http.StatusOK, list)
}

func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastID++
	todo.ID = lastID
	todoMap[lastID] = todo
	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func RemoveTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if todo, ok := todoMap[id]; ok {
		todo.Name = newTodo.Name
		todo.Completed = newTodo.Completed
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusBadRequest, Success{false})
	}
}

func main() {
	m := MakeWebHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
