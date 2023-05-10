package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	var mu sync.Mutex
	
	for _, s := range src {
		go func(s int) {
			res := s * 2
			
			// dst変数に排他制御をかける
			mu.Lock()
			dst = append(dst, res)
			mu.Unlock()
		}(s)
	}

	time.Sleep(time.Second)
	fmt.Println(dst)
}