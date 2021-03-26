//ch4/ex4.5/ex4.5.go
package main

import "fmt"

func main() {
	var a int16 = 3456
	var c int8 = int8(a) // int16타입에서 int8타입으로 변환

	fmt.Println(a)
	fmt.Println(c) // int8타입인 c값 출력
}
