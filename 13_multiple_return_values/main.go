package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func add(x, y int) (z1, z2 int) {
	z1 = x + y
	z2 = x - y
	return
}

func add2(v, w int) { //returns notting
	fmt.Println(v + w)
}

func add3(g, o int) int {
	return g + o
}
func main() {

	add2(22, 22) //function add2

	ans1, ans2 := add(4, 5) //function add
	fmt.Println(ans1, ans2)

	a, b := vals() //function vals
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

}
