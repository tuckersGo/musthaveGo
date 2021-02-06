//ex33/ex33.6/bankaccount/account.go
package bankaccount

type Account interface {
	Withdraw(money int) int
	Deposit(money int)
	Balance() int
}

func NewAccount() Account {
	return &innerAccount{balance: 1000}
}

type innerAccount struct {
	balance int
}

func (a *innerAccount) Withdraw(money int) int {
	a.balance -= money
	return a.balance
}

func (a *innerAccount) Deposit(money int) {
	a.balance += money
}

func (a *innerAccount) Balance() int {
	return a.balance
}
