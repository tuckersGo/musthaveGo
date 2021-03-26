//ch20/fedex/fedex.go
// Fedex에서 제공한 패키지입니다.
package fedex

import "fmt"

// Fedex에서 제공한 패키지 내 전송을 담당하는 구조체입니다.
type FedexSender struct {
}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends %v parcel\n", parcel)
}
