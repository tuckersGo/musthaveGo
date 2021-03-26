//appA/exA4/exA4.go
package main

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	Seed(1)
	fmt.Println(Random())
}
