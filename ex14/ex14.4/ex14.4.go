//ex14/ex14.4/ex14.4.go
package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) { // ❶ 파라미터로 Data 포인터를 받습니다.
	arg.value = 999 // ❸ arg 데이터 변경
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData(&data)                      // ❷ 인수로 data의 주소를 넘깁니다.
	fmt.Printf("value = %d\n", data.value) // ❹ data 필드값 출력
	fmt.Printf("data[100] = %d\n", data.data[100])
}
