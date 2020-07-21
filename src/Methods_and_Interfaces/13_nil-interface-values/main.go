package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// 具体的なM()の実装がないため下記はエラー
	//i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}