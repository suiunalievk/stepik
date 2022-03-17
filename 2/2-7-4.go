package main

import (
	"fmt"
)

func main() {
	var a string
	var b byte
	fmt.Scan(&a)

	for i := range a {
		if a[i] > b {
			b = a[i]
		}
	}
	fmt.Println(int(b - '0'))
}
