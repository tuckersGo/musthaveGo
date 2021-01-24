//ex5/ex5.7/ex5.7.go
package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b) // ❶ 값을 입력받습니다.
	if err != nil {              // 에러 발생 시
		fmt.Println(err) // 에러를 출력합니다.
	} else {
		fmt.Println(n, a, b)
	}
}
