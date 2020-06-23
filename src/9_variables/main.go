package main

import "fmt"

// var で変数を宣言
// 引数と同様に最後に肩を書く
var c, python, java bool

func main() {
	var i int
	// c,python,java はfalseが返る
	// i は 0 が返る
	fmt.Println(i, c, python, java)
}