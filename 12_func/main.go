package main

import "fmt"

func add(a int, b int) int { //1.two number addition
	temt := 0
	temt = a + b
	return temt
}
func squ(s int) int { //2.square of number
	return s * s
}
func swipping(s1, s2 *int) { //3.swiping of number
	// var temp int
	temp := *s1
	*s1 = *s2
	*s2 = temp
}
func main() {
	fmt.Println("1.addition of number")
	fmt.Println("2.square of number")
	fmt.Println("3.swiping of number")
	fmt.Println("------------------------")
	var swi int
	fmt.Println("enter your choice")
	fmt.Scanln(&swi)

	switch swi {
	case 1:
		var val, val2 int
		fmt.Println("enter the number")
		fmt.Scanln(&val)
		fmt.Println("enter the second number")
		fmt.Scanln(&val2)

		d := add(val, val2)
		println("sum of your addtion  is ", d)
	case 2:
		var sq1 int
		fmt.Println("enter the number")
		fmt.Scanln(&sq1)
		s := squ(sq1)
		fmt.Printf("square is %d", s)
	case 3:
		var swv1, swv2 int
		fmt.Println("enter the first value ")
		fmt.Scanln(&swv1)
		fmt.Println("enter the second value ")
		fmt.Scanln(&swv2)

		fmt.Printf("before swipping value is v1 %d v2 %d \n", swv1, swv2)
		swipping(&swv1, &swv2)
		fmt.Printf("before swipping value is v1 %d v2 %d", swv1, swv2)

	default:
		fmt.Println("enterd invalid number")
	}

}
