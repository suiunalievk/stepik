package main

import "fmt"

func main() {
	var i, max, count int
	for fmt.Scan(&i); i != 0; fmt.Scan(&i) {
		if i > max {
			max = i
			count = 1
		} else if max == i {
			count++
		}
	}
	fmt.Println(count)
}
