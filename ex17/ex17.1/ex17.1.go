//ex17/ex17.1/ex17.1.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // ❶ 시간값을 랜덤 시드로 설정

	n := rand.Intn(100)
	fmt.Println(n)
}
