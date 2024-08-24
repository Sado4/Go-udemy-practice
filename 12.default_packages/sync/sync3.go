package main

import (
	"fmt"
)

func main() {
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
	}()
	/*
		for {

		}
	*/
}
