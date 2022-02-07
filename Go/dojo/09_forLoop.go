package dojo

import (
	"fmt"
	"sync"
	"time"
)

func DevideLoop() {
	fmt.Println("Start")
	start := time.Now()
	var count int
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Microsecond)
		count++
	}
	end := time.Now()

	fmt.Println("Done：", count, "times")
	fmt.Println("通常：", end.Sub(start))

	//-------------------
	var w sync.WaitGroup
	var m sync.Mutex
	p := 10
	count = 0

	fmt.Println("Start")
	start = time.Now()
	for i := 0; i < 100; i += p {
		w.Add(1)
		go func() {
			defer w.Done()

			m.Lock()
			for j := 0; j < p; j++ {
				time.Sleep(1 * time.Microsecond)
				count++
			}
			m.Unlock()
		}()
	}
	w.Wait()
	end = time.Now()

	fmt.Println("Done：", count, "times")
	fmt.Println("並列：", end.Sub(start))

	//-------------------
	ch := make(chan int)
	count = 0

	fmt.Println("Start")
	start = time.Now()
	for i := 0; i < 100; i += p {
		go func() {
			sum := 0
			for j := 0; j < p; j++ {
				time.Sleep(1 * time.Microsecond)
				sum++
			}
			ch <- sum
		}()
	}
	for i := 0; i < 100/p; i++ {
		count += <-ch
	}
	end = time.Now()

	fmt.Println("Done：", count, "times")
	fmt.Println("並列：", end.Sub(start))
}
