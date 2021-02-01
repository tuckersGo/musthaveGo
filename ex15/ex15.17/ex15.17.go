//ex15/ex15.17/ex15.17.go
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello"
	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str)) // ❶
	addr1 := stringheader.Data                                    // ❷

	str += " World"            // ❸
	addr2 := stringheader.Data // ➍

	str += " Welcome!"         // ➎
	addr3 := stringheader.Data // ➏

	fmt.Println(str)
	fmt.Printf("addr1:\t%x\n", addr1)
	fmt.Printf("addr2:\t%x\n", addr2)
	fmt.Printf("addr3:\t%x\n", addr3)
}
