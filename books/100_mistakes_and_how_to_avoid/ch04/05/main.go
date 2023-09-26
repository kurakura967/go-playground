package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
		switch i {
		default:
		case 2:
			break // このbreakはswitch文を抜けるだけfor文を抜けるわけではない
		}
	}

loop:
	for j := 0; j < 5; j++ {
		fmt.Printf("j = %d\n", j)
		switch j {
		default:
		case 2:
			break loop // このbreakはfor文を抜ける
		}
	}
}
