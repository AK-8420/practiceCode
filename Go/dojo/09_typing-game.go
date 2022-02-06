package dojo

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func question() string {
	Q := []string{
		"prefiguration",
		"archetype",
		"epitome",
		"guide",
		"holotype",
		"image",
		"loadstar",
		"lodestar",
		"microcosm",
		"original",
		"paradigm",
		"pilot",
		"prototype",
		"template",
		"templet",
		"type specimen",
	}
	l := len(Q)
	t := time.Now().UnixNano()
	rand.Seed(t)
	return Q[rand.Intn(l)]
}
func inputs(r io.Reader) <-chan string {
	ch := make(chan string)
	// 無限ループで入力を受け付けるGoroutineを発行してから
	// チャンネルを返す
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func TypingGame() {
	ch1 := inputs(os.Stdin)
	ch2 := time.After(10 * time.Second)

	var q, a string
	n := 0

OuterLoop:
	for {
		q = question()
		fmt.Println(q)
		fmt.Print(">")
		select {
		case a = <-ch1:
			if q == a {
				fmt.Println("OK!")
				n++
			} else {
				fmt.Println("Wrong answer")
			}
		case <-ch2:
			fmt.Println("\nTime up! your score is", n)
			break OuterLoop
		}
	}
}
