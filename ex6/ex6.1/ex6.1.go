//ex6/ex6.1/ex6.1.go
package main

import "fmt"

func main() {
	var x int32 = 7 // ❶
	var y int32 = 3 // ❷

	var s float32 = 3.14 // ❸
	var t float32 = 5

	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y) // ❹ ❸
	fmt.Println("x % y = ", x%y) // ➎ ➍

	fmt.Println("s * t = ", s*t)
	fmt.Println("s / t = ", s/t) // ➏
}
