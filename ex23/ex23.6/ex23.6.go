//ex23/ex23.6/ex23.6.go
package main

import "fmt"

func f() {
	fmt.Println("f() 함수 시작")
	defer func() { // ❹ 패닉 복구
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}
	}()

	g() // ❶ g() -> h() 순서로 호출
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0)) // ❷ h() 함수 호출 - 패닉
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다.") // ❸ 패닉 발생!!
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램이 계속 실행됨") // ➎ 프로그램 실행 지속됨
}
