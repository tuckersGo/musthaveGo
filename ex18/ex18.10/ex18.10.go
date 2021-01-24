//ex18/ex18.10/ex18.10.go
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2 // 삭제할 인덱스

	for i := idx + 1; i < len(slice); i++ { // ❶ 요소 앞당기기
		slice[i-1] = slice[i]
	}

	slice = slice[:len(slice)-1] // ❷ 슬라이스로 마지막 값을 잘라줍니다.

	fmt.Println(slice)
}
