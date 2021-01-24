//ex10/ex10.8/ex10.8.go
package main

import "fmt"

type ColorType int // 별칭 ColorType을 선언하고 Const 열거값 정의 ❶
const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

// 각 ColorType 열거값에 따른 문자열을 반환하는 함수
func colorToString(color ColorType) string { // ❷
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}
