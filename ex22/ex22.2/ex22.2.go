//ex22/ex22.2/ex22.2.go
package main

import (
	"container/list"
	"fmt"
)

type Queue struct { // ❶ Queue 구조체 정의
	v *list.List
}

func (q *Queue) Push(val interface{}) { // ❷ 요소 추가
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} { // ❸ 요소을 반환하면서 삭제
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	queue := NewQueue() // ❹ 새로운 큐 생성

	for i := 1; i < 5; i++ { // ➎ 요소 입력
		queue.Push(i)
	}
	v := queue.Pop()
	for v != nil { // ➏ 요소 출력
		fmt.Printf("%v -> ", v)
		v = queue.Pop()
	}
}
