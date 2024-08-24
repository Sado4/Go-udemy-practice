package main

import "fmt"

func main() {
	var a []int = []int{100, 200}

	d := []int{1, 2, 3, 4, 5, 6}

	copy := copy(d, a)
	fmt.Println(copy)
	fmt.Println(d)
}
