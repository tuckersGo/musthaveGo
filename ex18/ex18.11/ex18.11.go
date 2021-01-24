//ex18/ex18.11/ex18.11.go
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	// ❶ 맨 뒤에 요소 추가
	slice = append(slice, 0)

	idx := 2 // 추가하려는 위치

	// ❷ 맨 뒤부터 추가하려는 위치까지 값을 하나씩 옮겨줍니다.
	for i := len(slice) - 2; i >= idx; i-- {
		slice[i+1] = slice[i]
	}

	// ❸ 값 변경
	slice[idx] = 100

	fmt.Println(slice)
}
