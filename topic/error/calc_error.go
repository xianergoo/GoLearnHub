package main

import (
	"errors"
	"fmt"
)

func Div(a int, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("除数不能为零")
	} else {
		return a / b, nil
	}
}

type MyError struct {
	err  error
	Msg  string
	Code int
}

func (e *MyError) Error() string {
	return e.Msg
}

func (e *MyError) Unwarp() error {
	return e.err
}

type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}

var targetErr *ErrorString

func main() {
	// _, err := Div(1, 0)
	// fmt.Printf("err.Error(): %v\n", err.Error())
	// fmt.Printf("err: %v\n", err)

	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%w]", err1)
	err3 := fmt.Errorf("err3: [%w]", err2)

	fmt.Println(errors.Is(err2, err3))
	fmt.Println(errors.Is(err3, err1))
	fmt.Println(err3)
	fmt.Println(errors.Unwrap(err3))
	fmt.Println(errors.Unwrap(errors.Unwrap(err3)))

	fmt.Println("******************************")
	err := fmt.Errorf("new error: [%w]", &ErrorString{s: "target err"})

	fmt.Println(err)
	fmt.Println(errors.As(err, &targetErr))
}
