package main

import (
	"bufio"
	"os"
)

func main() {
	f, err := os.OpenFile("bufio.d", os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriterSize(f, 1)
	w.Write([]byte("Hello World\n"))
	w.Write([]byte("Hello World\n"))
	w.Write([]byte("Hello World\n"))
	// Writeメソッドを実行する度にerrハンドリングを行う必要はない
	// Writer構造体にはerrフィールドがあり、複数回write実行しエラーが発生した場合、
	// 2回目以降のwriteは実行されない
	// Writer構造体のerrフィールドには、最後に発生したエラーが格納される
	// Flushメソッドを実行すると、errフィールドがnilでない場合、そのエラーが返される
	err = w.Flush()
	if err != nil {
		panic(err)
	}
}
