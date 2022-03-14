package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	arr := make([]int, a)

	for i := range arr {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr[3])
}
