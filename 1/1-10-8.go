package main

import "fmt"

func main() {
	var i int
	for fmt.Scan(&i); i <= 100; fmt.Scanln(&i) {
		if i >= 10 {
			fmt.Println(i)
		}
	}
}
