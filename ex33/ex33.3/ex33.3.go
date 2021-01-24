//ex33/ex33.3//ex33.3.go
package main

import (
	"fmt"
)

func main() {
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var slice1 []int = array[1:5]           // ❶ 배열 슬라이싱
	var slice2 []int = slice1[1:8:9]        // ❷ 슬라이스 슬라이싱
	var slice3 []int = make([]int, 5)       // ❸ make()
	var slice4 []int = make([]int, 0)       // ❹ 길이 0인 슬라이스
	var slice5 []int = []int{1, 2, 3, 4, 5} // ➎ 초기화
	var slice6 []int                        // ➏ 기본값은 nil

	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2, cap(slice2))
	fmt.Println("slice3", slice3)
	fmt.Println("slice4", slice4)
	if slice4 != nil {
		fmt.Println("slice4 is not nil")
	}
	fmt.Println("slice5", slice5)
	fmt.Println("slice6", slice6)
	if slice6 == nil {
		fmt.Println("slice6 is nil")
	}
}
