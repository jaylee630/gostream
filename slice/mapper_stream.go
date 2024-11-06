package slicestream

import "iter"

type mapperStream[IN, OUT any] struct {
	basicStream[OUT]
	mapper func(IN) OUT
	input  Stream[IN]
}

func (mp *mapperStream[IN, OUT]) attach() Stream[OUT] {
	mp.implementor = mp
	return mp
}

// 迭代器类型为 OUT
func (mp *mapperStream[IN, OUT]) iterator() iter.Seq[OUT] {
	return func(yield func(OUT) bool) {
		for n := range mp.input.iterator() { // n 为 IN 类型
			if !yield(mp.mapper(n)) { // 将 mapper 后的 OUT 类型 通过 yield 传出
				return
			}
		}
	}
}
