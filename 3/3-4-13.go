package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	Input string
}

func (b *Battery) String() string {
	cnt := strings.Count(b.Input, "1")
	if cnt > 10 {
		cnt = 10
	}

	return fmt.Sprintf("[%10s]", strings.Repeat("X", cnt))
}

func main() {
	batteryForTest := new(Battery)
	fmt.Scan(&batteryForTest.Input)

	fmt.Println(batteryForTest)
}
