//ex33/ex33.4//ex33.4.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input)) // ❶ 스캐너 생성
	scanner.Split(bufio.ScanWords)                        // ❷ 단어 단위로 검색

	count := 0
	for scanner.Scan() { // ❸ 스캔 반복
		count++
	}
	if err := scanner.Err(); err != nil { // ❹
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}
