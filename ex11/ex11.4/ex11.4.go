//ex11/ex11.4/ex11.4.go
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ { // ❶ 3번 반복합니다.
		for j := 0; j < 5; j++ { // ❷ 5번 반복합니다.
			fmt.Print("*") // ❸ * 를 출력합니다.
		}
		fmt.Println() // ❹ 줄바꿈합니다.
	}
}
