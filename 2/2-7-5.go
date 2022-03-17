package main

import (
	"fmt"
)

func main() {
	var a, b string
	var c int32
	fmt.Scan(&a)

	for _, val := range a {
		c = (val - 48) * (val - 48)
	}
	fmt.Println(c, " ", b)
}
