package slicestream

import "iter"

type (
	basicStream[T any] struct {
		implementor Stream[T] // 实现者，用来使具体流的方法遮蔽自身方法
	}
)

func (basic *basicStream[T]) Find(fn func(T) bool) (val T, isExist bool) {
	next, stop := iter.Pull(basic.implementor.iterator())
	defer stop()

	for v, ok := next(); ok; v, ok = next() {
		if fn(v) {
			return v, true
		}
	}

	return
}

func (basic *basicStream[T]) ForEach(fn func(T)) {
	next, stop := iter.Pull(basic.implementor.iterator())
	defer stop()

	for v, ok := next(); ok; v, ok = next() {
		fn(v)
	}
}

func (basic *basicStream[T]) ToSlice() (s []T) {
	next, stop := iter.Pull(basic.implementor.iterator())
	defer stop()

	for v, ok := next(); ok; v, ok = next() {
		s = append(s, v)
	}

	return
}

func (basic *basicStream[T]) Filter(predicate func(T) bool) Stream[T] {
	return (&filterStream[T]{predicate: predicate, input: basic.implementor}).attach()
}

func (basic *basicStream[T]) Map(mapper func(T) T) Stream[T] {
	return (&mapperStream[T, T]{mapper: mapper, input: basic.implementor}).attach()
}
