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
var i, l, m int
var ch = make(chan int, 10)
var buffer = []int{}

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

func Next() {
	for l > 0 {
		i = <-ch
		l--
		if i > m && !dif {
			for k := len(buffer); k > 0; k-- {
				if buffer[k-1] >= i {
					break
				}
				m = buffer[k-1]
				buffer = buffer[:k-1]
			}
			fmt.Println(m)
			buffer = []int{i}
			dif = !dif
		}
		if i < m && dif {
			for k := len(buffer); k > 0; k-- {
				if buffer[k-1] <= i {
					break
				}
				m = buffer[k-1]
				buffer = buffer[:k-1]
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
	//dif := false
	i = 0
	buffer = []int{<-ch}
	l = n - 1
	fmt.Println(buffer[0])
	// for the first dif
	for l > 0 {
		i = <-ch
		l--
		if i > buffer[0] {
			dif = true
			break
		} else if i < buffer[0] {
			dif = false
			break
		}
	}
	m = i
	buffer = []int{i}
	// next
	Next()
	// last
	fmt.Println(buffer[0])
	duration2 := time.Since(start) // count working time
	fmt.Println("Finished!\ndone in ", duration2)
}
