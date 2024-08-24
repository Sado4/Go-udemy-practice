package main

import "fmt"

/*
type error interface {
    Error() string
}
*/

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 1000}
}

func DoSomething() (int, error) {
	err := RaiseError()
	return 1, err
}

func main() {
	i, err := DoSomething()
	if err != nil {
		e, ok := err.(*MyError)
		if ok {
			fmt.Println(e.ErrCode)
		}
		fmt.Println(err)
	}
	fmt.Println(i)
}
