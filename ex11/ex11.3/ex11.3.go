//ex11/ex11.3/ex11.3.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for { // ❶ 무한 루프
		fmt.Println("입력하세요")
		var number int
		_, err := fmt.Scanln(&number) // ❷ 한줄 입력을 받습니다.
		if err != nil {               // ❸ 숫자가 아닌 경우
			fmt.Println("숫자를 입력하세요")

			// 키보드 버퍼를 비웁니다.
			stdin.ReadString('\n') // ❹ 키보드 버퍼를 지워줍니다.
			continue               // ➎ ❶ 로 돌아갑니다
		}
		fmt.Printf("입력하신 숫자는 %d입니다\n", number)
		if number%2 == 0 { // ➏ 짝수 검사를 합니다.
			break // ➐ for문을 종료합니다.
		}
	}
	fmt.Println("for문이 종료되었습니다.")
}
