package main

import (
	"fmt"
	"strings"
)

func main() {
	var text, str string

	fmt.Scan(&text, &str)

	fmt.Println(strings.Index(text, str))
}
