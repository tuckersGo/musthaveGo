//ex20/ex20.5/ex20.5.go
package main

import "fmt"

func PrintVal(v interface{}) { // ❶ 빈 인터페이스를 인수로 받는 함수
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		// 그외 타입인 경우 타입과 값을 출력합니다.
		fmt.Printf("Not supported type: %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)          // int
	PrintVal(3.14)        // float64
	PrintVal("Hello")     // string
	PrintVal(Student{15}) // Student
}
