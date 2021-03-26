//appB/ex33.9//ex33.9.go
package main

import "fmt"

func CallbyCopy(n int, b [5]int, s []int) {
	n = 3000
	b[0] = 1000
	s[3] = 500
}

func main() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	var c []int
	c = array[:]
	CallbyCopy(100, array, c) // ❶

	fmt.Println("array:", array)
	fmt.Println("c:", c)
}
