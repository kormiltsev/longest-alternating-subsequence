package main

import (
	"fmt"
	"math/rand"
)

var n = 1000
var size = 1000000000

func main() {
	fmt.Println(n)
	for i := 0; i < n; i++ {
		fmt.Println(rand.Intn(size))
	}
}
