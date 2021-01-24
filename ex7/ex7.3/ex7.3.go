//ex7/ex7.3/ex7.3.go
package main

import "fmt"

func printAvgScore(name string, math int, eng int, history int) { // ❶
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, "님 평균 점수는", avg, "입니다.")
}

func main() {
	printAvgScore("김일등", 80, 74, 95) // ❷
	printAvgScore("송이등", 88, 92, 53)
	printAvgScore("박삼등", 78, 73, 78)
}
