package main

import (
	"errors"
	"fmt"
	"net/http"
)

type transientError struct {
	err error
}

func (t transientError) Error() string {
	// transientError独自のエラーを作成する
	// %vはエラーを別のエラーに変換する
	return fmt.Sprintf("transient error: %v", t.err)
}

func getTransactionAmount(transactionID string) (float32, error) {
	//if len(transactionID) != 5 {
	//	return 0, fmt.Errorf("id is invalid: %s", transactionID)
	//}

	amount, err := getTransactionAmountFromDB(transactionID)
	if err != nil {
		return 0, fmt.Errorf("failed to get amount: %w", err)
	}
	return amount, nil
}

func getTransactionAmountFromDB(transactionID string) (float32, error) {
	return 0, transientError{fmt.Errorf("erorr from getTransactionAmountFromDB")}
}

func main() {
	_, err := getTransactionAmount("111")
	if err != nil {
		switch err := err.(type) {
		// transientErrorをwrapすると、エラーの型が変わってしまう
		case transientError:
			fmt.Println(http.StatusNotFound)
			fmt.Printf("%v\n", err)
		default:
			fmt.Println(http.StatusInternalServerError)
			fmt.Printf("%v\n", err)
		}
	}
	if err != nil {
		// errors.Asを使うと、エラーの型を保持したまま、エラーをwrapできる
		if errors.As(err, &transientError{}) {
			fmt.Println(http.StatusNotFound)
			fmt.Printf("%v\n", err)
		} else {
			fmt.Println(http.StatusInternalServerError)
			fmt.Printf("%v\n", err)
		}
	}

}
