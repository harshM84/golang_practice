package main

import "fmt"

func main() {
	var s []string
	// fmt.Println("uninit:", s, s == nil, len(s) == 0)
	// fmt.Println(s)

	s = make([]string, 3)
	fmt.Println("emp", s, "len", (s), "cap", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s, "asd", "qwe")
	fmt.Println("append", s)
	fmt.Println("set:", s)

	d := make([]string, len(s))
	copy(d, s)
	fmt.Println("copy", d)
	/////////////////////////////////////////////////////////////////////
	var km []int
	km = make([]int, 2)
	fmt.Println("emp", km, "len", (km), "cap", cap(km))

	km[0] = 12
	km[1] = 34
	fmt.Println(km)

	l := s[1:4]
	fmt.Println(l)

	///////////////////////////////////////////////////////////////

	// twoD := make([][]int, 3)
	// for i := 0; i <= 3; i++ {
	// 	innerLen := i + 1
	// 	twoD[i] = make([]int, innerLen)
	// 	for j := 0; j < innerLen; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	// fmt.Println("2d:-", twoD)
}
