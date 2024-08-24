package main

import (
	"fmt"
	"regexp"
)

func main() {
	re1 := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
	fmt.Println(re1.Split("ASHVJV<HABCXYZKNJBJVKABCXYZ", -1))

	re1 = regexp.MustCompile(`\s+`)
	fmt.Println(re1.Split("aaaaaaaaaa     bbbbbbb  cccccc", -1))
}
