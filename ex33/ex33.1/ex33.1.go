//ex33/ex33.1//ex33.1.go
package main

import "fmt"

func main() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	var b [5]int
	b = array // ❶

	var c []int
	// c = array 는 안된다.
	c = array[:] // ❷

	b[0] = 1000 // ❸
	c[3] = 500

	fmt.Println("array:", array)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}
