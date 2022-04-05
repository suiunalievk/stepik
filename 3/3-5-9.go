package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var a, res int
	for scanner.Scan() {
		a, _ = strconv.Atoi(scanner.Text())
		res += a
	}

	io.WriteString(os.Stdout, strconv.Itoa(res))
}
