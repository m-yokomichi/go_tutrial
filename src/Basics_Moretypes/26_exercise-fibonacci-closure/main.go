package main

import (
	"fmt"
)

func main() {
	f := fibonacci()
	for i := 0 ; i < 10 ; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func()int {
	sum := 0
	beforeNum := 0
	count := 0

	return func()int {
		switch count {
			case 0:
				count++
				return sum
			case 1:
				count++
				sum = 1
				return sum
			default :
				tmp := beforeNum
				beforeNum = sum
				sum = sum + tmp
				return sum					
		}
	}
}