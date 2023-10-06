package main

import "fmt"

func main() {
	var ptr *int

	fmt.Printf("null pointer %x \n", ptr)
	//////////////////////////////////////////////////////////////
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
	///////////////////////////////////////////////////////

	var a int = 30 //actual variable declration
	var ptrr *int  //pointer variable declaraton

	ptrr = &a //store address of pointer variable

	fmt.Printf("address of veriable %x \n", &a)

	//address stored in pointer veriable
	fmt.Printf("address stored in ptrr variable %x \n", ptrr)

	//access the value using pointer veriable
	fmt.Printf("value of *ptrr is %d \n", *ptrr)
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}
