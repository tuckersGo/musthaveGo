//ch16/ex16.3/ex16.3.go
package main

import (
	"ch16/ex16.3/exinit" // ⓬ exinit 패키지 임포트
	"fmt"
)

func main() { // ⓭ main() 함수
	fmt.Println("main function")
	exinit.PrintD()
}
