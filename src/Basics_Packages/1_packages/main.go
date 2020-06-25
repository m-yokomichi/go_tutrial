// mainパッケージから開始される
// golang.org/x/net/websocket とかの場合は websocketがパッケージ名になる
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
