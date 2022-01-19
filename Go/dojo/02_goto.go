package dojo

import "fmt"

func Goto() {
JUMP:
	for i := 0; ; i++ {
		if i%2 == 1 {
			break JUMP
		}
	}
	fmt.Println("finish.")
}
