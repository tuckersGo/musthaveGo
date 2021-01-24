//ex20/ex20.3/ex20.3.go
package main

import (
	"github.com/tuckersGo/musthaveGo/ex20/koreaPost"
	"github.com/tuckersGo/musthavego/ex20/fedex"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	// 우체국 전송 객체를 만듭니다.
	sender := &koreaPost.PostSender{} // ❶ *koreaPost.PostSender 타입
	SendBook("어린 왕자", sender)         // ❷ 타입이 맞지 않습니다.
	SendBook("그리스인 조르바", sender)
}
