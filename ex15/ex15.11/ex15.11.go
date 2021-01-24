//ex15/ex15.11/ex15.11.go
package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "Hell"
	str3 := "Hello"

	fmt.Printf("%s == %s  : %v\n", str1, str2, str1 == str2)
	fmt.Printf("%s != %s : %v\n", str1, str2, str1 != str2)
	fmt.Printf("%s == %s : %v\n", str1, str3, str1 == str3)
	fmt.Printf("%s != %s : %v\n", str1, str3, str1 != str3)
}
