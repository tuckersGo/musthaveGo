//ex15.5/ex15.5.go
package main

import "fmt"

func main() {
	str := "Hello World"

	// ❶ ‘H’, ‘e’, ‘l’, ‘l’, ‘o’, ‘ ‘, ‘W’, ‘o’, ‘r’, ‘l’, ‘d’ 문자코드 배열
	runes := []rune{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64}

	fmt.Println(str)
	fmt.Println(string(runes))
}
