//ex21/ex21.5//ex21.5.go
package main

import "fmt"

func main() {
	i := 0

	f := func() {
		i += 10 // ❶ i에 10 더하기
	}

	i++

	f() // ❷ f 함수 타입 변수를 사용해서 함수 리터럴 실행

	fmt.Println(i)
}
