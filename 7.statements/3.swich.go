package main

import (
	"fmt"
)

func main() {
	n := 4
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("どれでもない")
	}

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("どれでもない")
	}

	n2 := 6
	switch {
	case n2 > 0 && n2 < 4:
		fmt.Println("0 < n2 < 4")
	case n2 > 3 && n2 < 6:
		fmt.Println("3 < n2 < 6")
	}
	/*
		switch n := 2; n {
		case 1, 2, 3:
			fmt.Println("0 < n2 < 4")
		case n > 3 && n < 6:
			fmt.Println("3 < n2 < 6")
		}
	*/

}
