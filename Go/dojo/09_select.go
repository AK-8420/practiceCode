package dojo

import (
	"fmt"
	"time"
)

func Select() {
	//マルチチャンネルを可能にする
	ch1 := make(chan int)
	ch2 := make(chan string)
	var ch3 chan string // nilチャンネル

	go func() {
		ch1 <- 100
		ch1 <- 200
	}()
	go func() {
		ch2 <- "Hey!"
	}()
	go func() {
		ch3 <- "none"
	}()

	select {
	case v1 := <-ch1:
		fmt.Println(v1)
	case v2 := <-ch2:
		fmt.Println(v2)
		// nilチャンネルの場合は無視
	case v3 := <-ch3:
		fmt.Println(v3)
	}
	select {
	case v1 := <-ch1:
		fmt.Println(v1)
	case v2 := <-ch2:
		fmt.Println(v2)
	case v3 := <-ch3:
		fmt.Println(v3)
	}
	select {
	case v1 := <-ch1:
		fmt.Println(v1)
	case v2 := <-ch2:
		fmt.Println(v2)
	case v3 := <-ch3:
		fmt.Println(v3)
	}
	time.Sleep(5 * time.Second)
}
