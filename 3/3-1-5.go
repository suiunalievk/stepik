package main

import "fmt"

func work(x int) int {
	return x
}

func main() {
	var a int
	var m = make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		_, ok := m[a]
		if !ok {
			m[a] = work(a)
		}
		fmt.Printf("%d ", m[a])
	}
}
