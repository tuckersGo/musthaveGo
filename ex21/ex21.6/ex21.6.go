//ex21/ex21.6/ex21.6.go
package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeHello(writer Writer) { // ❷ writer 함수타입 변수 호출
	writer("Hello World")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg) // ❶ 함수 리터럴 외부 변수 f 사용
	})
}
