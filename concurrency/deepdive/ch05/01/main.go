package main

func restFunc() <-chan int {
	res := make(chan int)

	go func() {
		defer close(res)
		for i := 0; i < 5; i++ {
			res <- 1
		}
	}()

	return res
}

func main() {
	res := restFunc()

	// resの値を受け取る
	// for i := range res {
	// 	println(i)
	// }
	for i := 0; i < 5; i++ {
		println(<-res)
	}
}