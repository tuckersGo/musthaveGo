//Done
// 추상화 레벨
// 1. 로컬(타입, 별칭)
// 2. 메소드
// 3. 인터페이스
// 3이 제일 깊다.

package main

import "fmt"

// 인터페이스 Stringer는 메소드 String을 가진다.
type Stringer interface {
	String() string
}

// 구조체 포인트 *struct는 메소드 String()을 가지고 있다.
type Student struct {
	Age int
}

// *Student 구조체 타입의 String 메소드
func (s *Student) String() string {
	return fmt.Sprintf("Studnet Age:%d", s.Age)
}

// 인터페이스 Stringer를 인자로 받는 함수 PrintAge
func PrintAge(stringer Stringer) {
	// 기존 인터페이스는 String() 메소드만 접근할 수 있고, 해당 메소드 내부로는 접근할 수 없다.
	// 인터페이스를 본래의 구체화된 타입 즉, 구조체 *Student 타입으로 변환해서 접근할 수 있게 변환한다.

	s := stringer.(*Student)
	fmt.Printf("Age: %d\n", s.Age)
}

func main() {
	s := &Student{15}
	// 인터페이스 String은 구조체 s를 접목한다.
	PrintAge(s)
}
