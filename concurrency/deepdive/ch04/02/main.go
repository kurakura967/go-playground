package main

import (
	"fmt"
	"time"
)

func main() {
	// データが競合した場合
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	// srcの要素に対して何かの処理をしてdestに追加する
	for _, s := range src {
		go func(s int) {
			// 何か重い処理
			res := s*2
			// dstを複数のゴルーチンが同時に更新しようとするのでデータの競合が発生する
			// dstに追加
			dst = append(dst, res)
		}(s)
	}

	time.Sleep(time.Second)
	// 実行毎に結果が変わる
	fmt.Println(dst)
}
