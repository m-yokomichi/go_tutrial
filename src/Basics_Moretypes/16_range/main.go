package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main()  {
	// foreach的なやつ
	for i, v := range pow {
		fmt.Println(i, v, pow)
	}
}