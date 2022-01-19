package dojo

import "fmt"

func PrintArray() {
	// ゼロ値で初期化
	//var ns1 [3]int
	// 配列リテラルで初期化
	//var ns2 = [3]int{1,2,3}
	// 要素数を値から推論
	var ns3 = [...]int{10, 20, 30}
	// 5番目が50、10番目が100で他が0の要素数11の配列
	var ns4 = [...]int{5: 50, 10: 100}

	// 長さ
	fmt.Println(len(ns4))
	// スライス演算
	fmt.Println(ns3[1:2]) // 20が出力
	// 1番目～2番目の要素を取り出す
	fmt.Println(ns3[0:2])
}
