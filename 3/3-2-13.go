package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func cleaner(s string) string {
	tmp := make([]rune, 0)

	for _, r := range []rune(s) {
		if unicode.IsDigit(r) {
			tmp = append(tmp, r)
		}
	}

	return string(tmp)
}

func adding(a, b string) int64 {
	var res, res2 int64

	res, _ = strconv.ParseInt(cleaner(a), 10, 64)
	res2, _ = strconv.ParseInt(cleaner(b), 10, 64)
	return res + res2
}

func main() {
	var a, b string = "%^80", "hhhhh20&&&&nd"

	fmt.Println(adding(a, b))

}
