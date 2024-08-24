package main

import "fmt"

func main() {
	var a []int = []int{100, 200}
	a = append(a, 300)
	fmt.Println(a)

	l := make([]int, 0, 5)
	fmt.Println(l)

	fmt.Printf("len=%d cap=%d value=%v\n", len(l), cap(l), l)

	l = append(l, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(l), cap(l), l)

	l = append(l, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(l), cap(l), l)

	e := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(e), cap(e), e)

}
