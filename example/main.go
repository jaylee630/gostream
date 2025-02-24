package main

import (
	"fmt"
	slicestream "github.com/jaylee630/gostream/slice"
)

func main() {
	slicestream.
		Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11).
		Filter(func(n int) bool {
			return n%2 == 0
		}).
		Map(func(n int) int {
			return n * 100
		}).
		ForEach(func(n int) {
			fmt.Printf("%d ", n)
		})
}
