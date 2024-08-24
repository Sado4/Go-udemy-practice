package main

import "fmt"

func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))

}
