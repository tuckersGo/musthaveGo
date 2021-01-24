//ex29/ex29.7/ex29.7.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World") // ❶ 웹 핸들러 등록
	})

	err := http.ListenAndServeTLS(":3000", "localhost.crt", "localhost.key", nil) // ❷ 웹 서버 시작
	if err != nil {
		log.Fatal(err)
	}
}
