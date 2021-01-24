//ex22.1/ex22.1.go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()       // ❶ 새로운 리스트 생성
	e4 := l.PushBack(4)   // ❷ 리스트 뒤에 항목 추가
	e1 := l.PushFront(1)  // ❸ 리스트 앞에 항목 추가
	l.InsertBefore(3, e4) // ❹ e4 항목 앞에 항목 삽입
	l.InsertAfter(2, e1)  // ➎ e1 항목 뒤에 항목 삽입

	for e := l.Front(); e != nil; e = e.Next() { // ➏ 각 항목 순회
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	for e := l.Back(); e != nil; e = e.Prev() { // ➐ 각 항목 역순 순회
		fmt.Print(e.Value, " ")
	}
}
