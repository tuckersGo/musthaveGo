//ex17/ex17.3/ex17.3.go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(100) // ❶ 랜덤값 생성
	cnt := 10
	for {
		fmt.Printf("숫자값을 입력하세요:")
		n, err := InputIntValue() // ❷ 숫자값 입력
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			if n > r { // ❸ 숫자값 비교
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if n < r {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수:", cnt)
				break // ❹ 같을 경우 메시지를 출력하고 break로 종료
			}
			cnt++
		}
	}
}
