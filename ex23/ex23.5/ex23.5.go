//ex23/ex23.5/ex23.5.go
package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다") // ❶ Panic 발생
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(9, 3)
	divide(9, 0) // ❷ Panic 발생
}
