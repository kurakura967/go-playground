package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
)

// 並行処理を用いて複数の関数から1つのスライスにアクセスする
// ベンチマークテストを実行して逐次処理の場合と比較する

// AddFunc はスライスに要素を追加する関数
func AddFunc(ctx context.Context, dst *[]int) {
	defer trace.StartRegion(ctx, "AddFunc").End()
	// 重い処理
	//time.Sleep(1 * time.Second)
	// ランダムな値をスライスに追加する
	*dst = append(*dst, rand.Intn(10))
}

func _main() {
	ctx, task := trace.NewTask(context.Background(), "main")
	defer task.End()

	dst := &[]int{}
	for i := 0; i < 5; i++ {
		AddFunc(ctx, dst)
	}

	fmt.Printf("dst: %v\n", dst)
}

func main() {
	f, err := os.Create("./trace_sequential.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	trace.Start(f)

	_main()
	trace.Stop()
}
