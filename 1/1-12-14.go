package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	for _, value := range array {
		if a < value {
			a = value
		}
	}
	fmt.Println(a)
}
