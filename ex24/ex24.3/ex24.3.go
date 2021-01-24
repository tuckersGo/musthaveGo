//ex24/ex24.3/ex24.3.go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}    // ❶ 0원 잔고 통장
	wg.Add(10)                // ❷ WaitGroup 객체 생성
	for i := 0; i < 10; i++ { // ❸ Go 루틴 10개 생성
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 { // ➍ 잔고가 0이하면 패닉
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000      // ➎ 1000원 입금
	time.Sleep(time.Millisecond) // ➏ 잠시 쉬고
	account.Balance -= 1000      // ➐ 1000원 출금
}
