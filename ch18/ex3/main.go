package main

import (
	"github.com/tuckersGo/musthaveGo2/ch18/koreaPost"
	"github.com/tuckersGo/musthaveGo2/ch18/fedex"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	// 우체국 전송 객체를 만든다.
	send := &koreaPost.PostSender{} // *koreaPost.PostSender 타입
	SendBook("어린 왕자", send)
	SendBook("그리스인 조르바", send)
}