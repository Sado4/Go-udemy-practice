package main

import "fmt"

func main() {
	ch3 := make(chan int, 4)
	ch3 <- 5
	ch3 <- 7
	ch3 <- 9
	close(ch3)
	for ci := range ch3 {
		fmt.Println(ci)
	}
}
