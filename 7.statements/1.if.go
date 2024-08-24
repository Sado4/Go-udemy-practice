package main

import "fmt"

func main() {
	a := 1
	//if  else if  else
	if a == 2 {
		fmt.Println("2だ")
	} else if a == 1 {
		fmt.Println("1だ")
	} else {
		fmt.Println("それ以外")
	}

	if b := 100; b == 100 {
		fmt.Println("100だ")
	}

	x := 5
	if x := 2; true {
		fmt.Println(x)
		//2
	}

}
