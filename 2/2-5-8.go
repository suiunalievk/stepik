package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text := bufio.NewScanner(os.Stdin)
	text.Scan()
	str := text.Text()
	myRune := []rune(str)
	len := len(myRune) - 1
	for index := range myRune {
		if myRune[index] != myRune[len-index] {
			fmt.Println("Нет")
			return
		}
	}
	fmt.Println("Палиндром")
}
