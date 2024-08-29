package main

import "fmt"

const IntVal = 1

func main() {
	const Pi = 3.14
	//const Pi = 3

	const zero = 0.0 // untyped floating-point constant

	const (
		URL      = "http://xxx.jp"
		SiteName = "test"
	)

	const (
		A = 1
		B
		C
		D = "A"
		E
		F
	)

	// var big int = 9223372036854775807 + 1

	const big = 9223372036854775807 + 1

	fmt.Println(Pi, URL, SiteName)
	fmt.Println(big - 1)

	const (
		c0 = 1    // c0 == 0
		c1 = iota // c1 == 1
		c2 = iota // c2 == 2
	)

	fmt.Println(IntVal)
	fmt.Println(c0, c1, c2)
	fmt.Println(A, B, C, D, E, F)
}
