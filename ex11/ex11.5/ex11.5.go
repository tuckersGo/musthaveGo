//ex11/ex11.5/ex11.5.go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ { // ❶ 5번 반복합니다.
		for j := 0; j < i+1; j++ { // ❷ 현재 i값+1만큼 반복합니다.
			fmt.Print("*") // ❸ *를 출력합니다.
		}
		fmt.Println() // ❹ 줄바꿈합니다.
	}
}
