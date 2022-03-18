package main

import (
	"fmt"
	"math"
)

var p, v, k float64 = 6.0, 6.0, 1296.0

func T() float64 {
	return 6 / W()
}

func W() float64 {
	return math.Sqrt(k / M())
}

func M() float64 {
	return p * v
}

func main() {
	T()
	fmt.Println(T())
}
