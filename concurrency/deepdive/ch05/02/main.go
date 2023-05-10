package main

func generator() <-chan int {
	res := make(chan int)

	go func() {
		defer close(res)
		for {
			res <- 1
		}
	}()

	return res
}

func main() {
	// 別のゴールーチンが動き続ける
	res := generator()
	for i:= 0; i < 5; i++ {
		println(<-res)
	}
}