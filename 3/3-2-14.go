package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	a, _ := strconv.ParseFloat(strings.Split(strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), " ", ""), ",", "."), ";")[0], 64)
	b, _ := strconv.ParseFloat(strings.Split(strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), " ", ""), ",", "."), ";")[1], 64)

	fmt.Println(strconv.FormatFloat(a/b, 'f', 4, 64))
}
