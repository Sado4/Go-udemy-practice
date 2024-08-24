package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 4, 5, 6}
	test2 := test[4:]
	fmt.Println(test2)
	fmt.Printf("len=%d, cap=%d\n", len(test2), cap(test2))

	test3 := []int{1, 2, 3, 4, 5, 6}
	test4 := test3[:2]
	fmt.Println(test4)
	fmt.Printf("len=%d, cap=%d\n", len(test4), cap(test4))

	test5 := []int{1, 2, 3, 4, 5, 6}
	test6 := test5[3:4]
	fmt.Println(test6)
	fmt.Printf("len=%d, cap=%d\n", len(test6), cap(test6))

	test8 := test7[:2:2]
	fmt.Println(test8)
	fmt.Printf("len=%d, cap=%d\n", len(test8), cap(test8))

}
