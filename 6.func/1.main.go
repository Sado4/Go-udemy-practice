package main

import "fmt"

/*
func 関数名(引数　引数の型) 戻り値型 {
	関数の中身
	return 返す値
}
*/

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	r1 := Plus(10, 20)
	fmt.Println(r1)

	r2, _ := Div(30, 20)
	fmt.Println(r2)

	r3 := Double(100)
	fmt.Println(r3)

	noreturn()

	f := func(x, y int) int {
		return x + y
	}
	f(1, 2)

	func(x, y int) int {
		return x + y
	}(1, 2)

}
