//ex24/ex24.5/ex24.5.go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{} // ❶ 포크와 수저 뮤텍스
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "포크", "수저") // ❷ A는 포크 먼저
	go diningProblem("B", spoon, fork, "수저", "포크") // ❸ B는 수저 먼저
	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstname, secondname string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다.\n", name)
		first.Lock() // ❹ 첫 번째 뮤텍스를 획득 시도
		fmt.Printf("%s %s 획득\n", name, firstname)
		second.Lock() // ➎ 두 번째 뮤텍스를 획득 시도
		fmt.Printf("%s %s 획득\n", name, secondname)

		fmt.Printf("%s 밥을 먹습니다\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock() // ➏ 뮤텍스 반납
		first.Unlock()
	}
	wg.Done()
}
