package main

import "fmt"

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	ffl64 := 3.2
	fmt.Printf("%T\n", ffl64)

	var fl32 float32 = 5.4

	fmt.Printf("%T\n", float32(ffl64))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	//var u8 uint = 255 byte

	var c64 complex64 = -5 + 12i

}
