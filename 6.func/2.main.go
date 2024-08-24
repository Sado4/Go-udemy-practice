package main

func main() {
	f := func(x, y int) int {
		return x + y
	}
	f(1, 2)

	func(x, y int) int {
		return x + y
	}(1, 2)

}
