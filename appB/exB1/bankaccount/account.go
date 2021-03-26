//appB/exB1/bankaccount/account.go
package bankaccount

type Account interface { // ❶ 공개되는 인터페이스
	Withdraw(money int) int
	Deposit(money int)
	Balance() int
}

func NewAccount() Account { // ❷ 계좌 생성 함수 - 인터페이스 반환
	return &innerAccount{balance: 1000}
}

type innerAccount struct { // ❸ 공개되지 않는 구조체
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
