package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var prev float64
	z := float64(1)
	for math.Abs(z-prev) > 0.01 {
		prev = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
