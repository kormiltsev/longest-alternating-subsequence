package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var dif bool
var inbox = []int{}
var outcome = []int{}
var filename = "files/input.txt"

// reader for standard input
func readerdir() {
	fmt.Println("Enter file name to upload from file or use pipe: go run ./files/generatorstdin.go | go run ./cmd/task131022/main.go")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			reader(scanner.Text())
			return
		}
		inbox = append(inbox, i)
	}
}

// reader from file
func reader(filename string) {

	file, err := os.Open(filename)
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

}

func main() {
	start := time.Now() // count working time
	// get data
	readerdir()
	duration1 := time.Since(start) // count working time
	n := inbox[0]
	fmt.Println("Open and read file done or STDIN got.  N =", n)
	// complete if only 1 element
	if n <= 1 {
		return
	}
	dif := false
	i := 2
	m := 1
	newi := 0
	for i <= n {
		for k := i - 1; k >= m; k-- {
			if inbox[i]-inbox[k] < 0 && dif {
				newi = k
			} else if inbox[i]-inbox[k] > 0 && !dif {
				newi = k
			} else {
				break
			}
		}
		if newi != 0 {
			m = i
			outcome = append(outcome, inbox[newi])
			dif = !dif
			newi = 0
		}
		i++
	}
	outcome = append(outcome, inbox[m])
	duration2 := time.Since(start) // count working time
	// write to file
	writer(outcome) // to use stdout need to comment all "fmt.*" lines and uncomment next line:
	//writeout(outcome)
	duration3 := time.Since(start) // count working time
	fmt.Println("open and read file:", duration1, "\nprogramm finished (sinse start)", duration2, "\nwrite to file complete (sinse start)", duration3, "\n")
}

func writer(outcomes []int) {
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

func writeout(outcomes []int) {
	for _, val := range outcomes {
		fmt.Println(val)
	}
}
