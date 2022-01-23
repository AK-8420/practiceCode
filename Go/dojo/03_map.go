package dojo

import "fmt"

func PrintMap() {
	// makeで初期化
	// 容量指定できる
	//m1 = make(map[string]int, 10)

	// リテラルで初期化
	//m2 := map[string]int{} //空の場合
	m3 := map[string]int{"x": 10, "y": 5}

	value, ok := m3["z"]
	fmt.Println(value, ok) // ゼロ値とfalseを返す

	m3["z"] = 1
	delete(m3, "z") // 削除
	value, ok = m3["z"]
	fmt.Println(value, ok)
}
