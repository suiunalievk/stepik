package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(n uint) uint {
		var res = ""
		var res2 uint64
		str := strconv.Itoa(int(n))

		for i := range str {
			if str[i]%2 == 0 && str[i] != 48 {
				res += string(str[i])
			}
		}

		if res2, _ = strconv.ParseUint(res, 10, 32); res2 == 0 {
			res2 = 100
		}

		return uint(res2)
	}

	fmt.Println(fn(200))
}
