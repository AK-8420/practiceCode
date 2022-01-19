package dojo

import "fmt"

func PrintSlice() {
	// 長さと容量を指定して初期化(len() = 3, cap() = 10)
	// 各要素はゼロ値で初期化される
	s1 := make([]int, 3, 10)
	fmt.Println(s1)

	// スライスなので配列と違ってappend可能
	s1 = append(s1, 1)
	fmt.Println(s1)

	s2 := []int{1, 2, 3}
	s3 := append(s2, []int{4, 5, 6}...)
	fmt.Println(s3, "cap->", cap(s3)) // -> 6
	/*
		Go言語の仕様によると、スライスが容量オーバーしたら基本的に２倍の値を確保
		appendするたびにメモリ領域を確保してコピーする手順を踏むのは効率が悪いため。
	*/
	s3 = append(s3, []int{7, 8, 9}...)
	fmt.Println(s3, "cap->", cap(s3)) // -> 12
}
