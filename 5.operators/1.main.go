package main

import "fmt"

func main() {
	fmt.Println(1 + 2)
	fmt.Println("ABC" + "DE")

	fmt.Println(3 - 2)

	fmt.Println(5 * 6)

	fmt.Println(9 / 3)

	fmt.Println(9 / 4)

	n := 0
	n += 4
	fmt.Println(n)
	n++
	fmt.Println(n)
	n--
	fmt.Println(n)

	s := "演算子"
	s += "の解説"
	fmt.Println(s)
}
