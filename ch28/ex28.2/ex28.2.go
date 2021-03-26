//ch28/ex28.2/ex28.2.go
package main

import "fmt"

func fibonacci1(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	return fibonacci1(n-1) + fibonacci1(n-2) // ❶ 재귀 호출
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	one := 1
	two := 0
	rst := 0
	for i := 2; i <= n; i++ { // ❷ 반복문
		rst = one + two
		two = one
		one = rst
	}
	return rst
}

func main() {
	fmt.Println(fibonacci1(13))
	fmt.Println(fibonacci2(13))
}
