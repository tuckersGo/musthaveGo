//ex9/ex9.5/ex9.5.go
package main

import "fmt"

// 친구 중 부자가 있는가 반환 - 무조건 true 반환
func HasRichFriend() bool {
	return true
}

// 같이 간 친구 숫자를 반환 - 무조건 3을 반환
func GetFriendsCount() int {
	return 3
}

func main() {
	price := 35000

	if price > 50000 { // ❶
		if HasRichFriend() {
			fmt.Println("앗 신발끈이 풀렸네")
		} else {
			fmt.Println("나눠내자")
		}
	} else if price >= 30000 && price <= 50000 { // ❷
		if GetFriendsCount() > 3 { // ❸
			fmt.Println("어이쿠 신발끈이..")
		} else {
			fmt.Println("사람도 얼마 없는데 나눠내자")
		}
	} else {
		fmt.Println("오늘은 내가 쏜다")
	}
}
