package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"runtime/trace"
	"time"
	"os"
	"sync"
)

func RandomWait(ctx context.Context, i int) {
	defer trace.StartRegion(ctx, "randmWait").End()

	fmt.Printf("No.%d start\n", i+1)
	// ランダム時間Sleepする
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	fmt.Printf("No.%d done\n", i+1)
}

func _main()  {
	ctx, task := trace.NewTask(context.Background(), "main")
	defer task.End()

	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func (i int)  {
			defer wg.Done()
			RandomWait(ctx, i)
		}(i)
	}
	wg.Wait()
}

func main()  {
	// 結果出力用にファイルを作成する
	f, err := os.Create("./ch10/02/trace_after.out")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalln("Error:", err)
	}

	defer trace.Stop()

	_main()
}