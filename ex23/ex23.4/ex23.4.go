//ex23/ex23.4/ex23.4.go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) // ❶ 스캐너 생성
	scanner.Split(bufio.ScanWords)                      // ❷ 한 단어씩 끊어읽기

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err) // ➏ 에러 감싸기
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}
	return a * b, nil
}

// 다음 단어를 읽어서 숫자로 변환하여 반환합니다.
// 변환된 숫자, 읽은 글자수, 에러를 반환합니다.
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() { // ❸ 단어 읽기
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word) // ❹ 문자열을 숫자로 변환
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s err:%w", word, err) // ➎ 에러 감싸기
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) { // ➐ 감싸진 에러가 NumError인지 확인
			fmt.Println("NumberError:", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}
