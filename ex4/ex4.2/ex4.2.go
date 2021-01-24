//ex4/ex4.2/ex4.2.go
package main

import "fmt"

func main() {
	var minimumWage int = 10 // ❶ 변수 minimumWage 선언 및 초기화
	var workingHour int = 20 // ❷ 변수 workingHour 선언 및 초기화

	// ❸ 변수 income 선언 및 초기화
	var income int = minimumWage * workingHour

	// 변수 minimumWage, workingHour, income 출력
	fmt.Println(minimumWage, workingHour, income)
}
