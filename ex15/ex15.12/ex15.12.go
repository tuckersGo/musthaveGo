//ex15/ex15.12/ex15.12.go
package main

import "fmt"

func main() {
	str1 := "BBB"
	str2 := "aaaaAAA"
	str3 := "BBAD"
	str4 := "ZZZ"

	fmt.Printf("%s > %s : %v\n", str1, str2, str1 > str2)   // ❶
	fmt.Printf("%s < %s : %v\n", str1, str3, str1 < str3)   // ❷
	fmt.Printf("%s <= %s : %v\n", str1, str4, str1 <= str4) // ❸
}
