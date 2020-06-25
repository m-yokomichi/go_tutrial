package main

import "fmt"

func main() {
	sum := 1
	// ;を抜いてforを使用できる、多言語のwhileの扱い
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}