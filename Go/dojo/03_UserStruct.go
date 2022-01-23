package dojo

type User struct {
	userID int         // 模範解答だとstringだった
	result map[int]int // key = 試合番号、value = スコア
}
