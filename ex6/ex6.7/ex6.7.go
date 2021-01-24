//ex6/ex6.7/ex6.7.go
package main

import "fmt"

const epsilon = 0.000001 // ❶ 매우 작은 값

func equal(a, b float64) bool {
	if a > b {
		if a-b <= epsilon { // ❷ 작은 차이는 무시합니다.
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)          // ❸
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c)) // ❹

	a = 0.0000000000004 // ➎ 매우 작은 값으로 변경
	b = 0.0000000000002
	c = 0.0000000000007

	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))
}
