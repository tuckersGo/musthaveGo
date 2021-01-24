//ex6/ex6.5/ex6.5.go
package main

import "fmt"

func main() {
	var x int8 = 127 // ❶ 8비트 부호가 있는 정수 최댓값

	fmt.Printf("%d < %d + 1: %v\n", x, x, x < x+1) // ❷ 비교 연산 수행
	fmt.Printf("x\t= %4d, %08b\n", x, x)           // ❸ x값 확인
	fmt.Printf("x + 1\t= %4d, %08b\n", x+1, x+1)   // ➍ x + 1값 확인
	fmt.Printf("x + 2\t= %4d, %08b\n", x+2, x+2)   // ➎ x + 2값 확인
	fmt.Printf("x + 3\t= %4d, %08b\n", x+3, x+3)   // ➏ x + 3값 확인

	var y int8 = -128                              // 8비트 부호있는 정수 최솟값
	fmt.Printf("%d > %d - 1: %v\n", y, y, y > y-1) // ➐ 비교 연산 수행
	fmt.Printf("y\t= %4d, %08b\n", y, y)           // ➑ y값 확인
	fmt.Printf("y - 1\t= %4d, %08b\n", y-1, y-1)   // ➒ y - 1값 확인
}
