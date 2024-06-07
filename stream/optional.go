package stream

type Optional[T any] struct {
	v *T
}

func (o Optional[T]) Get() (v T, ok bool) {
	if o.v == nil {
		return *new(T), false
	}
	return *o.v, true
}

// IsPresent 判断是否有值
func (o Optional[T]) IsPresent() bool {
	return o.v != nil
}

// IfPresent 如果有值则执行fn函数
func (o Optional[T]) IfPresent(fn func(T)) {
	if o.v != nil {
		fn(*o.v)
	}
}
