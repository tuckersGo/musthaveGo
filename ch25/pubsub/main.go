package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(4)
	publisher := NewPublisher(ctx)
	subscriber1 := NewSubscriber("AAA", ctx)
	subscriber2 := NewSubscriber("BBB", ctx)

	go publisher.Update()

	subscriber1.Subscribe(publisher)
	subscriber2.Subscribe(publisher)

	go subscriber1.Update()
	go subscriber2.Update()

	go func() {
		tick := time.Tick(time.Second * 2)
		for {
			select {
			case <-tick:
				publisher.Publish("Hello Message")
			case <-ctx.Done():
				wg.Done()
				return
			}
		}
	}()

	fmt.Scanln()
	cancel()

	wg.Wait()
}
