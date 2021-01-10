//ex16.4/ex16.4.go
package main

import (
	"fmt"
	"tucker/customPkg" // ❶ customPkg 패키지 임포트
)

func main() {
	customPkg.ExportVar = 60         // ❷ customPkg의 ExportVar값 변경
	fmt.Println(customPkg.ExportVar) // ❸ 공개된 값 출력
	fmt.Println(customPkg.ExportConst)

	var custom = customPkg.ExportStruct{Name: "Chalie"} // ➍ 공개된 값 변경
	customPkg.PrintExportStruct(custom)                 // ➎ 공개된 값 출력
}
