//ex12/ex12.3/ex12.3.go
package main

import "fmt"

func main() {
	nums := [...]int{10, 20, 30, 40, 50} // ❶

	nums[2] = 300 // ❷

	for i := 0; i < len(nums); i++ { // ❸
		fmt.Println(nums[i]) // ➍
	}
}
