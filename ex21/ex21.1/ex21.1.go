//ex21/ex21.1/ex21.1.go
package main

import "fmt"

func sum(nums ...int) int { // ❶ 가변 인수를 받는 함수
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums) // // ❷ nums 타입 출력
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5)) // 5개의 인수를 사용합니다.
	fmt.Println(sum(10, 20))        // 2개의 인수를 사용합니다.
	fmt.Println(sum())              // 0개의 인수를 사용합니다.
}
