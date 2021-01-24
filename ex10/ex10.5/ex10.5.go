//ex10/ex10.5/ex10.5.go
package main

import "fmt"

func main() {
	temp := 18

	switch true {
	case temp < 10 || temp > 30:
		fmt.Println("바깥 활동하기 좋은 날씨가 아닙니다.")
	case temp >= 10 && temp < 20:
		fmt.Println("약간 추울 수 있으니 가벼운 겉옷을 준비하세요")
	// 이미 두 번째 case를 실행했기 때문에 검사하지 않습니다.
	case temp >= 15 && temp < 25:
		fmt.Println("야외활동 하기 좋은 날씨입니다")
	default:
		fmt.Println("따뜻합니다.")
	}
}
