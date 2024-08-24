package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			defer wg.Done()
			fmt.Println("process", i*1000)
		}()

	}
	fmt.Println("###################")
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	go consumer(ch, &wg)

	wg.Wait()
	close(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
