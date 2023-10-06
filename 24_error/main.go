// https://www.youtube.com/watch?v=gZmeeHM_1Uc
// advance error handlaning(go guru)
package main

import (
	"errors"
	"fmt"
)

func sum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("value is can't be zero or less then zero")
	}
	return a + b, nil
}

func main() {
	c, err := sum(10, -1)
	if err != nil {
		fmt.Println(err.Error()) //err is like ineface type and Error is like method and retuen the string
		return
	}
	fmt.Println(c)
}
