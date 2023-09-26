package main

import "fmt"

// map[string]intからキーを取り出す
func getKeys(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// map[string]int / map[int]stringからキーを取り出す
func getKeys2(m any) ([]any, error) {
	switch t := m.(type) {
	default:
		return nil, fmt.Errorf("unsupported type: %T", t)
	case map[string]int:
		var keys []any
		for k := range t {
			keys = append(keys, k)
		}
		return keys, nil
	case map[int]string:
		var keys []any
		for k := range t {
			keys = append(keys, k)
		}
		return keys, nil
	}
}

// ジェネリクスを使う場合 comparebleはキーの型が比較可能であることを示す
func getKeys3[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type customConstraint interface {
	~int | ~string
}

func getKeys4[K customConstraint, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type SliceFn[T any] struct {
	S       []T
	Compare func(T, T) bool
}

func (s SliceFn[T]) Len() int {
	return len(s.S)
}

func (s SliceFn[T]) Less(i, j int) bool {
	return s.Compare(s.S[i], s.S[j])
}

func (s SliceFn[T]) Swap(i, j int) {
	s.S[i], s.S[j] = s.S[j], s.S[i]
}

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(getKeys(m))

	m2 := map[int]string{1: "a", 2: "b", 3: "c"}
	res, err := getKeys2(m2)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	fmt.Println(getKeys3(m))
	fmt.Println(getKeys3(m2))

	fmt.Println(getKeys4(m))
	fmt.Println(getKeys4(m2))

	s := SliceFn[int]{
		S: []int{1, 2, 3},
		Compare: func(i, j int) bool {
			return i < j
		},
	}

	t := SliceFn[string]{
		S: []string{"c", "b", "a"},
		Compare: func(i, j string) bool {
			return i < j
		},
	}

	fmt.Println(s.Less(1, 2))
	fmt.Println(t.Less(1, 2))

}
