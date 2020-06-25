package main

import "fmt"

// naked return (naked = 裸)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
	a, b := split(17)
	fmt.Println(a)
	fmt.Println(b)
}