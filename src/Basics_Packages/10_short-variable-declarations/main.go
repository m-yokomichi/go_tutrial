package main

import "fmt"

// 関数外では暗黙的な宣言は使えない
// hoge := true  # エラー

func main() {
	var i, j int = 1, 2
	// := で暗黙的な型宣言ができる
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}