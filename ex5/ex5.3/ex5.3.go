//ex5/ex5.3/ex5.3.go
package main

import "fmt"

func main() {
	var a = 324.13455
	var b = 324.13455
	var c = 3.14

	fmt.Printf("%08.2f\n", a) // ❶ 최소너비:8 소수점이하:2자리 0을 채움.
	fmt.Printf("%08.2g\n", b) // ❷ 최소너비:8 총숫자: 2자리 0을 채움
	fmt.Printf("%8.5g\n", b)  // ❸ 최소너비:8 총숫자: 5자리
	fmt.Printf("%f\n", c)     // ❹ 소수점이하 6자리까지 출력
}
