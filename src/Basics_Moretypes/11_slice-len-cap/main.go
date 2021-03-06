package main

import (
	"fmt"
)

func main() {
	s := []int{1, 3, 5, 8, 10, 15}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	// 0 , 1 が削ぎ落とされる
	s = s[2:]
	printSlice(s)

	s = s[0:4]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}