//ex9/ex9.2/ex9.2.go
package main

import "fmt"

func main() {
	temp := 33 // ❶

	if temp > 28 { // ❷
		fmt.Println("에어컨을 켠다")
	} else if temp <= 3 {
		fmt.Println("히터를 켠다")
	} else {
		fmt.Println("대기한다")
	}
}
