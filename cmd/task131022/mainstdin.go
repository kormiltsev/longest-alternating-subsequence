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
var ch = make(chan int, 10)

// reader for standard input
func readerdir() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Panic(err)
		}
		ch <- i
	}
}

func main() {
	start := time.Now() // count working time
	// get data
	go readerdir()
	n := <-ch
	// complete if only 1 element
	if n <= 1 {
		fmt.Println(strconv.Itoa(<-ch))
		return
	}
	buffer := []int{}
	dif := false
	i := 0
	m := <-ch
	l := n - 1
	fmt.Println(m)
	// for the first dif
	for l > 0 {
		i = <-ch
		l--
		if i > m {
			dif = true
			buffer = []int{i}
			break
		} else if i < m {
			dif = false
			buffer = []int{i}
			break
		}
	}
	for l > 0 {
		i = <-ch
		l--
		if i > m && !dif {
			for k := len(buffer); k > 0; k-- {
				if buffer[k-1] >= i {
					break
				}
				m = buffer[k-1]
			}
			fmt.Println(m)
			buffer = []int{i}
			dif = !dif
		}
		if i < m && dif {
			for k := len(buffer); k > 0; k-- {
				if buffer[k-1] >= i {
					break
				}
				m = buffer[k-1]
			}
			fmt.Println(m)
			buffer = []int{i}
			dif = !dif
		}
		if i != m {
			buffer = append(buffer, i)
			m = i
		}
	}
	duration2 := time.Since(start) // count working time

	fmt.Println("Finished!\ndone in ", duration2)
}
