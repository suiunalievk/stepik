package main

import "fmt"

func main() {
	fmt.Print(vote(0, 1, 1))
}

func vote(x, y, z int) int {
	return (x + y + z) / 2
}
