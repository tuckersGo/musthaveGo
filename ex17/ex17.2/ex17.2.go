//ex17/ex17.2/ex17.2.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n) // ❶ int 타입값을 입력받음
	if err != nil {
		stdin.ReadString('\n') // ❷ 에러 발생 시 입력스트림을 비움
	}
	return n, err
}

func main() {
	for {
		fmt.Printf("숫자값을 입력하세요:")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			fmt.Println("입력하신 숫자는 ", n, " 입니다.")
		}
	}
}
