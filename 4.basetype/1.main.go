package main

import "fmt"

func main() {
	var s string = "Hello Golang"

	si := "100"

	fmt.Println(s)
	fmt.Println(si)

	fmt.Printf("%T\n", s)

	fmt.Println(`Hello
		World`)

	fmt.Println("\"")
	fmt.Println(`"`)

	//fmt.Println(s[0])
}
