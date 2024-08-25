package main

import (
	"fmt"
)

var i5 int = 10

func outer(o int) {
	// var o int = 20
	fmt.Println(o)
}

func main() {
	var num int = 10
	fmt.Println(num)

	var str string = "Hello"
	fmt.Println(str)

	var bo bool = true
	fmt.Println(bo)

	var num2, str2 = 20, "World"
	fmt.Println(num2, str2)

	var (
		num3 int    = 30
		str3 string = "Golang"
	)
	fmt.Println(num3, str3)

	var num4 int
	var str4 string
	var bo2 bool
	fmt.Println(num4, str4, bo2)

	num4 = 40
	str4 = "Hello"
	bo2 = true
	fmt.Println(num4, str4, bo2)

	test := "World"
	test2 := 50
	test3 := true
	fmt.Println(test, test2, test3)

	fmt.Println(i5)

	outer(46)
}
