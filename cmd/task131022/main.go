package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var dif bool
var inbox = []int{}
var outcome = []int{}

func reader() {

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

}

func main() {

	// get data
	reader()
	fmt.Println(inbox)
	n := inbox[0]
	// complete if only 1 element
	if n <= 1 {
		return
	}
	dif := false
	i := 2
	m := 1
	newi := 0
	//outcome = append(outcome, inbox[1])
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
			fmt.Printf("i=%d, m=%d\n", i, m)
			dif = !dif
			newi = 0
		}
		i++
	}

	// write to file
	fmt.Println("outcome: ", outcome)
	Writer(outcome)
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
