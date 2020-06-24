package main

import "fmt"

// 変数 := 値 で右側の値から型推論される
func main() {
	v := 0.5i // change me!
	fmt.Printf("v is of type %T\n", v)
}