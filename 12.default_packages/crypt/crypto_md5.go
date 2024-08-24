package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "ABCDE")

	//b := []byte{12, 34, 55, 3}

	fmt.Println(h.Sum(nil))
	//fmt.Println(h.Sum(b))

	//fmt.Println(b)

	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)

}
