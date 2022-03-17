package main

import (
	"fmt"
)

func main() {
	var text, output string

	fmt.Scan(&text)

	for i := range text {
		if i%2 != 0 {
			output += string(text[i])
		}
	}
	fmt.Println(output)
}
