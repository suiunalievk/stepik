package main

import "fmt"

func main() {
	var a = 8
	const b int = 10
	a = a + b
	//b = b + a
	fmt.Println(a)
}
