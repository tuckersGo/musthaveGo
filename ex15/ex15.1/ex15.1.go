//ex15/ex15.1/ex15.1.go
package main

import "fmt"

func main() {
	// ❶ 큰따옴표로 묶으면 특수 문자가 동작합니다.
	str1 := "Hello\t'World'\n"

	// ❷ 백쿼트로 묶으면 특수 문자가 동작하지 않습니다.
	str2 := `Go is "awesome"!\nGo is simple and\t'powerful'`
	fmt.Println(str1)
	fmt.Println(str2)
}
