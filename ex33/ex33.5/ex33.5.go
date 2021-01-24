//ex33/ex33.5/ex33.5.go
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Errorf("Create: %v\n", err)
		return
	}

	defer f.Close()

	const name, age = "Kim", 22
	n, err := fmt.Fprint(f, name, " is ", age, " years old.\n")

	// The n and err return values from Fprint are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Errorf("Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")
}
