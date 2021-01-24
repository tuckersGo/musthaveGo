//ex14/ex14.3/ex14.3.go
package main

import "fmt"

type Data struct { // ❶ Data형 구조체
	value int
	data  [200]int
}

func ChangeData(arg Data) { // ❷ 파라미터로 Data를 받습니다.
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData(data) // ❸ 인수로 data를 넣습니다.
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100]) // ❹ data 필드 출력
}
