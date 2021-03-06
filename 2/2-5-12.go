package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var text, output string
	output = "Ok"

	scan, err := fmt.Scan(&text)
	if err != nil {
		return
	}

	myRune := []rune(text)

	if utf8.RuneCountInString(text) < 5 {
		output = "Wrong password"
	}
	for _, value := range myRune {
		if !unicode.In(value, unicode.Latin) && !unicode.IsNumber(value) {
			output = "Wrong password"
		}
	}

	fmt.Println(output)
}
