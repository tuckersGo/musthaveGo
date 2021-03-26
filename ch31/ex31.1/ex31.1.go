//ch31/ex31.1/ex31.1.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type Todo struct { // ❶ 할 일 정보를 담는 Todo 구조체
	ID        int    `json:"id,omitempty"` // ❷ json 포맷으로 변환 옵션
	Name      string `json:"name"`
	Completed bool   `json:"completed,omitempty"`
}

var todoMap map[int]Todo
var lastID int = 0

func MakeWebHandler() http.Handler { // ❸ 웹 서버 핸들러 생성
	rd = render.New()
	todoMap = make(map[int]Todo)
	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")
	mux.HandleFunc("/todos", PostTodoHandler).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", RemoveTodoHandler).Methods("DELETE")
	mux.HandleFunc("/todos/{id:[0-9]+}", UpdateTodoHandler).Methods("PUT")
	return mux
}

type Todos []Todo // ❹ ID로 정렬하는 인터페이스

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
	rd.JSON(w, http.StatusOK, list) // ➎ ID로 정렬하여 전체 목록 반환
}

func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastID++ // ➏ 새로운 ID로 등록하고 만든 Todo 반환
	todo.ID = lastID
	todoMap[lastID] = todo
	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func RemoveTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // ➐ ID에 해당하는 할 일 삭제
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusNotFound, Success{false})
	}
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo // ➑ ID에 해당하는 할 일 수정
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
	n := negroni.Classic() // ➒ negroni 기본 핸들러로 감싼다
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
