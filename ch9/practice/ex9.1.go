package main

import "fmt"

func main() {
	temp := 33

	if temp > 28 {
		fmt.Println("에어컨을 키다")
	} else if temp <= 3 {
		fmt.Println("히터를 켜다")
	} else {
		fmt.Println("대기하다")
	}
}
