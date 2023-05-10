package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"sync"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
)

// AddFunc はスライスに要素を追加する関数
func AddFunc(ctx context.Context, dst *[]int) {

	defer trace.StartRegion(ctx, "AddFuncWithMutex").End()
	// 重い処理
	//time.Sleep(1 * time.Second)
	// ランダムな値をスライスに追加する
	num := rand.Intn(10)
	mu.Lock()
	*dst = append(*dst, num)
	mu.Unlock()
	wg.Done()
}

func _main() {
	ctx, task := trace.NewTask(context.Background(), "main")
	defer task.End()
	dst := &[]int{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go AddFunc(ctx, dst)
	}
	wg.Wait()
	fmt.Printf("dst: %v\n", dst)
}

func main() {
	f, err := os.Create("./trace_concurrency01.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()

	_main()
}
