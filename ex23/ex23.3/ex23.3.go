//ex23/ex23.3/ex23.3.go
package main

import "fmt"

type PasswordError struct { // ❶ 에러 구조체 선언
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string { // ❷ Error() 메서드
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8} // ❸ error 반환
	}

	return nil
}

func main() {
	err := RegisterAccount("myID", "myPw") // ➍ ID, PW 입력
	if err != nil {                        // ➎ 에러 확인
		if errInfo, ok := err.(PasswordError); ok { // ➏ 인터페이스 변환
			fmt.Printf("%v Len:%d RequireLen:%d\n",
				errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입되었습니다.")
	}
}
