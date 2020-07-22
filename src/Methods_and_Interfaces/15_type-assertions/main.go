package main

import (
	"fmt"
)

func main() {
	var i interface{} = "Hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// panicが発生する
	f = i.(float64)
	fmt.Println(f)
}