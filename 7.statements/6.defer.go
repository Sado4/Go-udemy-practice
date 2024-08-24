package main

import (
	"fmt"
	"os"
)

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func main() {
	TestDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	file.Write([]byte("Hello\n"))

	defer func() {
		fmt.Println("Func A")
		fmt.Println("Func B")
		fmt.Println("Func C")
	}()
}
