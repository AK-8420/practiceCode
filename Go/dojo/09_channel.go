package dojo

import (
	"fmt"
	"time"
)

func Channel1() {
	ch := make(chan int)

	go func() {
		ch <- 100
	}()

	go func() {
		v := <-ch
		fmt.Println(v)
	}()

	time.Sleep(2 * time.Second)
}
