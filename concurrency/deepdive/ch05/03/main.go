package main

func generator(done chan struct{}) <- chan int {
	res := make(chan int)

	go func ()  {
		defer close(res)

	LOOP:
		for {
			res <- 1
			select {
				// q: <- doneの意味は?
				// a: doneチャネルに値が送信されるまでブロックする
				// doneがcloseされると、このゴルーチンは終了する
				case <- done:
					break LOOP
				// q: case res <- 1:の意味は?
				// a: このケースは、resに値を送信できる場合に実行される
				case res <- 1:	
			}
		}
	}()
	return res
}

func main() {
	done := make(chan struct{})

	res := generator(done)
	for i := 0; i < 5; i++ {
		println(<-res)
	}
	close(done)
}