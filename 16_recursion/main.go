package main

import "fmt"

func main() {
	var fac int
	fmt.Println("enter the number")
	fmt.Scanln(&fac)
	//	var i int = 15
	fmt.Printf("factorial of %d is %d ", fac, fact(fac))

	//fmt.Println(countDown(3))
}
func fact(i int) int {
	if i <= 1 {
		return 1
	}
	return i * fact(i-1)
}

// func countDown(count int) int {
// 	if count >= 0 {
// 		countDown(count - 1)
// 	} else {
// 		fmt.Println("count down stop")
// 	}
// 	return count
// }
