//ex16/ex16.1/ex16.1.go
package main

import ( // ❶ 둘 이상의 패키지는 소괄호로 묶어줍니다.
	"fmt"
	"math/rand" // ❷ 패키지명은 rand입니다.
)

func main() {
	fmt.Println(rand.Int()) // ❸ 랜덤한 숫자값을 출력합니다.
}
