//ex14/ex14.2/ex14.2.go
package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	var p1 *int = &a // ❶ p1은 a의 메모리 공간을 가리킵니다.
	var p2 *int = &a // ❷ p2는 a의 메모리 공간을 가리킵니다.
	var p3 *int = &b // ❸ p3는 b의 메모리 공간을 가리킵니다.

	fmt.Printf("p1 == p2 : %v\n", p1 == p2)
	fmt.Printf("p2 == p3 : %v\n", p2 == p3)
}
