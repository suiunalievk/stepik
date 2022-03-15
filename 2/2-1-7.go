package main

import "fmt"

func main() {
	fmt.Print(minimumFromFour())
}

func minimumFromFour() int {
	var min, a int
	for i := 0; i < 4; i++ {
		fmt.Scan(&a)
		if a <= min || min == 0 {
			min = a
		}
	}
	return min
}
