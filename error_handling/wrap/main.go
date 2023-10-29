package main

import (
	"errors"
	"fmt"
)

var originalError = errors.New("original error")
var cleanUpError = errors.New("cleanUp error")

func cleanUp() error {
	return cleanUpError
}

func MyFunc() (err error) {
	defer func() {
		if _err := cleanUp(); _err != nil {
			// 関数を抜ける前にエラー(original error)をwrapする
			err = fmt.Errorf("cleanUp failed: %w", err)
		}
	}()

	return originalError
}

func MyFunc2() (err error) {
	defer func() {
		if _err := cleanUp(); _err != nil {
			err = errors.Join(err, _err)
		}
	}()

	return originalError
}

func main() {
	err := MyFunc()
	fmt.Println(err)
	if err != nil {
		fmt.Println(errors.Is(err, originalError))
		fmt.Println(errors.Is(err, cleanUpError))
	}

	err = MyFunc2()
	fmt.Println(err)
	if err != nil {
		fmt.Println(errors.Is(err, originalError))
		fmt.Println(errors.Is(err, cleanUpError))
	}
}
