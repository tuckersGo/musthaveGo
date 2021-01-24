//ex25/ex25.8/ex25.8.go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background()) // ❶ 컨텍스트 생성
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel() // ❷ 취소

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done(): // ❸ 취소 확인
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
