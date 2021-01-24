//ex9/ex9.1/ex9.1.go
package main

import "fmt"

func main() {
	light := "red"

	if light == "green" {
		fmt.Println("길을 건넌다")
	} else {
		fmt.Println("대기한다")
	}
}
