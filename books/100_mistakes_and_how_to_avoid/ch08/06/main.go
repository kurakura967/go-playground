package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// このコンテキストは4秒後にキャンセルされる
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := Child(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	select {
	case <-ctx.Done():
		fmt.Println("done")
	}
}

func Child(ctx context.Context) (bool, error) {

	//go Sleep() // この関数は3秒後に終了する
	var err error
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			err = ctx.Err()
			return false, err
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
	}
}
