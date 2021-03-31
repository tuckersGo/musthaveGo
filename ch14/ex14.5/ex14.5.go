package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u // 1 탈출 분석으로 u 메모리가 사라지지 않음
}

func main() {
	userPointer := NewUser("AAA", 23)

	fmt.Println(userPointer)
}
