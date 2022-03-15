package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	fmt.Println(1 + ((a - 1) % 9))
}
