package main

import "fmt"

func main() {
	var a float32
	fmt.Scanln(&a)
	if a <= 0 {
		fmt.Printf("число %2.2f не подходит", a)
	} else if a > 10000 {
		fmt.Printf("%e", a)
	} else {
		fmt.Printf("%.4f", a*a)
	}
}
