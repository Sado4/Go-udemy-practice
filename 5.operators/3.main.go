package main

import "fmt"

func main() {
	fmt.Println(true && false == true)
	fmt.Println(true || false == true)

	fmt.Println(!true)
	fmt.Println(!false)
}
