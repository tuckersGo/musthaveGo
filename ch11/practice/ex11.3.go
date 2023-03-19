// Done
// continue, break문
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// NewReader는 인수로 입력되는 입력 스트림으로 Reader 객체 생성
	// 표준 입력 스트림 os.Stdin
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력하세요.")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요.")
			// 키보드 버퍼 제거
			// 줄바꿈 문자가 나올 때까지 읽기, os.Stdin이 비워진다.
			// 줄바꿈 문자가 나올 때까지 계속 읽어들인다. 마지막 줄바꿈 문자는 if문에 빠져나가면 사라진다.
			stdin.ReadString('\n')
			continue
		}

		fmt.Printf("입력하신 숫자는 %d입니다. \n", number)
		if number%2 == 0 {
			break
		}
	}
	fmt.Println("for문이 종료됬습니다")
}
