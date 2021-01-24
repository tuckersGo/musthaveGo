//ex26/ex26.3/ex26.3.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 찾은 라인 정보
type LineInfo struct { // ❶
	lineNo int
	line   string
}

// 파일 내 라인 정보
type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 { //
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.1 word filepath")
		return
	}

	word := os.Args[1] // ❷
	path := os.Args[2]

	findInfos := FindWordInAllFiles(word, path) // ❸ 파일 찾기
	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("--------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("--------------------------------")
		fmt.Println()
	}
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	filelist, err := GetFileList(path) // ❶ 파일 리스트 가져오기
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. err:", err)
		return findInfos
	}

	for _, filename := range filelist { // ❷ 각 파일별로 검색
		findInfos = append(findInfos, FindWordInFile(word, filename))
	}
	return findInfos
}

func FindWordInFile(word, filename string) FindInfo {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		return findInfo
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file) // ❸ 스캐너를 만듭니다.
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) { // ❹ 한 줄씩 읽으면 단어 포함 여부 검색
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	return findInfo
}
