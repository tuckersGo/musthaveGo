// 쇼트서킷, 좌변이 참이면 우변을 검사하지 않고 바로 실행하는 것
package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 증가")
	}

	if true || IncreaseAndReturn() < 5 {
		fmt.Println("2 증가")
	}
	fmt.Println("cnt:", cnt)
}
