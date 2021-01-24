//ex11/ex11.6/ex11.6.go
package main

import "fmt"

func main() {
	dan := 2
	b := 1
	for {
		for {
			fmt.Printf("%d * %d = %d\n", dan, b, dan*b)
			b++
			if b == 10 { // ❶
				break // 안쪽 for문을 종료합니다.
			}
		}
		b = 1
		dan++
		if dan == 10 { // ❷
			break // 바깥쪽 for문을 종료합니다.
		}
	}
	fmt.Println("for문이 종료되었습니다.")
}
