//ex18/ex18.8/ex18.8.go
package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}

	slice2 := make([]int, len(slice1)) // ❶ slice1과 같은 길이의 슬라이스 생성

	for i, v := range slice1 { // ❷ slice1의 모든 요소값 복사
		slice2[i] = v
	}

	slice1[1] = 100 // ❸ slice1 요솟값 변경
	fmt.Println(slice1)
	fmt.Println(slice2)
}
