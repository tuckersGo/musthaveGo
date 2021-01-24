//ex20/ex20.6/ex20.6.go
package main

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker // ❶ 기본값은 nil입니다.
	att.Attack()     // ❷ att가 nil이기 때문에 런타임 에러가 발생합니다.
}
