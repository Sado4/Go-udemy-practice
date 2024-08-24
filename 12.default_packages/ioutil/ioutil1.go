package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Open("foo.txt")
	bs1, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs1))

	bs2, _ := ioutil.ReadFile("foo.txt")
	fmt.Println(string(bs2))

}
