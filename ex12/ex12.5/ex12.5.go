//ex12/ex12.5/ex12.5.go
package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a { // ❶ 배열 a 원소 출력
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()         // ❷ 개행
	for i, v := range b { // ❸ 배열 b 원소 출력
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a // ❹ a 배열을 b변수에 복사

	fmt.Println()         // 개행
	for i, v := range b { // ➎ 배열 b 원소 출력
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}
