package main

import (
	"fmt"
	slicestream "github.com/jaylee630/gostream/slice"
)

type Foo struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

func main() {
	var limit = int64(1)
	s, totalCount := slicestream.
		Of(Foo{1, "foo"}, Foo{2, "bar"}, Foo{3, "baz"}).
		Fuzzy(func() (fuzzyFields []string, fuzzyValue string) {
			return []string{"Name"}, "ba"
		}).
		//Filter(func(n Foo) bool {
		//	return n.Id%2 == 0
		//}).
		//Map(func(n int) int {
		//	return n * 100
		//}).
		Sort(func(a, b Foo) bool {
			return a.Id > b.Id
		}).
		Pager(&limit, nil)

	fmt.Println("=====>", s, totalCount)
}
