//ex26/ex26.1//ex26.1.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 { // ❶ 실행 인수 개수 확인
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.1 word filepath")
		return
	}

	word := os.Args[1] // ❷ 실행 인수 가져오기
	path := os.Args[2]
	fmt.Println("찾으려는 단어:", word)
	PrintAllFiles(path)
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles(path string) {
	filelist, err := GetFileList(path) // ❸ 파일목록 가져오기
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. err:", err)
		return
	}
	fmt.Println("찾으려는 파일 리스트")
	for _, name := range filelist {
		fmt.Println(name)
	}
}
