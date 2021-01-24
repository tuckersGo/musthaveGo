//ex13/ex13.3/ex13.3.go
package main

import "fmt"

type User struct { // 일반 고객용 구조체
	Name string
	ID   string
	Age  int
}

type VIPUser struct { // VIP 고객용 구조체
	User     // ❶ 필드명 생략
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 40},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이 %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이 %d VIP 레벨: %d VIP 가격: %d만원\n",
		vip.Name, // ❷ . 하나로 접근할 수 있습니다.
		vip.ID,
		vip.Age,
		vip.VIPLevel,
		vip.Price, // 여러 줄로 초기화할 때는 제일 마지막 값 뒤에 꼭 쉼표를 달아주세요.
	)
}
