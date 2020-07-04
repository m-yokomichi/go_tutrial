package main

import "fmt"

func main()  {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// データではなく元の配列部分を示している(ポインタ)
	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)

	// b[0](Paulの部分)を書き換えると names[1]が書き換えられる
	b[0] = "HogeHoge"
	fmt.Println(a, b)
	fmt.Println(names)

}