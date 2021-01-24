//ex26/ex26.5/ex26.5.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 { //
		fmt.Println("2개 이상의 실행인자가 필요합니다. ex) ex26.1 word filepath")
		return
	}

	word := os.Args[1]
	path := os.Args[2]

	cnt, ch := FindWordInAllFiles(word, path)
	recvCnt := 0
	for findInfo := range ch {
		fmt.Println(findInfo.filename)
		fmt.Println("--------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("--------------------------------")
		fmt.Println()
		recvCnt++
		if recvCnt == cnt {
			// all received
			break
		}
	}
}

func GetFileList(pattern string) ([]string, error) {
	filelist := []string{}
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			matched, _ := filepath.Match(pattern, info.Name())
			if matched {
				filelist = append(filelist, path)
			}
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}
	return filelist, nil
}

func FindWordInAllFiles(word, path string) (int, chan FindInfo) {
	filelist, err := GetFileList(path) // 실행인자 가져오기
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. err:", err)
		return 0, nil
	}

	ch := make(chan FindInfo)
	cnt := len(filelist)
	for _, filename := range filelist {
		go FindWordInFile(word, filename, ch)
	}
	return cnt, ch
}

func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		ch <- findInfo
		return
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	ch <- findInfo
}
