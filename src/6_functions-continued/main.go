package main

import "fmt"

// ２つ以上の引数が同じ型の場合は最後の方を残して省略して記述できる
func add(x, y , z int) int {
	return x + y + z
}

func main() {
	fmt.Println(add(42, 13, 1))
}