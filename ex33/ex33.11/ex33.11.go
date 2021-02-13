//ex33/ex33.11/ex33.11.go
package main

import (
	"embed"
	"net/http"
)

// jpgs holds the static images used on the home page.
//go:embed static/*
var files embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(files))) // ❶
	http.ListenAndServe(":3000", nil)
}
