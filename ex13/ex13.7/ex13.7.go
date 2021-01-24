//ex13/ex13.7/ex13.7.go
package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 // 1바이트
	B int  // 8바이트
	C int8 // 1바이트
	D int  // 8바이트
	E int8 // 1바이트
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}
