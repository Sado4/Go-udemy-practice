package main

import (
	"fmt"
)

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := later()

	fmt.Println(f("Golang"))   // => ""
	fmt.Println(f("is"))       // => "Golang"
	fmt.Println(f("awesome!")) // => "is"
}
