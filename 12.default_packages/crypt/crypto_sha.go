package main

import (
	"fmt"
	"io"

	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

func main() {
	s1 := sha1.New()
	io.WriteString(s1, "ABCDE")
	fmt.Printf("%x\n", s1.Sum(nil))

	s256 := sha256.New()
	io.WriteString(s256, "ABCDE")
	fmt.Printf("%x\n", s256.Sum(nil))

	s512 := sha512.New()
	io.WriteString(s512, "ABCDE")
	fmt.Printf("%x\n", s512.Sum(nil))
}
