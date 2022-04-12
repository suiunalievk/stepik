package main

import "fmt"

func main() {

	var a, hours, minutes int

	fmt.Scan(&a) // считаем переменную 'a' с консоли
	hours = a / 30
	minutes = a % 30 * 2
	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
