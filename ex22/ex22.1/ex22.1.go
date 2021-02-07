//ex22/ex22.1/ex22.1.go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New()       // ❶ 새로운 리스트 생성
	e4 := v.PushBack(4)   // ❷ 리스트 뒤에 요소 추가
	e1 := v.PushFront(1)  // ❸ 리스트 앞에 요소 추가
	v.InsertBefore(3, e4) // ❹ e4 요소 앞에 요소 삽입
	v.InsertAfter(2, e1)  // ➎ e1 요소 뒤에 요소 삽입

	for e := v.Front(); e != nil; e = e.Next() { // ➏ 각 요소 순회
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() { // ➐ 각 요소 역순 순회
		fmt.Print(e.Value, " ")
	}
}
