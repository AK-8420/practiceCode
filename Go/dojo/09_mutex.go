package dojo

import (
	"fmt"
	"sync"
	"time"
)

func Mutex() {
	var m sync.Mutex
	m.Lock()
	go func() {
		time.Sleep(3 * time.Second)
		m.Unlock()
		fmt.Println("Unlock 1")
	}()
	m.Lock() //一旦ここで止まる！
	m.Unlock()
	fmt.Println("Unlock 2")
}

func RWMutex() {
	var m sync.RWMutex
	m.RLock()
	go func() {
		time.Sleep(3 * time.Second)
		m.RUnlock()
		fmt.Println("Unlock 1")
	}()
	m.RLock() //Read & Writeは止まるけど、read & readは止まらない
	m.RUnlock()
	fmt.Println("Unlock 2")
}
