package main

import (
	"fmt"
	"strconv"
)

func main() {
	fl64 := 10.2

	n := int(fl64)
	nn := float64(n)

	var s string = "100"

	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)

	h := "Hello World"
	fmt.Println(string(h[0]))

	var ii int = 321
	var ss string
	ss = strconv.Itoa(ii)
	fmt.Println(ss)

}
