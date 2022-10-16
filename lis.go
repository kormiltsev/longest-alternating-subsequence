package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var inbox = []int{}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("not int")
		}
		inbox = append(inbox, i)
	}

	fmt.Println(inbox)
}
