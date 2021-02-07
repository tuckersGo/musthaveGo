//ex18/ex18.6/ex18.6.go
package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3} // ❶ len:3 cap:3 슬라이스 생성

	slice2 := append(slice1, 4, 5) // ❷ append() 함수로 요소 추가

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100 // ❸ slice1 요솟값 변경

	fmt.Println("After change second element")
	fmt.Println("slice:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500) // ➍ slice1 요솟값 변경

	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}
