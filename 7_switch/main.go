package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 2
	var i int
	fmt.Println("enter the number")
	fmt.Scanln(&i)
	// fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("entered invalid number")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is weekend")
	default:
		fmt.Println("it is weekday")
	}
}
