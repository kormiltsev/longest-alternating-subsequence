package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var n = 1000

func main() {
	err := os.Remove("files/input.txt")
	if err != nil {
		log.Println(err)
	}

	set := []int{}
	set = append(set, n)
	for i := 0; i < n; i++ {
		set = append(set, rand.Intn(1000000000))
	}

	file, err := os.OpenFile("files/input.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	dwriter := bufio.NewWriter(file)

	for _, data := range set {
		_, err = dwriter.WriteString(strconv.Itoa(data) + "\n")
		if err != nil {
			log.Printf("failed write string: %s", err)
		}
	}

	dwriter.Flush()
}
