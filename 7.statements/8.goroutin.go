package main

import (
	"fmt"
)

func sub() {
	i := 0
	for {
		fmt.Println(i)
		i++
	}
}

func main() {
	go sub()
	/*
		go func() {
			i := 0
			for {
				fmt.Println(i)
				i++
			}
		}()
	*/
	for {
		fmt.Println("main")
	}

	/*
		go fmt.Println("追加")
		fmt.Printf("CPU数: %d\n", runtime.NumCPU())
		fmt.Printf("Goroutin数: %d\n", runtime.NumGoroutine())
		fmt.Printf("Version: %d\n", runtime.Version())
	*/

}
