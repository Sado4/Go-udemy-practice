package main

import "fmt"

func main() {
	var i []int
	fmt.Println(cap(i))

	var a []int = []int{100, 200}

	b := []string{"A", "B"}
	fmt.Println(b)

	l := make([]int, 0, 5)
	fmt.Println(l)

	a[0] = 1000
	fmt.Println(a)

	d := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(d)

	fmt.Println(d[0])

	fmt.Println(d[2:4])

	fmt.Println(d[:2])

	fmt.Println(d[2:])

	fmt.Println(d[:])

	fmt.Println(d[len(d)-1])

	fmt.Println(d[1 : len(d)-1])

}
