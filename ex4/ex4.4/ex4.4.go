//ex4/ex4.4/ex4.4.go
package main

import "fmt"

func main() {
	a := 3              // int
	var b float64 = 3.5 // float64

	var c int = int(b)  // ❶ float64에서 int로 변환
	d := float64(a * c) // int에서 float6로 변환

	var e int64 = 7
	f := int64(d) * e // int에서 int64로 변환

	var g int = int(b * 3) // float64에서 int로 변환
	var h int = int(b) * 3 // float64에서 int로 변환 g와 값이 다릅니다.
	fmt.Println(g, h, f)
}
