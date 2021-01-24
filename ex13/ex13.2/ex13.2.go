//ex13/ex13.2/ex13.2.go
package main

import "fmt"

type User struct { // 일반 고객용 구조체
	Name string
	ID   string
	Age  int
}

type VIPUser struct { // VIP 고객용 구조체
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 40},
		3,
		250, // 여러 줄로 초기화할 때는 제일 마지막 값 뒤에 꼭 쉼표를 달아주세요.
	} // ❶ User를 포함한 VIPUser 구조체 변수를 초기화합니다.

	fmt.Printf("유저: %s ID: %s 나이 %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이 %d VIP 레벨: %d VIP 가격: %d만원\n",
		vip.UserInfo.Name, // ❷ UserInfo 안의 Name
		vip.UserInfo.ID,   // ❸ UserInfo 안의 ID
		vip.UserInfo.Age,
		vip.VIPLevel, // ❹ VIPUser의 VIPLevel
		vip.Price,    // ➎ 마지막에 쉼표
	)
}
