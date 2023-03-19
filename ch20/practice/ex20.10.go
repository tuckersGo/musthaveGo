// Done
package main

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

// 구조체 포인터 변수 *file(리시버)는 Read()라는 메소드 함수를 가지고 있다. Close()라는 메소드 함수는 없다
func (f *File) Read() {
}

func ReadFile(reader Reader) {
	// 인터페이스를 Closer() 구조체 변수 타입으로 변환한다. -> 런타임 에러 발생
	// 타입 변환 성공 여부 반환
	// .(convertType)은 두 개를 변수로 받으면, 두 번째로는 불리언 타입(True, False)로 반환해준다. 따라서, 런타임 에러는 발생하지 않는다.

	if c, ok := reader.(Closer); ok {
		c.Close()
	}

	/*
		c, ok := reader.(Closer)
		 if ok {
		 	c.Close()
		 }
		 fmt.Println("not converted")
	*/
}

func main() {
	file := &File{}
	ReadFile(file)
}
