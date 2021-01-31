//ex16/ex16.3/ex16.3.go
package main

import (
	"fmt"

	"github.com/tuckersGo/musthaveGo/ex16/expkg"
)

func main() {
	expkg.PrintSample()   // ❶ expkg 패키지 함수 사용
	fmt.Println(expkg.PI) // ❷ expkg 패키지 상수 사용
}
