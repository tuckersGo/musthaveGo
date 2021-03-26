//ch30/ex30.1/ex30.1.go
package main

import (
	"encoding/json"
	"net/http"
	"sort"

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

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
