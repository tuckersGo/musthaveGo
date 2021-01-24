//ex15/ex15.6/ex15.6.go
package main

import "fmt"

func main() {
	str := "hello 월드"    // ❶ 한글과 영문자가 섞인 문자열
	runes := []rune(str) // ❷ []rune 타입으로 타입 변환

	fmt.Printf("len(str) = %d\n", len(str))     // ❸ string 타입 길이
	fmt.Printf("len(runes) = %d\n", len(runes)) // ➍ []rune 타입 길이
}
