package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	fmt.Print(string(a[2]) + string(a[1]) + string(a[0]))
}
