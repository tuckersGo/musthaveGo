//appA/exA6/exA6.go
package main

import (
	"embed"
	"net/http"
)

// jpgs holds the static images used on the home page.
//go:embed static/*
var files embed.FS // ❶ 파일을 추가합니다.

func main() {
	http.Handle("/", http.FileServer(http.FS(files))) // ❷ 파일 서버 실행
	http.ListenAndServe(":3000", nil)
}
