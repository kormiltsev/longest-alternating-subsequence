package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var mas = []int{}

func main() {
	// get data
	income := []int{1, 7, 5, 7, 4, 6, 8, 3, 5, 3, 7, 7, 9, 4, 5, 11}
	n := len(income)
	outcome := []int{}

	// corner case
	if n < 3 {
		outcome = income
		fmt.Println(outcome)
		return
	}

	compare := income[0]
	var dif bool

	for len(outcome) == 0 {
		for _, value := range income {
			if value-compare < 0 {
				// for chanels change dif to opposite
				dif = true
				outcome = append(outcome, compare)
				// for chanels add value to outcome
				break
			}
			if value-compare > 0 {
				// for chanels change dif to opposite
				dif = false
				outcome = append(outcome, compare)
				// for chanels add value to outcome
				break
			}
		}
	}

	for _, value := range income {
		if dif && (value-compare) < 0 {
			outcome = append(outcome, value)
			compare = value
			dif = !dif
		}
		if !dif && (value-compare) > 0 {
			outcome = append(outcome, value)
			compare = value
			dif = !dif
		}
	}

	fmt.Println(outcome)
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
