package main

import "fmt"

//Done
// 인터페이스 Stringer는 메소드 String을 조절 할 수 있다.
type Stringer interface {
	String() string
}

// 구조체, 필드 정의
type Student struct {
	Name string
	Age  int
}

// 인터페이스에 구현하지 않는 메소드는 별도로 구현
// 리시버 s는 Student 구조체 타입을 가진다. 또한, 메소드 String을 가진다.
func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func main() {
	student := Student{"철수", 12} // 구조체 student 초기화
	var stringer Stringer        // 인터페이스 stringer는 Stringer 인터페이스이다. 인터페이스 안의 메소드 String은 구조체 Student를 가지고 있다

	stringer = student                    // 인터페이스에 구조체 접목.
	fmt.Printf("%s\n", stringer.String()) // 인터페이스로 메소드에 접근. 메소드는 구조체를 사용할 수 있다.
}
