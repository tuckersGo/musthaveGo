//ch20/ex20.2/ex20.2.go
package main

import "github.com/tuckersGo/musthaveGo/ch20/fedex"

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	// Fedex 전송 객체를 만듭니다.
	sender := &fedex.FedexSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}
