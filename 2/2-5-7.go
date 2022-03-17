package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimRight(text, "\n")
	myRune := []rune(text)

	if unicode.IsUpper(myRune[0]) && string(myRune[len(myRune)-2]) == "." {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
