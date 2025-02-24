package slicestream

import (
	"iter"
	"sort"
)

type sortStream[T any] struct {
	basicStream[T]
	sortFunc func(T, T) bool
	input    Stream[T] // 传入的流
}

func (st *sortStream[T]) attach() Stream[T] {
	// 将当前实现者设置为自己
	st.implementor = st
	return st
}

func (st *sortStream[T]) iterator() iter.Seq[T] {
	var data []T
	for n := range st.input.iterator() {
		data = append(data, n)
	}
	sort.Slice(data, func(i, j int) bool {
		return st.sortFunc(data[i], data[j])
	})
	return func(yield func(T) bool) {
		for _, n := range data {
			// 过滤通过，则调用yield函数传出当前n
			if !yield(n) {
				return
			}
		}
	}
}
