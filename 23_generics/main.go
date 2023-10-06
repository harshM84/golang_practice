package main

import "fmt"

/*func has(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}*/ //brfore using generic

func has[t comparable](list []t, value t) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	fmt.Printf("has a %v \n", has([]string{"a", "b"}, "a"))
	fmt.Printf("has c %v \n", has([]string{"a", "b"}, "c"))
	fmt.Printf("has 2 %v \n", has([]int{1, 2, 3}, 2))

}
