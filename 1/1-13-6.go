package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if ((a+b)/2)*10%10 == 0 {
		fmt.Printf("%d", (a+b)/2)
	} else {
		fmt.Printf("%.1f", float32(a+b)/2)
	}
}
