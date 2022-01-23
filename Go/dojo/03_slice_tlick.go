package dojo

import "fmt"

func PrintSlice2() {
	s1 := []int{1, 2, 3, 4, 5, 6}

	//カット
	i := 1
	j := 3
	s2 := append(s1[:i], s1[j:]...)
	fmt.Println(s2) // -> [1 4 5 6]

	//削除(1)
	s1 = []int{1, 2, 3, 4, 5, 6}
	s3 := append(s1[:i], s1[i+1:]...)
	fmt.Println(s3) // -> [1 3 4 5 6]

	//削除(2)
	s1 = []int{1, 2, 3, 4, 5, 6}
	s4 := s1[:i+copy(s1[i:], s1[i+1:])]
	fmt.Println(s4) // -> [1 3 4 5 6]
}
