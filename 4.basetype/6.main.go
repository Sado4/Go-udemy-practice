package main

import "fmt"

func main() {
	var arr1 [2]int
	fmt.Println(arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)

	var arr5 [5]int = [5]int{1, 2, 3}
	fmt.Println(arr5)

	fmt.Println(arr4[0])
	fmt.Println(arr4[2-1])

	arr4[1] = "F"
	fmt.Println(arr4)

	//fmt.Println(arr1 == arr5)
	//fmt.Println(arr2 == arr3)

}
