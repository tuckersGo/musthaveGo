//ex5/ex5.2/ex5.2.go
package main

import "fmt"

func main() {
	var a = 123
	var b = 456
	var c = 123456789

	fmt.Printf("%5d, %5d\n", a, b)    // ❶ 최소 너비보다 짧은 값 너비 지정
	fmt.Printf("%05d, %05d\n", a, b)  // ❷ 최소 너비보다 짧은 값 0 채우기
	fmt.Printf("%-5d, %-05d\n", a, b) // ❸ 최소 너비보다 짧은 값 왼쪽 정렬

	fmt.Printf("%5d, %5d\n", c, c)    // ➍ 최소 너비보다 긴 값 너비 지정
	fmt.Printf("%05d, %05d\n", c, c)  // ➎ 최소 너비보다 긴 값 0 채우기
	fmt.Printf("%-5d, %-05d\n", c, c) // ➏ 최소 너비보다 긴 값 왼쪽 정렬
}
