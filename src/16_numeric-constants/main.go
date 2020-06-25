package main

import "fmt"

// 型のない定数はその状況によって必要な型をとる
const (
	// 2進数で 100桁左にシフトする 1000000.....0(100個)
	Big = 1 << 100
	// 99桁右にシフトする
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}