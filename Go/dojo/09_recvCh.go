package dojo

import "fmt"

func makeCh() chan int {
	return make(chan int)
}

// 受信専用のチャンネル↓
func recvCh(ch <-chan int) int {
	return <-ch
}
func ReceiveCh() {
	ch := makeCh()
	// 送信専用のチャンネル↓
	go func(ch chan<- int) {
		ch <- 333
	}(ch)
	fmt.Println(recvCh(ch))
}
