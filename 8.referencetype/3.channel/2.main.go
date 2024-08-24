package main

import (
	"fmt"
	"time"
)

func reciever1(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	var ch1 chan int
	var ch2 chan int
	ch1 = make(chan int)
	ch2 = make(chan int)

	ch1 <- 1

	go reciever1(ch1)
	go reciever1(ch2)

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		//
		time.Sleep(50 * time.Millisecond)
		i++
	}
}
