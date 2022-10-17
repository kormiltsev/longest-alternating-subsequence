package main

import (
	"fmt"
	"math/rand"
)

var n = 1000
var size = 1000000000
var st = []int{}

func randomer() {
	for i := 0; i < n; i++ {
		st = append(st, rand.Intn(size))
	}
}

func main() {
	randomer() //v2
	fmt.Println(n)
	// v2 work without time spending on rand func
	for _, v := range st {
		fmt.Println(v)
	}
	//-----------------------------------------

	// v1
	//for i := 0; i < n; i++ {
	//fmt.Println(rand.Intn(size)) //rand is spend a lot of time
	// fmt.Println("103030303")  //used for testing
	// fmt.Println("203030303")
	// fmt.Println("303030303")
	// fmt.Println("103030303")
	//time.Sleep(time.Millisecond * 5) // also work with delay
	//}
}
