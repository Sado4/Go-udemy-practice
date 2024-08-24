package main

import "fmt"

func main() {

	i := 0
	for {
		i++
		fmt.Println(i)
		if i == 3 {
			break
		}
	}

	point := 1
	for point < 10 {
		fmt.Println(point)
		point++
	}

	for i := 0; i < 100; i++ {
		if i == 3 {
			fmt.Println(3)
			continue
		}
		if i > 30 {
			fmt.Println("抜けた")
			break
		}
		fmt.Println(i)
	}

	arr := [3]int{1, 2, 3}

	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	for _, v := range arr {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for _, v := range m {
		fmt.Println(v)
	}

	for i := 0; i < 100; i += 2 {
		fmt.Println(i)
	}
}
