//ch29/ex29.5/ex29.5.go
package main

import (
	"fmt"
	"net/http"
)

func MakeWebHandler() http.Handler { // ❶ 핸들러 인스턴스를 생성하는 함수
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})
	return mux
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
