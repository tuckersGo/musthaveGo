//ex21/ex21.4/ex21.4.go
package main

import "fmt"

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {

		return func(a, b int) int { // ❶ 함수 리터럴을 사용해서 더하기 함수를 정의하고 반환
			return a + b
		}
	} else if op == "*" {

		return func(a, b int) int { // ❷ 함수 리터럴을 사용해서 곱하기 함수를 정의하고 반환
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator("*")

	result := fn(3, 4) // ❸ 함수 타입 변수를 사용해서 함수 호출
	fmt.Println(result)
}
