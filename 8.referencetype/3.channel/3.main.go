package main

import (
	"fmt"
	"time"
)

func reciever2(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)

	close(ch1)

	//ch1 <- 1 //エラー

	fmt.Println(<-ch1)

	ii, ok := <-ch1
	fmt.Println(ii, ok)

	ch2 := make(chan int, 20)

	go reciever2("1.goroutin", ch2)
	go reciever2("2.goroutin", ch2)
	go reciever2("3.goroutin", ch2)

	iii := 0
	for iii < 100 {
		ch2 <- iii
		iii++
	}
	close(ch2)
	time.Sleep(3 * time.Second)

}
