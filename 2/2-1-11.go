package main

import "fmt"

func main() {
	fmt.Print(sumInt(1, 2, 3))
}

func sumInt(a ...int) (int, int) {
	var b int
	for _, value := range a {
		b += value
	}
	return len(a), b
}
