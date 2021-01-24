//ex7/ex7.1/ex7.1.go
package main

import "fmt"

func Add(a int, b int) int { // ❶
	return a + b // ❷
}

func main() {
	c := Add(3, 6) // ❸
	fmt.Println(c) // ➍
}
