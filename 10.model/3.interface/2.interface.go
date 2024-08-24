package main

import "fmt"

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

func main() {
	vs := []Stringify{
		&Person{Name: "Taro", Age: 15},
		&Car{Number: "XXX-012", Model: "RX-78"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}
}
