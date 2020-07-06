package main

import (
	"fmt"
)

func main()  {
	q := []int{1, 3, 5, 7}

	fmt.Println(q)

	r := []bool{true, false, true, true, false, false}
	fmt.Println(r)

	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{1, false},
		{20, true},
		{5, false},
		{13, false},
		{78, true},
	}

	fmt.Println(s)
}
