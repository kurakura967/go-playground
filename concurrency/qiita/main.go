package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	// 並行処理によるスライスへのアクセス
	dst := &[]int{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go AddFuncWithMutex(dst)
	}
	wg.Wait()
	fmt.Printf("conccurency-> dst: %v\n", dst)

	// 逐次処理によるスライスへのアクセス
	dst = &[]int{}
	for i := 0; i < 5; i++ {
		AddFunc(dst)
	}
	fmt.Printf("sequential-> dst: %v\n", dst)
}

func AddFuncWithMutex(dst *[]int) {
	// 重い処理
	//time.Sleep(1 * time.Second)
	mutex.Lock()
	defer mutex.Unlock()

	*dst = append(*dst, rand.Intn(10))
	wg.Done()
}

func AddFunc(dst *[]int) {
	// 重い処理
	// time.Sleep(1 * time.Second)
	*dst = append(*dst, rand.Intn(10))
}
