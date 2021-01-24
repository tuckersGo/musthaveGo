//ex10/ex10.2/ex10.2.go
package main

import "fmt"

func main() {

	day := 3

	if day == 1 {
		fmt.Println("첫째 날입니다")
		fmt.Println("오늘은 팀미팅이 있습니다.")
	} else if day == 2 {
		fmt.Println("둘째 날입니다")
		fmt.Println("새로운 팀원 면접이 있습니다.")
	} else if day == 3 {
		fmt.Println("셋째 날입니다.")
		fmt.Println("설계안을 확정하는 날입니다.")
	} else if day == 4 {
		fmt.Println("넷째 날입니다.")
		fmt.Println("예산을 확정하는 날입니다.")
	} else if day == 5 {
		fmt.Println("다섯째 날입니다.")
		fmt.Println("최종 계약하는 날입니다.")
	} else {
		fmt.Println("프로젝트를 진행하세요")
	}
}
