package slicestream

import "iter"

type Stream[T any] interface {
	// 流迭代器自定义实现
	iterator() iter.Seq[T]

	// 用于在 Stream 的基础上绑定或附加新的行为。
	// 可以理解为创建了一个新的 Stream 实例，使得流操作可以连续调用，形成流式链式调用风格
	attach() Stream[T]

	// Filter 过滤流中的元素，返回新的 Stream 实例。
	// predicate 是一个条件函数，用于测试每个元素是否满足条件，
	Filter(predicate func(T) bool) Stream[T]

	// Map 用于将流中的每个元素应用某种转换，返回新的 Stream 实例。
	// 提案：规范：允许方法中的类型参数:https://github.com/golang/go/issues/49085
	Map(mapper func(T) T) Stream[T]

	// Find 用于在流中查找符合条件的元素，返回找到的第一个元素和一个标识是否找到的布尔值
	Find(fn func(T) bool) (val T, isExist bool)

	// ForEach 对每个元素执行 fn 操作，用于遍历流中的所有元素并对其执行某些操作
	ForEach(fn func(T))

	// ToSlice 将流中的所有元素收集到一个切片中
	ToSlice() []T
}

func Of[T any](elems ...T) Stream[T] {
	return OfSlice(elems)
}

func OfSlice[T any](elems []T) Stream[T] {
	return (&sliceStream[T]{items: elems}).attach()
}
