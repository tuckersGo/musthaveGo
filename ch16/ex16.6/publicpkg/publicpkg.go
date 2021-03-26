//ch16/ex16.6/publicpkg/publicpkg.go
package publicpkg

import "fmt"

const (
	PI = 3.1415   // 외부로 공개되는 상수
	pi = 3.141516 // 외부로 공개되지 않는 상수
)

var ScreenSize int = 1080 // 외부로 공개되는 변수
var screenHeight int      // 외부로 공개되지 않는 변수

func PublicFunc() { // 외부로 공개되는 함수
	const MyConst = 100 // 외부로 공개되지 않습니다.
	fmt.Println("This is a public function", MyConst)
}

func privateFunc() { // 외부로 공개되지 않는 함수
	fmt.Println("This is a private function")
}

type MyInt int       // 외부로 공개되는 별칭 타입
type myString string // 외부로 공개되지 않는 별칭 타입

type MyStruct struct { // 외부로 공개되는 구조체
	Age  int    // 외부로 공개되는 구조체 필드
	name string // 외부로 공개되지 않는 구조체 필드
}

func (m MyStruct) PublicMethod() { // 외부로 공개되는 메서드
	fmt.Println("This is a public method")
}

func (m MyStruct) privateMethod() { // 외부로 공개되지 않는 메서드
	fmt.Println("This is a private method")
}

type myPrivateStruct struct { // 외부로 공개되지 않는 구조체
	Age  int    // 외부로 공개되지 않는 구조체 필드
	name string // 외부로 공개되지 않는 구조체 필드
}

func (m myPrivateStruct) PrivateMethod() { // 외부로 공개되지 않는 메서드
	fmt.Println("This is a private method")
}
