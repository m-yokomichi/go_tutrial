package main

import "fmt"

// 型は変数名の後ろに書く
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}