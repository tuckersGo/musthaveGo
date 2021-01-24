//ex12/ex12.1/ex12.1.go
package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2} // ❶

	for i := 0; i < 5; i++ { // ❷
		fmt.Println(t[i]) // ❸
	}
}
