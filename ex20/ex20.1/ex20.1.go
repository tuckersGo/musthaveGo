//ex20/ex20.1/ex20.1.go
package main

import "fmt"

type Stringer interface { // ❶ Stringer 인터페이스 선언
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string { // ❷ Student의 String() 메서드

	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name) // ❸ 문자열 만들기
}

func main() {
	student := Student{"철수", 12} // Student 타입
	var stringer Stringer        // Stringer 타입

	stringer = student // ➎ stringer값으로 student 대입

	fmt.Printf("%s\n", stringer.String()) // ➏ stringer의 String() 메서드 호출
}
