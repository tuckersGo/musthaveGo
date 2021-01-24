//ex15/ex15.3/ex15.3.go
package main

import "fmt"

func main() {
	var char rune = '한'

	fmt.Printf("%T\n", char) // ❶ char 타입 출력
	fmt.Println(char)        // ❷ char값 출력
	fmt.Printf("%c\n", char) // ❸ 문자 출력
}
