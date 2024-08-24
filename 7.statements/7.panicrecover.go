package main

import "fmt"

func main() {
	panic("rungime error")
	fmt.Println("START")

	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("START")

}
