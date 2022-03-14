package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	arr := make([]int, a)

	for i := range arr {
		fmt.Scan(&arr[i])

		if i%2 == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}
}
