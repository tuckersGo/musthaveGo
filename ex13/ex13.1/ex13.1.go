//ex13/ex13.1/ex13.1.go
package main

import "fmt"

type House struct { // ❶ House 구조체를 정의합니다.
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house House               // ❷ House 구조체 변수를 선언합니다.
	house.Address = "서울시 강동구 ..." // ❸ 각 필드값을 초기화합니다.
	house.Size = 28
	house.Price = 9.8
	house.Type = "아파트"

	fmt.Println("주소:", house.Address) // ❹ 필드값을 출력합니다.
	fmt.Printf("크기 %d평\n", house.Size)
	fmt.Printf("가격: %.2f억원\n", house.Price) // ➎ 소수점 2자리까지 출력합니다.
	fmt.Println("타입:", house.Type)
}
