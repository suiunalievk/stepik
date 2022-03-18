package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)

	for _, val := range a {
		fmt.Print((val - 48) * (val - 48))
	}
}
