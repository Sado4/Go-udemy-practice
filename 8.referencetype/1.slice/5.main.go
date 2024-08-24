package main

import "fmt"

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	//fmt.Println("A", "B", "C")
	fmt.Println(Sum(1, 2, 3))

	fmt.Println(Sum(1, 2, 3, 4, 5, 6))

	fmt.Println(Sum())

	sl := []int{10, 20, 30}
	fmt.Println(Sum(sl...))
}
