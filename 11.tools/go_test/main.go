package main

import (
	"fmt"
	//"./alib"
	"kiso_blog/go_test/alib"
)

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Println(alib.Average(s))
}
