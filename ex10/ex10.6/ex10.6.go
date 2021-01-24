//ex10/ex10.6/ex10.6.go
package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {
	switch age := getMyAge(); age { // getMyAge() 결과값 비교 ❶
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is", age) // age값 사용
	}

	fmt.Println("age is", age) // error - age 변수는 사라졌습니다. ❷
}
