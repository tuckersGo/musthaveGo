//ex10/ex10.1/ex10.1.go
package main

import "fmt"

func main() {

	a := 3

	switch a { // ❶
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3: // ❷
		fmt.Println("a == 3")
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 4")
	}
}
