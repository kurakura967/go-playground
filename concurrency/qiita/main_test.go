package main

import (
	"testing"
)

func benchMarkAddFunc() {
	dst := &[]int{}
	for i := 0; i < 5; i++ {
		AddFunc(dst)
	}
}

func benchMarkAddFuncWithMutex() {
	dst := &[]int{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go AddFuncWithMutex(dst)
	}
	wg.Wait()
}

func BenchmarkSequential(b *testing.B) {

	for i := 0; i < b.N; i++ {
		benchMarkAddFunc()
	}
}

func BenchmarkConcurrency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchMarkAddFuncWithMutex()
	}
}
