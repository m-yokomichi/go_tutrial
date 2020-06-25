package main

import (
	"fmt"
	"math"
)

// 型()で型変換を行う
// goでは明示的な変換が必要
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z  = uint(f)
	fmt.Println(x, y, z)
}
