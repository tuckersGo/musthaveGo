//ch30/ex30.3/ex30.3.go
package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student // ❶ 학생 목록을 저장하는 맵
var lastId int

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter() // ❷ gorilla/mux를 만듭니다.
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	//-- ❸ 여기에 새로운 핸들러 등록 --//
	mux.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET")
	mux.HandleFunc("/students", PostStudentHandler).Methods("POST")
	mux.HandleFunc("/students/{id:[0-9]+}", DeleteStudentHandler).Methods("DELETE")

	students = make(map[int]Student) // ❹ 임시 데이터 생성
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2

	return mux
}

type Students []Student // Id로 정렬하는 인터페이스
func (s Students) Len() int {
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0) // ➎ 학생 목록을 Id로 정렬
	for _, student := range students {
		list = append(list, student)
	}

	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list) // ➏ JSON 포맷으로 변경
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // ❶ id를 가져옵니다.
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // ❷ id에 해당하는 학생이 없으면 에러
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student) // ❶ JSON 데이터 변환
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastId++ // ❷ id를 증가시킨후 맵에 등록
	student.Id = lastId
	students[lastId] = student
	w.WriteHeader(http.StatusCreated)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // ❶ id를 가져옵니다.
	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // ❷ id에 해당하는 학생이 없으면 에러
		return
	}
	delete(students, id)
	w.WriteHeader(http.StatusOK) // ❸ StatusOK 반환
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
