//ex10/ex10.3/ex10.3.go
package main

import "fmt"

func main() {

	day := 3

	switch day {
	case 1:
		fmt.Println("첫째 날입니다")
		fmt.Println("오늘은 팀미팅이 있습니다.")
	case 2:
		fmt.Println("둘째 날입니다")
		fmt.Println("오늘은 면접이 있습니다.")
	case 3:
		fmt.Println("셋째 날입니다.")
		fmt.Println("설계안을 확정하는 날입니다.")
	case 4:
		fmt.Println("넷째 날입니다.")
		fmt.Println("예산을 확정하는 날입니다.")
	case 5:
		fmt.Println("다섯째 날입니다.")
		fmt.Println("최종 계약하는 날입니다.")
	default:
		fmt.Println("프로젝트를 진행하세요")
	}
}
