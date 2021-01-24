//ex15/ex15.7/ex15.7.go
package main

import "fmt"

func main() {
	str := "Hello 월드!" // ❶ 한영이 섞인 문자열

	for i := 0; i < len(str); i++ { // ❶ 문자열 크기를 얻어 순회
		// ❸ 바이트 단위로 출력
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])
	}
}
