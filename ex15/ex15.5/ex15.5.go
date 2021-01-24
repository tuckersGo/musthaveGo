//ex15/ex15.5/ex15.5.go
package main

import "fmt"

func main() {
	str := "Hello World"

	// ❶ ‘H’, ‘e’, ‘l’, ‘l’, ‘o’, ‘ ‘, ‘W’, ‘o’, ‘r’, ‘l’, ‘d’ 문자코드 배열
	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	fmt.Println(str)
	fmt.Println(string(runes))
}
