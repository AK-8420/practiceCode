package main

import (
	"Go/dojo"
	"fmt"
)

//他のパッケージからユーザ定義型を持ってくる
type TestUser dojo.User

func main() {
	const s = "Hello" + ", world!"
	fmt.Println(s)

	//dojo.Print5()
	//dojo.Goto()
	//dojo.OMIKUJI(3)
	//dojo.PrintArray()
	//dojo.PrintSlice()
	//dojo.PrintSlice2()
	//dojo.PrintMap()
	fmt.Println(dojo.Swap1(1, 2))
	fmt.Println(dojo.Swap2(3, 4))
	dojo.Swap3(5, 6)
}
