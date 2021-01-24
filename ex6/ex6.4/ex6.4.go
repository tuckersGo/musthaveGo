//ex6/ex6.4/ex6.4.go
package main

import "fmt"

func main() {
	var x int8 = 16   // ❶ 부호가 있는 정수, 부호 비트값이 0인 수
	var y int8 = -128 // ❷ 부호가 있는 정수, 부호 비트값이 1인 수
	var z int8 = -1   // ❸ 모든 비트값이 1인 정수
	var w uint8 = 128 // ➍ 부호 없는 정수, 최상위 비트값이 1인 양수

	fmt.Printf("x:%08b x>>2: %08b x>>2: %d\n", x, x>>2, x>>2)               // ➎ 오른쪽 시프트❶
	fmt.Printf("y:%08b y>>2: %08b y>>2: %d\n", uint8(y), uint8(y>>2), y>>2) // ➏ 오른쪽 시프트❷
	fmt.Printf("z:%08b z>>2: %08b z>>2: %d\n", uint8(z), uint8(z>>2), z>>2) // ➐ 오른쪽 시프트❸
	fmt.Printf("w:%08b w>>2: %08b w>>2: %d\n", uint8(w), uint8(w>>2), w>>2) // ➑ 오른쪽 시프트❹
}
