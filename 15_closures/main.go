package main

import (
	"fmt"
	//"strconv"
)

func main() {
	// func(name string) {
	// 	fmt.Printf("aur %s bhai ,all good", name)
	// }("toofani")

	// x := "1111"
	// value, _ := strconv.ParseInt(x, 10, 64)
	// fmt.Println(value)
	nextInt := intseq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intseq()
	fmt.Println(newInt())
}
func intseq() func() int {
	i := 0
	fmt.Println("outer func exrcuted and retured another function,with intial value", i)
	return func() int {
		fmt.Print("inner function executed with variable scope from other function")
		i++
		return i
	}
}
