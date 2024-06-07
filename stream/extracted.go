package stream

// OptFunc 定义操作切片的函数
type OptFunc[T any] func([]T)

func Map[T any, R any](s Stream[T], mapper func(T) R) Stream[R] {
	mapped := make([]R, 0)

	for _, el := range s.source {
		mapped = append(mapped, mapper(el))
	}
	return FromSlice(mapped)
}

func FlatMap[T any, R any](s Stream[T], mapper func(T) Stream[R]) Stream[R] {
	streams := make([]Stream[R], 0)
	s.ForEach(func(t T) {
		streams = append(streams, mapper(t))
	})

	newEl := make([]R, 0)
	for _, str := range streams {
		newEl = append(newEl, str.ToSlice()...)
	}

	return FromSlice(newEl)
}

// Concat 拼接流
func Concat[T any](s Stream[T], others ...Stream[T]) Stream[T] {
	return s.Concat(others...)
}

// GroupingBy 分组
// keyMapper 分组的key
// keyMapper 分组的value
// opts 分组后,各组values的操作,例如排序,去重,等等
func GroupingBy[T any, K string | int | int32 | int64, R any](s Stream[T], keyMapper func(T) K, valueMapper func(T) R, opts ...OptFunc[R]) map[K][]R {
	groups := make(map[K][]R)
	s.ForEach(func(t T) {
		key := keyMapper(t)
		groups[key] = append(groups[key], valueMapper(t))
	})
	for _, vs := range groups {
		for _, opt := range opts {
			opt(vs)
		}
	}
	return groups
}

func ToMap[T any, K string | int | int32 | int64, R any](s Stream[T], keyMapper func(T) K, valueMapper func(T) R, opts ...func(oldV, newV R) R) map[K]R {
	res := make(map[K]R)
	for _, item := range s.source {
		var (
			key   = keyMapper(item)
			value = valueMapper(item)
			oldV  R
			ok    bool
		)

		if oldV, ok = res[key]; !ok {
			res[key] = value
			continue
		}

		var newV R
		for _, opt := range opts {
			if opt != nil {
				newV = opt(oldV, value)
			}
			res[key] = newV
		}
	}
	return res
}
