package main

import "fmt"

func main() {
	var i, n, k, sum int
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&k)
		if k > 9 && k < 100 && k%8 == 0 {
			sum += k
		}
	}
	fmt.Println(sum)
}
