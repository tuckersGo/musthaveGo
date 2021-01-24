//ex8.3/ex8.3.go
package main

import (
	"fmt"

	"github.com/fatih/color"
)

// ❶ 상수값에 코드를 부여합니다.
const Red int = 0
const Blue int = 1
const Green int = 2

// ❸ 코드값에 따라서 다른 색깔의 텍스트를 출력합니다.
func PrintColor(clr int, text string) {
	if clr == Red {
		color.Red(text)
	} else if clr == Blue {
		color.Blue(text)
	} else if clr == Green {
		color.Green(text)
	} else {
		fmt.Println(text)
	}
}

func main() {
	// ❷ PrintColor() 함수를 호출하여 색깔있는 글자를 출력합니다.
	PrintColor(Red, "Prints text in red.")
	PrintColor(Blue, "Prints text in blue.")
	PrintColor(Green, "Prints text in green.")
}
