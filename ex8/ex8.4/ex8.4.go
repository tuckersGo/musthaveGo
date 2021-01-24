//ex8/ex8.4/ex8.4.go
package main

import "fmt"

const PI = 3.14              // ❶ 타입 없는 상수
const FloatPI float64 = 3.14 // ❷ float64 타입 상수

func main() {
	var a int = PI * 100      // ❸ 오류 발생하지 않습니다.
	var b int = FloatPI * 100 // ➍ 타입 오류 발생

	fmt.Println(a)
	fmt.Println(b)
}
