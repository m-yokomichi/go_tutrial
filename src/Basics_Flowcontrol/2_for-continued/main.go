package main

import "fmt"

func main() {
	sum := 1
	// 初期値と後処理は任意 （他の言語と同様）
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}