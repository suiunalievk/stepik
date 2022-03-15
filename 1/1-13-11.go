package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)

	for fmt.Scan(&b); b >= a; b-- {
		if b%7 == 0 {
			fmt.Printf("%d", b)
			break
		} else if b == a {
			fmt.Println("NO")
		}
	}
}
