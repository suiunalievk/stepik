package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, count int
	var max = math.MaxInt64
	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b < max {
			max = b
			count = 1
		} else if b == max {
			count++
		}
	}

	fmt.Printf("%d", count)
}
