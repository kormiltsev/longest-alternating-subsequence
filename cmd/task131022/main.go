package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var ch = make(chan int)
var dif bool
var compare int

func reader() {
	// fmt.Println("[f] to read from file files/input.txt or input quantity and then data by elements.")
	// scanner := bufio.NewScanner(os.Stdin)
	// if scanner.Text() != "f" {
	// 	i, err := strconv.Atoi(scanner.Text())
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	ch <- i
	// } else {
	inbox := []int{}
	file, err := os.Open("files/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in, _ := strconv.Atoi(scanner.Text())
		inbox = append(inbox, in)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ch <- len(inbox)
	for _, in := range inbox {
		ch <- in
	}
}

func main() {

	outcome := []int{}

	// get data
	go reader()

	// manage by element
	n := <-ch
	fmt.Println("waiting elements: ", n)
	// complete if only 1 element
	if n == 1 {
		outcome = append(outcome, <-ch)
		Writer(outcome)
		return
	}
	// get first element
	compare = <-ch
	outcome = append(outcome, compare)
	fmt.Printf("incoming: %d, i=%d\n", compare, n)
	n--
	// wait for first element not equal to firts one
	for n > 0 {
		x := <-ch
		fmt.Printf("incoming: %d, i=%d\n", x, n)
		n--
		if x-compare > 0 {
			compare = x
			dif = true
			outcome = append(outcome, compare)
			break
		} else if x-compare < 0 {
			dif = false
			outcome = append(outcome, compare)
			break
		}
	}
	// continue
	for n > 0 {
		x := <-ch
		fmt.Printf("incoming: %d, i=%d\n", x, n)
		if x-compare > 0 && !dif {
			compare = x
			dif = true
			outcome = append(outcome, compare)
		} else if x-compare < 0 && dif {
			compare = x
			dif = false
			outcome = append(outcome, compare)
		}
		n--
	}
	// write to file
	Writer(outcome)
	fmt.Println(outcome)
}

func Writer(outcomes []int) {
	file, err := os.OpenFile("files/output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	dwriter := bufio.NewWriter(file)

	for _, data := range outcomes {
		_, err = dwriter.WriteString(strconv.Itoa(data) + "\n")
		if err != nil {
			log.Printf("failed write string: %s", err)
		}
	}

	dwriter.Flush()

}
