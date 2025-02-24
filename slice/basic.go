package slicestream

import (
	"iter"
)

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

func (basic *basicStream[T]) Pager(limit, offset *int64) (s []T, totalCount int64) {
	next, stop := iter.Pull(basic.implementor.iterator())
	defer stop()

	var data = make([]T, 0)

	for v, ok := next(); ok; v, ok = next() {
		data = append(data, v)
	}

	totalCount = int64(len(data))

	if limit == nil || len(data) == 0 {
		return
	}
	// 缺省offset
	if offset == nil {
		offset = new(int64)
	}

	var (
		start, end = *offset, *limit + *offset
	)
	// 头索引越界，直接返回空
	if start < 0 || end < 0 || start > totalCount {
		return
	}
	// 尾索引越界，直接返回总量
	if end > totalCount {
		return data[start:], totalCount
	}
	return data[start:end], totalCount
}

func (basic *basicStream[T]) Filter(predicate func(T) bool) Stream[T] {
	return (&filterStream[T]{predicate: predicate, input: basic.implementor}).attach()
}

func (basic *basicStream[T]) Map(mapper func(T) T) Stream[T] {
	return (&mapperStream[T, T]{mapper: mapper, input: basic.implementor}).attach()
}

func (basic *basicStream[T]) Fuzzy(fuzzyFunc func() (fuzzyFields []string, fuzzyValue string)) Stream[T] {
	return (&fuzzyStream[T]{fuzzyFunc: fuzzyFunc, input: basic.implementor}).attach()
}

func (basic *basicStream[T]) Sort(sortFunc func(T, T) bool) Stream[T] {
	return (&sortStream[T]{sortFunc: sortFunc, input: basic.implementor}).attach()
}
