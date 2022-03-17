package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)
	myRune := []rune(a)
	myRune2 := make([]rune, 0)
	for index := range myRune {
		myRune2 = append(myRune2, myRune[index])

		if index != len(myRune)-1 {
			myRune2 = append(myRune2, '*')
		}
	}

	fmt.Println(string(myRune2))
}
