package main

// 해당 인터페이스는 메소드 String()을 가지고 있다.
type Stringer interface {
	String() string
}

type Student struct {
}

func (s Student) String() string {
	return fmt.println("Hi")
}

func main() {
	var stringer Stringer
	stringer.(*Student)
}
