//ex20/ex20.9/ex20.9.go
package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
}

func (s *Student) String() string {
	return "Student"
}

type Actor struct {
}

func (a *Actor) String() string {
	return "Actor"
}

func ConvertType(stringer Stringer) {
	// ❷ 런타임 에러 발생: *Student 타입은 Stringer 인터페이스로 쓰일 수 있지만
	// stringer값이 *Student 타입이 아니기 때문에 에러가 발생합니다.
	student := stringer.(*Student)
	fmt.Println(student)
}

func main() {
	// ❶ *Actor 구조체 값을 ConvertType() 함수의 인수로 사용합니다.
	actor := &Actor{}
	ConvertType(actor)
}
