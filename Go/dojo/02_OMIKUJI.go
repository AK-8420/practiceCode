package dojo

import (
	"fmt"
	"math/rand"
	"time"
)

func OMIKUJI(n int) {
	t := time.Now().UnixNano()
	rand.Seed(t)
	for i := 0; i < n; i++ {
		s := rand.Intn(6) + 1
		switch s {
		case 1:
			fmt.Println("凶")
		case 2, 3:
			fmt.Println("吉")
		case 4, 5:
			fmt.Println("中吉")
		case 6:
			fmt.Println("大吉")
		}
	}
}
