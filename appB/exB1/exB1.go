//appB/exB1/exB1.go
package main

import (
	"fmt"

	"musthaveGo/exB1/bankaccount"
)

func main() {
	account := bankaccount.NewAccount()
	account.Deposit(1000)
	fmt.Println(account.Balance())
}
