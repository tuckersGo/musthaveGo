//ex23/ex23.2/ex23.2.go
package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("제곱근은 항상 양수여야 합니다. f:%g", f) // ❶ f가 음수이면 에러 반환
	}
	return math.Sqrt(f), nil
}

func main() {
	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("Error: %v\n", err) // ❷ 에러 출력
		return
	}
	fmt.Printf("Sqrt(-2) = %v\n", sqrt)
}
