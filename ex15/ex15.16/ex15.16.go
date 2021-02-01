//ex15/ex15.16/ex15.16.go
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str)) // ❶
	sliceheader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str:\t%x\n", stringheader.Data) // ❷
	fmt.Printf("slice:\t%x\n", sliceheader.Data)
}
