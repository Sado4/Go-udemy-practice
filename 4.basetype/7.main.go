package main

import "fmt"

func main() {
	var x interface{}
	fmt.Println(x)

	x = 1
	x = 3.14
	x = "A"
	x = [...]int{3, 4, 5, 6}
	//x = 2
	//var y interface{} = 3
	//fmt.Println(x + y)
}
