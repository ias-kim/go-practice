package main

import (
	"github.com/tuckersGo/musthaveGo2/ch18/fedex"
	"github.com/tuckersGo/musthaveGo2/ch18/koreaPost"
)

// Sender 인터페이스 만들기
type Sender interface {
	Send(parcel string)
}

// Sender 인터페이스를 입력으로 받기
func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	// 우체국 전송 객체, Fedex 전송 객체 모두 SendBook 인수로 사용할 수 있다.
	// 우체국 전송 객체 만들기
	koreaPostSender := &koreaPost.PostSender{}
	SendBook("어린 왕자", koreaPostSender)
	SendBook("그리스인 조르바", koreaPostSender)

	// Fedex 전송 객체 만들기
	fedexSender := &fedex.FedexSender{}
	SendBook("어린 왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)
}