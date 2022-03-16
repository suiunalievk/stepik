package main

import "fmt"

func main() {
	var a, b = 2, 3
	fmt.Println(a, " ", b)
	test(&a, &b)

}

func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Print(*x1, " ", *x2)
}
