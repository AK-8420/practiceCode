package dojo

import (
	"sync"
	"time"
)

func WaitRoutine() {
	var w sync.WaitGroup

	w.Add(1)
	go func() { w.Done() }()
	w.Add(1)
	go func() { time.Sleep(2 * time.Second); w.Done() }()

	w.Wait() //2個揃うまで待つ
}
