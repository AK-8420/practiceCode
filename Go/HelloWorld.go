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

	//No.2-------
	//dojo.Print5()
	//dojo.Goto()
	//dojo.OMIKUJI(3)

	//No.3-------
	//dojo.PrintArray()
	//dojo.PrintSlice()
	//dojo.PrintSlice2()
	//dojo.PrintMap()
	//fmt.Println(dojo.Swap1(1, 2))
	//fmt.Println(dojo.Swap2(3, 4))
	//dojo.Swap3(5, 6)

	//No.9-------
	//dojo.Routine1()
	//dojo.Routine2()
	//dojo.Channel1()
	//dojo.Select()
	//dojo.ReceiveCh()
	//dojo.ForSelect()
	//dojo.Mutex()
	//dojo.RWMutex()
	//dojo.WaitRoutine()
	//dojo.TypingGame()
	dojo.DevideLoop()
}
