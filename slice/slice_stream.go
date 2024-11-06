package slicestream

import (
	"iter"
)

type sliceStream[T any] struct {
	basicStream[T]
	items []T
}

func (sl *sliceStream[T]) attach() Stream[T] {
	sl.implementor = sl
	return sl
}

func (sl *sliceStream[T]) iterator() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, n := range sl.items {
			if !yield(n) {
				return
			}
		}
	}
}

func (sl *sliceStream[T]) ToSlice() []T {
	return sl.items
}
