package main

import "time"

func main() {
	// よくやるバグ

	for i := 0; i < 3; i++ {
		// go func ()  {
		// 	// このiはforループのiではなく、
		// 	// ゴルーチンが実行された時点でのiの値を参照している
		// 	// そのため、3が3回表示される
		// 	println(i)
		// }()

		go func(i int) {
			// 引数としてiを渡すことで、
			// ゴルーチンが実行された時点でのiの値を参照する
			// そのため、0, 1, 2が表示される
			println(i)
		}(i)
	}

	time.Sleep(1 * time.Second)
}
