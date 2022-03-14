package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Printf("It is %d hours %d minutes.", a/3600, a/60%60)
}
