package main

import "fmt"

type T struct {
	ID   int
	Name string
}

func (t *T) String() string {
	return fmt.Sprintf("ID=%v Name=%v", t.ID, t.Name)
}

func main() {
	t := &T{ID: 1, Name: "Taro"}
	fmt.Println(t)

}
