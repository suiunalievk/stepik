package main

import "fmt"

func main() {
	fmt.Print(fibonacci(4))
}

func fibonacci(n int) int {
	var fibo, fibo2 int = 0, 1

	for i := 0; i < n; i++ {
		fibo, fibo2 = fibo2, fibo+fibo2
	}
	return fibo
}
