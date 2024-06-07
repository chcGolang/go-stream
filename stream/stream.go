package stream

import "sort"

type Stream[T any] struct {
	source []T
}

// FromSlice creates stream from slice.
func FromSlice[T any](source []T) Stream[T] {
	return Stream[T]{source: source}
}

// Of creates a stream whose elements are the specified values.
func Of[T any](elems ...T) Stream[T] {
	return FromSlice(elems)
}

/*****************************Stream中间处理*******************************/

// Reverse 对流中元素进行返转操作
func (s Stream[T]) Reverse() Stream[T] {
	var itemLen = len(s.source)
	source := make([]T, itemLen)

	for i := 0; i < itemLen; i++ {
		source[i] = s.source[itemLen-1-i]
	}

	return FromSlice(source)
}

// Filter 对流中元素进行过滤操作
func (s Stream[T]) Filter(predicate func(item T) bool) Stream[T] {
	source := make([]T, 0)

	for _, v := range s.source {
		if predicate(v) {
			source = append(source, v)
		}
	}

	return FromSlice(source)
}

// Skip 跳过当前流前面指定个数的元素
func (s Stream[T]) Skip(n int) Stream[T] {
	if n <= 0 {
		return s
	}
	source := make([]T, 0)
	l := len(s.source)
	if n > l {
		return FromSlice(source)
	}
	for i := n; i < l; i++ {
		source = append(source, s.source[i])
	}

	return FromSlice(source)
}

// Limit 截取当前流前指定个数的元素
func (s Stream[T]) Limit(maxSize int) Stream[T] {
	if s.source == nil {
		return s
	}
	if maxSize < 0 {
		return FromSlice([]T{})
	}
	source := make([]T, 0, maxSize)
	for i := 0; i < len(s.source) && i < maxSize; i++ {
		source = append(source, s.source[i])
	}
	return FromSlice(source)
}

// Concat 拼接流
func (s Stream[T]) Concat(others ...Stream[T]) Stream[T] {
	if len(others) == 0 {
		return s
	}
	source := make([]T, 0)
	source = append(source, s.source...)
	for _, other := range others {
		source = append(source, other.source...)
	}
	return FromSlice(source)
}

// Distinct 对流中元素进行去重操作
func (s Stream[T]) Distinct() Stream[T] {
	return s.DistinctBy(func(item T) any {
		return item
	})
}

// DistinctBy 对流中元素进行去重操作
func (s Stream[T]) DistinctBy(fn func(item T) any) Stream[T] {
	source := make([]T, 0)

	distinct := map[any]bool{}

	for _, v := range s.source {
		key := fn(v)
		if _, ok := distinct[key]; !ok {
			distinct[key] = true
			source = append(source, v)
		}
	}

	return FromSlice(source)
}

// Sorted 对流中元素进行排序操作
func (s Stream[T]) Sorted(less func(a, b T) bool) Stream[T] {
	source := make([]T, 0, len(s.source))
	source = append(source, s.source...)

	sort.Slice(source, func(i, j int) bool {
		return less(source[i], source[j])
	})
	return FromSlice(source)
}

// Peek 对流中元素指针进行消费操作
func (s Stream[T]) Peek(consumer func(item *T)) Stream[T] {
	for k := range s.source {
		consumer(&s.source[k])
	}

	return s
}

// PeekP 对流中元素进行消费操作
func (s Stream[T]) PeekP(consumer func(item T)) Stream[T] {
	for _, v := range s.source {
		consumer(v)
	}
	return s
}

func (s Stream[T]) Map(fn func(item T) T) Stream[T] {
	return Map[T](s, fn)
}

func (s Stream[T]) MapToString(mapper func(T) string) Stream[string] {
	return Map[T](s, mapper)
}

func (s Stream[T]) MapToInt(mapper func(T) int) Stream[int] {
	return Map[T](s, mapper)
}

func (s Stream[T]) MapToInt32(mapper func(T) int32) Stream[int32] {
	return Map[T](s, mapper)
}

func (s Stream[T]) MapToInt64(mapper func(T) int64) Stream[int64] {
	return Map[T](s, mapper)
}

func (s Stream[T]) MapToFloat64(mapper func(T) float64) Stream[float64] {
	return Map[T](s, mapper)
}

func (s Stream[T]) MapToFloat32(mapper func(T) float32) Stream[float32] {
	return Map[T](s, mapper)
}

func (s Stream[T]) FlatMap(mapper func(T) Stream[T]) Stream[T] {
	return FlatMap[T](s, mapper)
}

func (s Stream[T]) FlatMapToString(mapper func(T) Stream[string]) Stream[string] {
	return FlatMap[T](s, mapper)
}

