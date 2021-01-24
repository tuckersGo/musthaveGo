//ex6/ex6.2/ex6.2.go
package main

import "fmt"

func main() {
	var x1 int8 = 34   // ❶ 8비트 정수, 00100010
	var x2 int16 = 34  // ❷ 16비트 정수, 00000000 00100010
	var x3 uint8 = 34  // ❸ 8비트 부호가 없는 정수, 00100010
	var x4 uint16 = 34 // ➍ 16비트 부호가 없는 정수,00000000 00100010

	fmt.Printf("^%d = %5d,\t %08b\n", x1, ^x1, uint8(^x1)) // ➎
	fmt.Printf("^%d = %5d,\t %016b\n", x2, ^x2, uint16(^x2))
	fmt.Printf("^%d = %5d,\t %08b\n", x3, ^x3, ^x3)
	fmt.Printf("^%d = %5d,\t %016b\n", x4, ^x4, ^x4)
}
