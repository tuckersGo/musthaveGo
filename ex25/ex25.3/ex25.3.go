//ex25/ex25.3/ex25.3.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { // ❷ 데이터를 계속 기다린다.
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done() // ❹ 실행되지 않는다.
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2 // ❶ 데이터를 넣는다.
	}
	wg.Wait() // ❸ 작업 완료를 기다린다.
}
