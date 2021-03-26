//ch20/koreaPost/post.go
// 우체국에서 제공한 패키지입니다.
package koreaPost

import "fmt"

// 우체국에서 제공한 패키지 내 전송을 담당하는 구조체입니다.
type PostSender struct {
}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("우체국에서 택배 %v를 보냅니다.\n", parcel)
}
