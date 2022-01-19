package dojo

import "fmt"

func Print5() {
	// 1つ目に宣言した定数の右辺を下に引き継ぐ
	const (
		a = iota
		b
	)
	// シフト演算
	const (
		c = 1 << iota
		d
		e
	)
	fmt.Println(a, b, c, d, e)
}
