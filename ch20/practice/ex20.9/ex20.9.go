// Done
package main

import "fmt"

// 인터페이스 Stringer는 메소드 String을 가진다.
// 메소드 String()을 가지기만 하면, 인터페이스로 사용될 수 있다.
type Stringer interface {
	String() string
}

type Student struct {
}

// 구조체 변수 *Student는 String() 메소드를 가진다.
func (s *Student) String() string {
	return "Student"
}

type Actor struct {
}

// 구조체 변수 *Actor는 String() 메소드를 가진다
func (a *Actor) String() string {
	return "Actor"
}

// 메소드가 여러 개네?

func ConvertType(stringer Stringer) {
	// 함수 ConvertType의 인터페이스 stringer는 구조체 포인터 변수 *Student를 가르키고 있다.
	student := stringer.(*Student)
	fmt.Println(student)
}

func main() {
	student := &Student{}
	ConvertType(student)
	// 아래, 런타임 error
	// actor := &Actor{}
	// ConvertType(actor)
}
