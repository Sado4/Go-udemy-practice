package main

import (
	"fmt"
)

func main() {
	ch4 := make(chan int, 30)
	ch5 := make(chan string, 20)
	ch5 <- "AAA"
	//e1 := <-ch4
	//e2 := <-ch5
	//fmt.Println(e2)
	//fmt.Println(e1)

	select {
	case e1 := <-ch4:
		fmt.Println(e1 + 1000)
	case e2 := <-ch5:
		fmt.Println(e2 + "e1")
	default:
		fmt.Println("どちらでもない")
	}

	ch6 := make(chan string, 5)
	ch7 := make(chan string, 5)
	ch8 := make(chan string, 5)

	ch6 <- "A"
	ch7 <- "B"
	ch8 <- "C"

	select {
	case s1 := <-ch6:
		fmt.Println(s1)
	case s2 := <-ch7:
		fmt.Println(s2)
	case s3 := <-ch8:
		fmt.Println(s3)
	}

	ch9 := make(chan int, 5)
	ch10 := make(chan int, 5)
	ch11 := make(chan int, 5)

	go func() {
		for {
			i1 := <-ch9
			ch10 <- i1 * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch10
			ch11 <- i2 - 1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch9 <- n:
			n++
		case i3 := <-ch11:
			fmt.Println("recieve", i3)
		default:
			if n > 100 {
				//close(ch9)
				break L
			}
		}
	}
}
