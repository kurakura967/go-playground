package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

func DoSomething(ctx context.Context, c chan int) {

	defer trace.StartRegion(ctx, "DoSomething").End()
	// 重い処理
	time.Sleep(1 * time.Second)
	// ランダムな値をスライスに追加する
	num := rand.Intn(10)
	c <- num
}

func _main() {
	ctx, task := trace.NewTask(context.Background(), "main")
	defer task.End()

	var dst []int
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go DoSomething(ctx, c)
	}

	for i := 0; i < 5; i++ {
		num := <-c
		dst = append(dst, num)
	}
	fmt.Printf("dst: %v\n", dst)
}

func main() {
	f, err := os.Create("./trace_concurrency02.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()

	_main()
}