func (s Stream[T]) FlatMapToInt(mapper func(T) Stream[int]) Stream[int] {
	return FlatMap[T](s, mapper)
}
func (s Stream[T]) FlatMapToInt32(mapper func(T) Stream[int32]) Stream[int32] {
	return FlatMap[T](s, mapper)
}
func (s Stream[T]) FlatMapToInt64(mapper func(T) Stream[int64]) Stream[int64] {
	return FlatMap[T](s, mapper)
}
func (s Stream[T]) FlatMapToFloat64(mapper func(T) Stream[float64]) Stream[float64] {
	return FlatMap[T](s, mapper)
}
func (s Stream[T]) FlatMapToFloat32(mapper func(T) Stream[float32]) Stream[float32] {
	return FlatMap[T](s, mapper)
}

/*****************************Stream的终止*********************************/

// FindLast 获取最后一个元素
func (s Stream[T]) FindLast() Optional[T] {
	if s.source == nil || len(s.source) == 0 {
		return Optional[T]{v: nil}
	}
	return Optional[T]{v: &s.source[len(s.source)-1]}
}

// FindFirst 获取第一个元素
func (s Stream[T]) FindFirst() Optional[T] {
	if s.source == nil || len(s.source) == 0 {
		return Optional[T]{v: nil}
	}
	return Optional[T]{v: &s.source[0]}
}

// ForEach 对元素进行逐个遍历
func (s Stream[T]) ForEach(action func(item T)) {
	for _, v := range s.source {
		action(v)
	}
}

// Reduce 对流中元素进行聚合处理
func (s Stream[T]) Reduce(accumulator func(itemA, itemB T) T) Optional[T] {
	var cnt = 0
	var initial T
	for _, v := range s.source {
		if cnt == 0 {
			cnt++
			initial = v
			continue
		}
		cnt++
		initial = accumulator(initial, v)
	}

	if cnt == 0 {
		return Optional[T]{v: nil}
	}
	return Optional[T]{v: &initial}
}

// AnyMatch 返回此流中是否存在元素满足所提供的条件
func (s Stream[T]) AnyMatch(predicate func(item T) bool) bool {
	for _, v := range s.source {
		if predicate(v) {
			return true
		}
	}

	return false
}

// AllMatch 返回此流中是否全都满足条件
func (s Stream[T]) AllMatch(predicate func(item T) bool) bool {
	for _, v := range s.source {
		if !predicate(v) {
			return false
		}
	}

	return true
}

// NoneMatch 返回此流中是否全都不满足条件
func (s Stream[T]) NoneMatch(predicate func(item T) bool) bool {
	return !s.AnyMatch(predicate)
}

// Count 返回流中元素个数
func (s Stream[T]) Count() (count int) {
	return len(s.source)
}

// Max 返回流中最大元素
func (s Stream[T]) Max(comparator func(newItem T, oldItem T) bool) Optional[T] {
	var (
		max T
		ok  bool
	)
	if max, ok = s.FindFirst().Get(); !ok {
		return Optional[T]{v: nil}
	}
	s.ForEach(func(t T) {
		if comparator(t, max) {
			max = t
		}
	})
	return Optional[T]{v: &max}
}

// Min 返回流中最小元素
func (s Stream[T]) Min(comparator func(newItem T, oldItem T) bool) Optional[T] {
	var (
		min T
		ok  bool
	)
	if min, ok = s.FindFirst().Get(); !ok {
		return Optional[T]{v: nil}
	}

	s.ForEach(func(t T) {
		if comparator(t, min) {
			min = t
		}
	})
	return Optional[T]{v: &min}
}

// ToSlice 将流处理后转化为切片
func (s Stream[T]) ToSlice() []T {
	return s.source
}

// ToMapString 将流处理后转化为string map
func (s Stream[T]) ToMapString(keyMapper func(T) string, valueMapper func(T) T, opts ...func(oldV, newV T) T) map[string]T {
	return ToMap(s, keyMapper, valueMapper, opts...)
}

// ToMapInt 将流处理后转化为int map
func (s Stream[T]) ToMapInt(keyMapper func(T) int, valueMapper func(T) T, opts ...func(oldV, newV T) T) map[int]T {
	return ToMap(s, keyMapper, valueMapper, opts...)
}

// GroupingByString 分组统计,key为字符串
func (s Stream[T]) GroupingByString(groupFunc func(T) string, opts ...OptFunc[T]) map[string][]T {
	return GroupingBy(s, groupFunc, func(t T) T {
		return t
	}, opts...)
}

// GroupingByInt 分组统计,key为int
func (s Stream[T]) GroupingByInt(groupFunc func(T) int, opts ...OptFunc[T]) map[int][]T {
	return GroupingBy(s, groupFunc, func(t T) T {
		return t
	}, opts...)
}
