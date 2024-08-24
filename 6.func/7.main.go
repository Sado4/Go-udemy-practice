package main

import (
	"fmt"
)

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()

	fmt.Println(ints()) // => "1"
	fmt.Println(ints()) // => "2"
	fmt.Println(ints()) // => "3"

	otherInts := integers()
	fmt.Println(otherInts()) // => "1"
}
