package main

import (
	"fmt"
)

func main()  {
	s := []int{2,4,11,3,23,56}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

}