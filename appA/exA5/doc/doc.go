// Package doc This package is example package for explaining go doc.
// You can see the detail of go doc with https://pkg.go.dev/golang.org/x/tools/cmd/godoc link.
//appA/exA5/doc/doc.go
package doc

import "fmt"

//CharSize 상수 설명입니다.
const CharSize = 3

const (
	//CharColorRed 빨간색 -iota
	CharColorRed = iota
	//CharColorBlue 파란색 -iota
	CharColorBlue
	//CharColorGreen 초록색 - iota
	CharColorGreen
)

//PrintDoc 문서 출력 함수 설명입니다.
func PrintDoc() {

}

//TextDoc 문서 작성시 구조체 예제입니다.
type TextDoc struct {
	//Msg 내부 메세지입니다.
	Msg string

	//size 를 나타냅니다.
	size int
}

//NewTextDoc 외부로 공개되는 함수 설명입니다.
func NewTextDoc() *TextDoc {
	return &TextDoc{}
}

//PrintDoc 외부로 공개되는 메서드 설명입니다.
// t.PrintDoc()
//같이 사용합니다.
func (t *TextDoc) PrintDoc() {
	fmt.Println("This is TextDoc PrintDoc method")
}
