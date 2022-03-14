package main

import "fmt"

func main() {
	var a, b uint8
	var workArray [10]uint8

	for i := range workArray {
		fmt.Scan(&workArray[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Scan(&a, &b)
		workArray[a], workArray[b] = workArray[b], workArray[a]
	}

	for _, value := range workArray {
		fmt.Printf("%d ", value)
	}
}
