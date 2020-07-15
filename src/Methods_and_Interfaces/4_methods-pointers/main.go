package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func main() {
	v := Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v.Abs())
}

func (v Vertex) Abs()float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// ポインタレシーバ
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}