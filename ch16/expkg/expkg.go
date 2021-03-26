//ch16/expkg/expkg.go
package expkg

import "fmt"

// PI 상수값은 외부로 공개되는 상수입니다.
const PI = 3.14159265359

// gravity 상수값은 외부로 공개되지 않습니다.
const gravity = 9.8

// ExportedVar 변수는 외부로 공개되는 변수입니다.
var ExportedVar int

// nonexportVar 변수는 외부로 공개되지 않는 변수입니다.
var nonexportVar int

// ExportStruct 는 외부로 공개되는 struct 타입입니다.
type ExportStruct struct {
	Name string // 외부로 공개되는 필드입니다.
	age  int    // 외부로 공개되지 않는 필드입니다.
}

type nonexportStruct struct {
	Name string // nonexportStruct struct 가 외부로 공개되지 않기 때문에 공개되지 않습니다.
	age  int    // 외부로 공개되지 않습니다.
}

// 외부로 공개되는 함수입니다.
func PrintSample() {
	fmt.Println("This is Github expkg Sample")
}

// 외부로 공개되지 않는 함수입니다.
func nonExportFunc() {
	fmt.Println("This is not exported")
}

// 원둘레를 구하는 함수
func Circumference(radius float64) float64 {
	return 2 * radius * PI
}
