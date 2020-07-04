package main

import "fmt"

func main()  {
	primes := [6]int{2,4,6,7,8, 1}

	// スライス
	var s []int = primes[1:4]
	fmt.Println(s)
}