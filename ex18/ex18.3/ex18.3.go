//ex18/ex18.3/ex18.3.go
package main

import "fmt"

func main() {
	var slice []int

	for i := 1; i <= 10; i++ { // ❶ 요소를 하나씩 추가
		slice = append(slice, i)
	}

	slice = append(slice, 11, 12, 13, 14, 15) // ❷ 한 번에 여러 요소 추가
	fmt.Println(slice)
}
