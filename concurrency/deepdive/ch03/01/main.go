package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

func getLuckyNum() {
	fmt.Println("....")

	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	fmt.Printf("Lucky number is %d\n", num)
}

func main() {
	fmt.Println("what is your lucky number?")
	// go getLuckyNum()
	// 新しいゴルーチンを作成しているので、
	// このスリープをコメントアウトすると、
	// 何も表示されないままプログラムが終了してしまう
	// time.Sleep(5 * time.Second)

	// ゴールーチンの待ち合わせ
	var wg sync.WaitGroup
	// wgの内部カウンターを1増やす
	wg.Add(1)

	// for i := 0; i < 2; i++ {
	// 	go func ()  {
	// 		// wgの内部カウンターを1減らす
	// 		defer wg.Done()
	// 		getLuckyNum()
	// 	}()
	// }
	go func ()  {
		// wgの内部カウンターを1減らす
		defer wg.Done()
		getLuckyNum()
	}()
	// wgの内部カウンターが0になるまで待機する
	wg.Wait()
}
