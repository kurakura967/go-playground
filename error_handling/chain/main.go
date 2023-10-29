package main

import (
	"errors"
	"fmt"
	"strconv"
)

type MyError struct {
	Err error
}

func (e *MyError) Error() string {
	return fmt.Sprintf("MyError: %v", e.Err)
}

func (e *MyError) Unwrap() error {
	return e.Err
}

func main() {
	err1 := getMyError()
	fmt.Printf("%v\n", err1) // MyError: strconv.Atoi: parsing "a": invalid syntax

	err2 := getStrconvError()
	fmt.Printf("%v\n", err2) // strconv.Atoi: parsing "a": invalid syntax

	var numError *strconv.NumError
	// err1はnumErrorをwrapしているので、errors.Asでtrueが返る
	fmt.Println(errors.As(err1, &numError))
	fmt.Println(errors.As(err2, &numError))

	fmt.Println(errors.Is(err1, numError))
}

func getMyError() error {
	_, err := strconv.Atoi("a")
	return &MyError{Err: err}
}

func getStrconvError() error {
	_, err := strconv.Atoi("a")
	return err
}
