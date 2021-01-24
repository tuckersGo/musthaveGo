//ex24/ex24.4/ex24.4.go
package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex // ❶ 패키지 전역 변수 뮤텍스

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()         // ❷ 뮤텍스 획득
	defer mutex.Unlock() // ❸ defer를 사용한 Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
