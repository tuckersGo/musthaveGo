//ex12/ex12.6/ex12.6.go
package main

import "fmt"

func main() {
	a := [2][5]int{ // ❶ 이중 배열 선언
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9}, // ❷ 여러 줄에 걸쳐 초기화할 때는 쉼표를 찍자!
	}
	for _, arr := range a { // ❸ arr값은 순서대로 a[0]의 배열 a[1]의 배열
		for _, v := range arr { // ❹ v값은 순서대로 a[0]과 a[1] 배열의 각 원소
			fmt.Print(v, " ") // ➎ v값 출력
		}
		fmt.Println()
	}
}
