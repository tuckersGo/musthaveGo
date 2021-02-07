//ex18/ex18.1/ex18.1.go
package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 { // ❶ slice 길이가 0인지 확인
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10 // ❷ 에러 발생
	fmt.Println(slice)
}
