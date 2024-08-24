package main

import (
	"fmt"
)

func callFunction(f func()) {
	f()
}

func main() {
	callFunction(func() {
		fmt.Println("I'm a function")
	}) // => "I'm a function"
}
