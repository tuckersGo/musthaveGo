//ex22/ex22.7/ex22.7.go
package main

import "fmt"

func main() {
	m := make(map[int]int) // ❶ 맵 생성
	m[1] = 0               // ❷ 요소 추가
	m[2] = 2
	m[3] = 3

	delete(m, 3)      // ❸ 요소 삭제
	delete(m, 4)      // ❹ 없는 요소 삭제 시도
	fmt.Println(m[3]) // ➎ 삭제된 요소 값 출력
	fmt.Println(m[1]) // ➏ 존재하는 요소 값 출력
}
