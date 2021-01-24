//ex9/ex9.6/ex9.6.go
package main

import "fmt"

func getMyAge() (int, bool) {
	return 33, true
}

func main() {

	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age) // age 값에 접근가능하다 - ❶
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}

	fmt.Println("Your age is", age) // Error - age는 소멸됐다 - ❷
}
