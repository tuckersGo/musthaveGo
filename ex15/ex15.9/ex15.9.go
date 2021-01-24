//ex15/ex15.9/ex15.9.go
package main

import "fmt"

func main() {
	str := "Hello 월드!"      // ❶ 한영 문자가 섞인 문자열
	for _, v := range str { // ❷ range를 이용한 순회
		fmt.Printf("타입:%T 값:%d 문자:%c\n", v, v, v) // ❸ 출력
	}
}
