//ch16/ch16/ex16.4/ex16.4.go
package main

import (
	"fmt"
	"tucker/custompkg"
)

func main() {
	custompkg.ExportVar = 60         // ❷ customPkg의 ExportVar값 변경
	fmt.Println(custompkg.ExportVar) // ❸ 공개된 값 출력
	fmt.Println(custompkg.ExportConst)

	var custom = custompkg.ExportStruct{Name: "Chalie"} // ➍ 공개된 값 변경
	custompkg.PrintExportStruct(custom)                 // ➎ 공개된 값 출력
}
