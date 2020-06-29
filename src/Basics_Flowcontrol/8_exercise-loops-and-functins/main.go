package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	// お手製
	fmt.Println(Sqrt(2))
	// math.Sqrt
	fmt.Println(math.Sqrt(2))
}