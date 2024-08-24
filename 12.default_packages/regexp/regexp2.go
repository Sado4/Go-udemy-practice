package main

import (
	"fmt"
	"regexp"
)

func main() {

	re1 := regexp.MustCompile(`\s+`)

	fmt.Println(re1.ReplaceAllString("AAA BBB CCC", ","))

	re1 = regexp.MustCompile(`、|。`)
	fmt.Println(re1.ReplaceAllString("私は、Golangを使用する、プログラマー。", ""))

}
