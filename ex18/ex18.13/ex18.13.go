//ex18/ex18.13/ex18.13.go
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

// ❶ []Student의 별칭 타입 Students
type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age } // ❷ 나이 비교
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Student{
		{"화랑", 31}, {"백두산", 52}, {"류", 42},
		{"켄", 38}, {"송하나", 18}}

	sort.Sort(Students(s)) // ❸ 정렬
	fmt.Println(s)
}
