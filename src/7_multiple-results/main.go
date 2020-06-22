package main

import "fmt"

func swap(x, y string) (string, string) {
	// 複数の戻り値を返すことができる
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
