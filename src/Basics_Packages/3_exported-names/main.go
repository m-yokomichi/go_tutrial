/*
math.Pi のように大文字から始まる名前は外部パッケージから参照できる名前
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}