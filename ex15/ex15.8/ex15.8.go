//ex15/ex15.8/ex15.8.go
package main

import "fmt"

func main() {
	str := "Hello 월드!" // ❶ 한영 문자가 섞인 문자열
	arr := []rune(str) // ❷ 문자열을 []rune으로 형변환

	for i := 0; i < len(arr); i++ { // ❸ 문자열 크기를 얻어 순회
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
	}
}
