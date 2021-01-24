//ex12/ex12.2/ex12.2.go
package main

const Y int = 3 // ❶

func main() {
	x := 5                     // ❷
	a := [x]int{1, 2, 3, 4, 5} // ❸

	b := [Y]int{1, 2, 3} // ➍

	var c [10]int // ➎
}
