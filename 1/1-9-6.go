package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a)
	b = a / 100
	c = a / 10 % 10
	d = a % 10
	if b != c && c != d && d != b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
