package main

import "fmt"

func main() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	for k, v := range m {
		if v {
			// 新しく要素を追加すると、次の反復でその要素が取り出されるかもしないし、取り出されないかもしれない
			m[10+k] = true
		}
	}
	// 実行毎に結果が異なる
	//fmt.Println(m)

	// mapのコピーを作成する
	m2 := make(map[int]bool)
	for k, v := range m {
		m2[k] = v
	}

	// 読み込むmapと書き込むmapを分ける
	for k, v := range m {
		m2[k] = v
		if v {
			m2[10+k] = true
		}
	}
	fmt.Println(m2)
}
