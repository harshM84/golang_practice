package main

import (
	"fmt"
)

func main() {
	// var num int

	// fmt.Printf("enter the fist number")
	// fmt.Scanln(&num)

	// fmt.Printf("this first name %v", fname)
	// var num = 2
	// var numm int
	// for w := 1; w <= 10; w++ {
	// 	numm = num * w
	// 	fmt.Printf("this is number %v * %v = %v\n", num, w, numm)
	// }

	// for a:= 1; a <= 10; a++ {

	// 	//fmt.Printf("%v * \n", a,b)

	// 	for b := 1; b <= 10; b++ {
	// 		fmt.Printf("%v * %v = \n",a,b)
	// 	}

	// }

	//day one
	// var ThisIsvar = "trial"
	// var a int = 50000
	// var b, c, d int = 4, 5, 6
	// var x, z = 7000, "zreo"
	// n, m := 8000000000, "harsh"

	// fmt.Print(b)
	// fmt.Print(c)
	// fmt.Print(d)

	// fmt.Print(n)
	// fmt.Print(m)

	// fmt.Print(x)
	// fmt.Print(z)

	// fmt.Println("this is and trial")
	// fmt.Println("this is ", ThisIsvar, "and trial")
	// fmt.Println("this is harsh")
	// fmt.Println(a)
	// fmt.Print(a)

	// fmt.Printf("this is trial %v and trial.\n", ThisIsvar)
	// fmt.Printf("this is trial %T and trial.\n", ThisIsvar)

	// const squares = 2
	// var circles = 0

	// fmt.Println("Squares:", squares)
	// fmt.Println("Circles:", circles)

	// //squares=1
	// circles = 7

	// fmt.Println("Squares:", squares)
	// fmt.Println("Circles:", circles)

	/*var num int
	var j int
	fmt.Println("enter the number")
	fmt.Scan(&num)

	for i := 0; i <= 50; i++ {
		j = num * i

		fmt.Printf("%v * %v = %v \n", num, i, j)
	}*/

	/*
		//arr and slices
		var a [5]int
		fmt.Println("emp", a)

		arr := []string{"This", "is", "the", "tutorial", "of", "Go", "language"}
		fmt.Println("arr", arr)

		var my_slice_1 = []string{"Geeks", "for", "Geeks"}
		fmt.Println("this is trial", my_slice_1)
	*/

	greeter()
	greeter()
	greeter()
}

func greeter() {
	fmt.Println("this is harsh")
	for i := 0; i <= 50; i++ {
		fmt.Println(i)
	}
	//return(i)
	//fmt.Print(i)
}
