//ex6/ex6.9/ex6.9.go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Float).SetString("0.1") // ❶
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b) // ❷
	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d)) // ❸
}
