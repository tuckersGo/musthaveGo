//ex18.1/ex18.1.go
package main

import "fmt"

func main() {
	var slice []int

	if slice == nil { // ❶ slice가 nil인지 확인
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10 // ❷ 에러 발생
	fmt.Println(slice)
}
