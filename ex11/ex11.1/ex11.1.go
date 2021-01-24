//ex11/ex11.1/ex11.1.go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ { // ❶ 초기문; 조건문; 후처리
		fmt.Print(i, ", ") // ❸ i 값을 출력합니다.
	}

	//  fmt.Println(i)              // ❸ Error - i는 이미 사라졌습니다.
}
