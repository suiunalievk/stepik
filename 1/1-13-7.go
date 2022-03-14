package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b == 0 {
			c++
		}

	}
	fmt.Printf("%d", c)
}
