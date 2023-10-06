//variables in go

package main

import (
	"fmt"
)

func main() {
	var a = "intial"
	fmt.Println(a)

	var b, c = 1, 2 //multiples variables at once
	fmt.Println(c, b)

	var d = true && false
	fmt.Println(d)

	var h = true
	fmt.Println(h)

	var e int //with out intialization of values
	fmt.Println(e)

	f := "apple" //can also declare like this
	println(f)

	g := 5 //this only returns in inside the function only
	print(g)

	hars := 23

	fmt.Printf("this is the ,%b \n", hars)
	fmt.Printf("this is the ,%X \n", hars)

	y := 40
	z := 40

	fmt.Printf("%v of type %T \n", y, z)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.22
	fmt.Printf("%v of type %T \n", m, m)

}
