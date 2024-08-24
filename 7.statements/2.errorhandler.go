package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "100"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T %v\n", i, i)
}
