package main

import (
	"fmt"
)

func main() {
	// ループ用のスライス
	pow := make([]int, 10)

	for i := range pow {
		fmt.Printf("i = %v\n", i)
		pow[i] = 1 << uint(i) // 2**i
	}

	for _, value := range pow {
		fmt.Printf("value = %v\n", value)
	}
}