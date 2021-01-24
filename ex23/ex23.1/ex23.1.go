//ex23/ex23.1/ex23.1.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename) // ❶ 파일 열기
	if err != nil {
		return "", err // ❷ 에러 나면 에러 반환
	}
	defer file.Close()          // ❸ 함수 종료 직전 파일 닫기
	rd := bufio.NewReader(file) // ❹ 파일 내용 읽기
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename) // ➎ 파일 생성
	if err != nil {                  // ➏ 에러 나면 에러 반환
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line) // ➐ 파일에 문자열 쓰기
	return err
}

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename) // ➑ 파일 읽기 시도
	if err != nil {
		err = WriteFile(filename, "This is WriteFile") // ➒ 파일 생성
		if err != nil {                                // ❿ 에러를 처리
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}
		line, err = ReadFile(filename) // ⓫ 다시 읽기 시도
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다.", err)
			return
		}
	}
	fmt.Println("파일내용:", line) // ⓬ 파일 내용 출력
}
