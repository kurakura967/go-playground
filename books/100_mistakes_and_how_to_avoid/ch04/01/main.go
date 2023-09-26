package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	// rangeはループ開始前に1度だけ評価される
	for range s {
		s = append(s, 10)
	}
	fmt.Println(s)

	t := []int{1, 2, 3}
	// len(t)はループごとに評価されるので、無限ループになる
	for i := 0; i < len(t); i++ {
		t = append(t, 10)
	}
}
