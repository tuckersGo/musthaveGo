//ex5/ex5.8/ex5.8.go
package main

import (
	"bufio" // ❶ io를 담당하는 패키지
	"fmt"
	"os" // 표준 입출력 등을 가지고 있는 패키지
)

func main() {
	stdin := bufio.NewReader(os.Stdin) // ❷ 표준 입력을 읽는 객체

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil { // 에러 발생 시
		fmt.Println(err)       // 에러 출력
		stdin.ReadString('\n') // ❸ 표준 입력 스트림 지우기
	} else {
		fmt.Println(n, a, b)
	}
	n, err = fmt.Scanln(&a, &b) // ❹ 다시 입력받기
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
