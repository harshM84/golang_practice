package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 5, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 5 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "animal", "b": "ball"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}
}
