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

// 메소드가 여러 개네?

func ConvertType(stringer Stringer) {
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
