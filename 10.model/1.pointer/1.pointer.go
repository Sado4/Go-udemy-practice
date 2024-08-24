package main

import "fmt"

func double(i int) {
	i = i * 2
}

func double2(i *int) {
	*i = *i * 2
}

func main() {
	var n int = 100

	fmt.Println(n)

	fmt.Println(&n)
	//>>0xc000136008

	double(n)
	fmt.Println(n)

	var p *int = &n

	fmt.Println(&p)

	fmt.Println(p)
	//>>0xc000136008

	fmt.Println(*p)
	//>>100

	n = 200
	fmt.Println(*p)
	//200
	*p = 300
	fmt.Println(n)
	//300

	double2(&n)
	fmt.Println(n)

	double2(p)
	fmt.Println(*p)

	fmt.Println(&n)
	fmt.Println(*&n)
	fmt.Println(&*&n)

}
