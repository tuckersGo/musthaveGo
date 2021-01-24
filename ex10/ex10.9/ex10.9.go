//ex10/ex10.9/ex10.9.go
package main

import "fmt"

func main() {

	a := 3

	switch a { // ❶
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
		break
	case 3: // ❷
		fmt.Println("a == 3")
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 4")
	}
}
