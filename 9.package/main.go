package main

import (
	"fmt"
	f "fmt"
	"golang_udemy/9.package/foo"
)

func appName() string {
	const AppName = "Go App"
	var Version = "1.0"
	return AppName + "" + Version
}

/*
func appName(s string) (b string) {
	var s string
	const b = "string"
	return b
}

func appName(s string) (b string) {
	{
		var s string
		const b = "string"
		return b
	}
}
*/

func main() {
	//f.Println(foo.min)

	f.Println(foo.ReturnMin())

	fmt.Println(appName())
}
