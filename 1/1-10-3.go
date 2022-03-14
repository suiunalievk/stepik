package main

import "fmt"

func main() {
	var i, n, sum int
	fmt.Scan(&i, &n)
	for ; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
}
