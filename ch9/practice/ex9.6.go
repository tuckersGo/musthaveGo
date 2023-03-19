/*
if 초기문; 조건문 {
}
*/
package main

import "fmt"

func getMyAge() (int, bool) {
	return 28, true
}

func main() {
	// 초기문에서 선언한 변수의 범위는 if문으로 한정된다.
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("Your are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}

	//	fmt.Println("Your age is", age)
}
