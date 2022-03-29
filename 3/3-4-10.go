package main

import (
	"fmt"
)

func readTask() (value1, value2, operation interface{}) {
	return 5.0, 0.1, "+"
}

func CheckFloat(a ...interface{}) bool {
	for _, k := range a {
		if _, ok := k.(float64); !ok {
			fmt.Printf("value=%v: %T", k, k)
			return false
		}
	}
	return true
}

func main() {
	value1, value2, operation := readTask()
	var res float64
	if CheckFloat(value1, value2) {
		switch operation.(type) {
		case string:
			switch operation {
			case "+":
				res = value1.(float64) + value2.(float64)
			case "-":
				res = value1.(float64) - value2.(float64)
			case "*":
				res = value1.(float64) * value2.(float64)
			case "/":
				res = value1.(float64) / value2.(float64)
			default:
				fmt.Print("неизвестная операция")
				return
			}
		default:
			fmt.Print("неизвестная операция")
			return
		}
		fmt.Printf("%.4f", res)
	}
}
