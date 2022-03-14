package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println((a/100 + a/10%10 + a%10))
}
