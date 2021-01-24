//ex10/ex10.4/ex10.4.go
package main

import "fmt"

func main() {

	day := "thursday"

	switch day {
	case "monday", "tuesday": // ❶
		fmt.Println("월, 화요일은 수업 가는 날입니다.")
	case "wednesday", "thursday", "friday": // ❷
		fmt.Println("수, 목, 금요일은 실습 가는 날입니다.")
	}
}
