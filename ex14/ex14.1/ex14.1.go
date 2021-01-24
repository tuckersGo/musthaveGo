//ex14/ex14.1/ex14.1.go
package main

import "fmt"

func main() {
	var a int = 500
	var p *int // ❶ int 포인터 변수 p 선언

	p = &a // ❷ a의 메모리 주소를 변수 p의 값으로 대입(복사)

	fmt.Printf("p의 값: %p\n", p)            // ❸ 메모리 주솟값 출력
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p) // ❹ p가 가리키는 메모리의 값 출력
	*p = 100                               // ➎ p가 가리키는 메모리 공간의 값을 변경합니다.
	fmt.Printf("a의 값: %d\n", a)            // ➏ a값 변화 확인
}
