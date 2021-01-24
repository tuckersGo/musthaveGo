//ex25/ex25.1/ex25.1.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) // ❶ 채널 생성

	wg.Add(1)
	go square(&wg, ch) // ❷ Go 루틴 생성
	ch <- 9            // ❸ 채널에 데이터를 넣는다.
	wg.Wait()          // ❹ 작업이 완료되길 기다린다.
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch // ➎ 데이터를 빼온다

	time.Sleep(time.Second) // 1초 대기
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}
