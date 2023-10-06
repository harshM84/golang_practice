package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 23
	fmt.Println("this is value", m)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("this is length:", len(m))

	delete(m, "k2")
	fmt.Println("removed", m)

	//can aslo use thsi syntax for map

	n := map[string]int{"foo": 2, "bar": 2}

	fmt.Println("this is new", n)

}
