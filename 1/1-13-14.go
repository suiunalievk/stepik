package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var fibo, fibo2 int = 0, 1
	for i := 1; ; i++ {
		fibo, fibo2 = fibo2, fibo+fibo2

		if a == fibo {
			fmt.Println(i)
			break
		} else if a > fibo && a < fibo2 {
			fmt.Println(-1)
			break
		}
	}
}

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144 и так до бесконечности
