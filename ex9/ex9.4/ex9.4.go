//ex9/ex9.4/ex9.4.go
package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() < 5 { // ❶ 함수가 호출되지 않습니다.
		fmt.Println("1 increase")
	}
	if true || IncreaseAndReturn() < 5 { // ❷ 함수가 호출되지 않습니다.
		fmt.Println("2 increase")
	}

	fmt.Println("cnt:", cnt)
}
