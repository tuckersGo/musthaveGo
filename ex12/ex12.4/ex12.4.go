//ex12/ex12.4/ex12.4.go
package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2} // ❶

	for i, v := range t { // ❷
		fmt.Println(i, v) // ❸
	}
}
