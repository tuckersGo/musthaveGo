//ex13/ex13.6/ex13.6.go
package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   // ❶ 4바이트
	Score float64 // 8바이트
}

func main() {
	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user))
}
