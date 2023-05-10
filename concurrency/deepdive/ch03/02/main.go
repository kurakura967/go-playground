package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getLuckyNum(c chan<- int)  {
	fmt.Println("....")

	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	// メインゴールーチンにラッキーナンバーを伝える
	c <- num
}

func main() {
	fmt.Println("what is your lucky number?")

	// ラッキーナンバーをメインゴールーチンで表示する
	c := make(chan int)
	go getLuckyNum(c)

	// メインゴールーチンはチャンネルcからラッキーナンバーを受け取るまで待機する
	num := <- c

	fmt.Printf("Lucky number is %d\n", num)
	close(c)
}
