//ex21/ex21.3/ex21.3.go
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int { // ❶ op에 따른 함수 타입 반환
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else { // ❷ +나 *가 아니면 nil 반환
		return nil
	}
}

func main() {
	// ❸ int 타입 인수 2개를 받아서 int 타입 반환을 하는 함수 타입 변수
	var operator func(int, int) int
	operator = getOperator("*")

	var result = operator(3, 4) // ❹ 함수 타입 변수를 사용해서 함수 호출
	fmt.Println(result)
}
