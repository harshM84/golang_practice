package main

import "fmt"

func main() {
	/*
		var arr1 =[5]int{} //not initialized
		var arr2 =[4]int{1,2} //patialy initialize
		var arr3 =[3]int {1,2,3}//fully initialize
		var arr3 =[...]int {1,2,3}// hare size of array is not declare & we can also
		//wite like this

	*/
	//static array
	var arry = [3]string{"harsh", "is", "good"}
	var arry2 = [...]string{"this", "is", "harsh"}
	fmt.Println(arry)
	fmt.Println(arry2)

	//dynamic array

	var arry1 = [5]
	fmt.Println("enter the number")

	for i := 0; i <= 4; i++ {
		fmt.Scan(&arry1)
	}

	for z := 0; z <= 4; z++ {
		fmt.Printf("this is your array %v \n", arry1)
	}

}
