package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// switchに条件書いていない場合は switch (true) と書くと同じ
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}