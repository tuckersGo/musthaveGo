//ch29/ex29.6/ex29.6.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHandler() http.Handler { // ❶ 핸들러 인스턴스를 생성하는 함수
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"aaa", 16, 87}
	data, _ := json.Marshal(student)                   // ❷ Student 객체를 []byte로 변환
	w.Header().Add("content-type", "application/json") // ❸ json 포맷임을 표시
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data)) // ❹ 결과 전송
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
