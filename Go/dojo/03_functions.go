package dojo

import "fmt"

func Swap1(x, y int) (int, int) {
	return y, x
}

func Swap2(x, y int) (x2, y2 int) {
	y2, x2 = x, y // :=ではない
	return
}

func Swap3(x, y int) {
	//　変数使わずに入れ替えられる
	x, y = y, x
	fmt.Println(x, y)
}
