package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	c := make(chan int)

	for _, s := range src {
		go func(s int, c chan int) {
			res := s*2
			c <- res
		}(s, c)
	}

	for _ = range src {
		// 並行処理の結果を受け取る
		num := <- c
		dst = append(dst, num)
	}
	fmt.Println(dst)
	close(c)
}