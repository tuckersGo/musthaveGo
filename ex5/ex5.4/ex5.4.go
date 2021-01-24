//ex5/ex5.4/ex5.4.go
package main

import "fmt"

func main() {
	str := "Hello\tGo\t\tWorld\n\"Go\"is Awesome!\n" // ❶ 문자열

	fmt.Print(str)        // ❷ 문자열을 기본 서식으로 출력
	fmt.Printf("%s", str) // ❸ 문자열을 %s 서식으로 출력
	fmt.Printf("%q", str) // ➍ 문자열을 %q 서식으로 출력
}
