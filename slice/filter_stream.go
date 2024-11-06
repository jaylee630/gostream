package slicestream

import (
	"iter"
)

type filterStream[T any] struct {
	basicStream[T]
	predicate func(T) bool
	input     Stream[T] // 传入的流
}

func (filter *filterStream[T]) attach() Stream[T] {
	// 将当前实现者设置为自己
	filter.implementor = filter
	return filter
}

func (filter *filterStream[T]) iterator() iter.Seq[T] {
	return func(yield func(T) bool) {
		for n := range filter.input.iterator() {
			if !filter.predicate(n) {
				continue
			}
			// 过滤通过，则调用yield函数传出当前n
			if !yield(n) {
				return
			}
		}
	}
}
