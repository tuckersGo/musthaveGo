//ex5.4/ex5.4.go
package main

import "fmt"

func main() {
	a := 3              // int
	var b float64 = 3.5 // float64

	var c int = int(b)  // float64에서 int로 변환 - ❶
	d := float64(a) * b // int에서 float6로 변환

	var e int64 = 7
	f := int64(a) * e // int에서 int64로 변환

	var g int = int(b * 3) // float64에서 int로 변환 - ❷
	var h int = int(b) * 3 // float64에서 int로 변환 g와 값이 다릅니다 - ❸
	fmt.Println(g, h)
}
