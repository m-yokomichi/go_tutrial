package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// forのように前処理を記述することができる
	// 前処理で宣言された変数はifのスコープ内が有効
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}