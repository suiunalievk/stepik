package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	arr := make([]int, a)

	for i := range arr {
		fmt.Scan(&arr[i])

		if arr[i] > 0 {
			a++
		}
	}
	fmt.Printf("%d", a)
}
