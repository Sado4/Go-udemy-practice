package main

import "fmt"

func main() {
	fmt.Println("A")
	goto L
	fmt.Println("B")
L:
	fmt.Println("C")

	for {
		for {
			for {
				fmt.Println("Goto")
				goto D
			}
		}
	}
D:
	fmt.Println("end")

}
