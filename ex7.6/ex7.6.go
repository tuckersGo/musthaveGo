//ex7.6/ex7.6.go
package main

import "fmt"

func printNo(n int) {
	if n == 0 { // ❶ 재귀 호출 탈출 조건
		return
	}
	fmt.Println(n)
	printNo(n - 1) // ❷ 재귀 호출
}

func main() {
	printNo(3) // ❸ 함수 호출
}
