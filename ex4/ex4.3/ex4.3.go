//ex4/ex4.3/ex4.3.go
package main

import "fmt"

func main() {
	var a int = 3 // 가장 기본 형태
	var b int     // 초깃값 생략 - 초깃값은 타입별 기본값으로 대체
	var c = 4     // 타입 생략 - 변수의 타입은 우변 값의 타입이 됩니다
	d := 5        // 선언대입문 := 을 사용해서 var 키워드와 타입 생략

	fmt.Println(a, b, c, d)
}
