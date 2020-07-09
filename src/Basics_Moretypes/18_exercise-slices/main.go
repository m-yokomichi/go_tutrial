package main

import "golang.org/x/tour/pic"

// 長さdyのsliceに、各要素が8bitのuint型で長さdxのsliceを割り当てたものを返す
func Pic(dy, dx int) [][]uint8  {
	// 長さdyのsliceを作成
	pic := make([][]uint8, dy)
	for y := range pic {
		// 長さdxの []int8を作成
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			// (x+y)/2 、 x*y 、 x^y などがある
			pic[y][x] = uint8(x * y)
		}
	}

	return pic
}
func main()  {
	pic.Show(Pic)
}