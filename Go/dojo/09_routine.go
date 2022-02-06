package dojo

import "fmt"
import "time"

func Routine1() {
	go func() {
		fmt.Println("別のゴールーチン")
	}()
	fmt.Println("mainゴールーチン")
	time.Sleep(50 * time.Millisecond)
}

func Routine2() {
	defer fmt.Println("All finished.")

	go func() {
		defer fmt.Println("goroutine1 done.")
	}()
	go func() {
		defer fmt.Println("goroutine2 done.")
		time.Sleep(1 * time.Second)
	}()

	time.Sleep(3 * time.Second)
}
