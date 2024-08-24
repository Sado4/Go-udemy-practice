package main

import "fmt"

func main() {

	var i int = 100

	fmt.Println(i + 50)

	var ii int64 = 50

	//fmt.Println(i + ii)

	fmt.Println(int32(ii))
	fmt.Printf("%T\n", int32(ii))
}
