//ex22/ex22.8/ex22.8.go
package main

import "fmt"

const M = 10 // ❶ 나머지 연산의 분모

func hash(d int) int {
	return d % M // ❷ 나머지 연산
}

func main() {
	m := [M]int{} // ❸ 값을 저장할 배열 생성

	m[hash(23)] = 10  // ❹ 키 23에 값 설정
	m[hash(259)] = 50 // ➎ 키 259에 값 설정

	fmt.Printf("%d = %d\n", 23, m[hash(23)])
	fmt.Printf("%d = %d\n", 259, m[hash(259)])
}
