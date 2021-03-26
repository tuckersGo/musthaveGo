//appB/ex33.7/ex33.7.go
package main

import (
	"fmt"

	"musthaveGo/ex33.6/bankaccount"
)

func main() {
	account := bankaccount.NewAccount()
	account.Deposit(1000)
	fmt.Println(account.Balance())
}
