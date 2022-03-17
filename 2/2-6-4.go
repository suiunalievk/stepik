package main

import "fmt"

func divide(a, b int) (int, string) {
	if b == 0 {
		return 0, "division by zero!"
	}
	return a / b, ""
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	res, _ := divide(a, b)

	//if err != nil {
	//	fmt.Println("ошибка")
	//	return
	//}
	fmt.Println(res)
}
