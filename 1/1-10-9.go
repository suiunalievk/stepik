package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	for i := 1; a <= c; i++ {
		a += a * b / 100
		if a >= c {
			fmt.Println(i)
			break
		}
	}
}
