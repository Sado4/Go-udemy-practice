package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("表示")

	print("Hello")
	println("Hello!")

	fmt.Print("Hello")
	fmt.Println("Hello!")

	fmt.Println("Hello", "world!")
	fmt.Println("Hello", "world!")

	fmt.Printf("%s\n", "Hello")
	fmt.Printf("%#v\n", "Hello")

	fmt.Sprint("Hello")
	fmt.Sprintf("Hello")
	fmt.Sprintln("Hello")

	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")

}
