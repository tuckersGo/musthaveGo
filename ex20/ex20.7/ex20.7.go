//ex20/ex20.7/ex20.7.go
package main

import "fmt"

type Stringer interface { // ❶ 인터페이스
	String() string
}

type Student struct { // ❷ 구조체
	Age int
}

func (s *Student) String() string { // ❸ Student 타입의 String() 메서드
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) { // ➍

	s := stringer.(*Student)       // ➎ *Student 타입으로 타입 변환
	fmt.Printf("Age: %d\n", s.Age) // ➏ s.Age 출력

}

func main() {
	s := &Student{15} // ➐ *Student 타입 변수 s 선언 및 초기화

	PrintAge(s) // ➑ 변수 s 를 인터페이스 인수로 PrintAge() 함수 호출
}
