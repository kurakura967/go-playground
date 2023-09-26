package main

import "fmt"

type Foo struct {
	Bar string
}

type Store struct {
	m map[string]*Foo
}

func updateMap(m map[string]*Foo, key string) {
	m[key].Bar = "baz"
}

type Customer struct {
	ID      string
	Balance float64
}

type store struct {
	m map[string]*Customer
}

func (s *store) storeCustomers(customers []Customer) {
	// スライスを反復処理すると固定のアドレスを持つcが作成される
	for _, c := range customers {
		current := c
		// customerのアドレスを表示する
		fmt.Printf("%p\n", &current)
		s.m[c.ID] = &current
	}
}

func main() {
	s := Store{
		m: map[string]*Foo{
			"foo": &Foo{Bar: "bar"},
		},
	}

	fmt.Println("s.m[\"foo\"].Bar:", s.m["foo"].Bar)
	updateMap(s.m, "foo")
	fmt.Println("s.m[\"foo\"].Bar:", s.m["foo"].Bar)

	store := &store{
		m: make(map[string]*Customer),
	}

	store.storeCustomers([]Customer{
		{ID: "1", Balance: 100},
		{ID: "2", Balance: 200},
		{ID: "3", Balance: 300},
	})

	for _, c := range store.m {
		fmt.Printf("ID: %s, Balance: %.2f\n", c.ID, c.Balance)
	}

	c := []Customer{
		{ID: "1", Balance: 100},
		{ID: "2", Balance: 200},
		{ID: "3", Balance: 300},
	}
	// スライスを反復処理すると一時的な変数が作成される
	// 変数のアドレスは同じだが、値は異なる
	for _, c := range c {
		fmt.Printf("%p\n", &c)
	}
}
