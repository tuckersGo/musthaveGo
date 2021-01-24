//ex22/ex22.6/ex22.6.go
package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product) // ❶ 맵 생성

	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 200}
	m[78] = Product{"자", 1000}
	m[345] = Product{"샤프", 3000}
	m[897] = Product{"샤프심", 500}

	for k, v := range m { // ❸ 맵 순회
		fmt.Println(k, v)
	}
}
