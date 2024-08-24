package main

import "fmt"

var i int = 100
var s string = "Golang"

var t, f bool = true, false

var (
	ii int    = 1000
	ss string = "Go"
)

var i2 int

var sss string

//i3 := 100

func main() {
	fmt.Println(i2)
	i2 = 150
	fmt.Println(i2)

	fmt.Println(sss)
	sss = "Go!!"
	fmt.Println(sss)

	i2 = 200
	fmt.Println(i2)

	i3 := 200
	fmt.Println(i3)

	//i3 := 300

	i3 = 300

	//i3 = "string"

}
