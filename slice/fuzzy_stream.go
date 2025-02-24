package slicestream

import (
	"fmt"
	"iter"
	"reflect"
	"strings"
)

const (
	JsonTag = "json"
)

type fuzzyStream[T any] struct {
	basicStream[T]
	fuzzyFunc func() (fuzzyFields []string, fuzzyValue string)
	input     Stream[T] // 传入的流
}

func (fuzzy *fuzzyStream[T]) attach() Stream[T] {
	// 将当前实现者设置为自己
	fuzzy.implementor = fuzzy
	return fuzzy
}

func (fuzzy *fuzzyStream[T]) iterator() iter.Seq[T] {
	fuzzyFields, fuzzyValue := fuzzy.fuzzyFunc()
	return func(yield func(T) bool) {
		for n := range fuzzy.input.iterator() {
			dataMap, toMapErr := ToMap(n, JsonTag)
			if toMapErr != nil {
				continue // 跳过当前数据
			}
			for _, field := range fuzzyFields {
				value, ok := dataMap[field]
				if !ok {
					continue
				}
				valueStr, isStr := value.(string)
				if !isStr {
					continue
				}
				if strings.Contains(valueStr, fuzzyValue) {
					// 过滤通过，则调用yield函数传出当前n
					if !yield(n) {
						return
					}
				}
			}
		}
	}
}

// ToMap 结构体转为Map[string]interface{}
func ToMap[T any](in T, tagName string) (map[string]any, error) {
	out := make(map[string]any)

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}
