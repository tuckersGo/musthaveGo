//ex17/ex17.2/ex17.2.go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	BALANCE        = 1000
	EARN_POINT     = 500
	LOSE_POINT     = 100
	VICTORY_POINT  = 5000
	GAMEOVER_POINT = 0
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n) // ❶ int 타입값을 입력받습니다.
	if err != nil {
		stdin.ReadString('\n') // ❷ 에러 발생시 입력스트림을 비웁니다.
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())

	balance := BALANCE

	for {
		fmt.Print("1~5사이의 값을 입력하세요:")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else if n < 1 || n > 5 {
			fmt.Println("1~5사이의 값만 입력하세요.")
		} else {
			r := rand.Intn(5) + 1
			if n == r {
				balance += EARN_POINT
				fmt.Println("축하합니다. 맞추셨습니다. 남은 돈:", balance)
				if balance >= VICTORY_POINT {
					fmt.Println("게임 승리")
					break
				}
			} else {
				balance -= LOSE_POINT
				fmt.Println("꽝 아쉽지만 다음 기회를.. 남은 돈:", balance)
				if balance <= GAMEOVER_POINT {
					fmt.Println("게임 오버")
					break
				}
			}
		}
	}
}
