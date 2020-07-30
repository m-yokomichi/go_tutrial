package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func Walk(t *tree.Tree, ch chan int) {
	var walk func (t *tree.Tree)
	walk = func (t *tree.Tree) {
		if t == nil {
			return 
		}

		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}

	walk(t)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int) , make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for t1Value := range ch1 {
		t2Value := <- ch2
		if t1Value != t2Value {
			return false
		}
	}

	return true
}
func main() {
	ch := make(chan int)
	go Walk(tree.New(3), ch)
	for i := range ch {
		fmt.Println(i)
	}

	var same bool
	same = Same(tree.New(1), tree.New(1))
	fmt.Println(same)

	same = Same(tree.New(1), tree.New(2))
	fmt.Println(same)
}