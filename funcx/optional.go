package funcx

type Optional[T any] struct {
	data *T
}

func NewOptionalSome[T any](data T) *Optional[T] {
	d := data
	return &Optional[T]{&d}
}

func NewOptionalNone[T any]() *Optional[T] {
	return &Optional[T]{nil}
}

func (o *Optional[T]) IsNone() bool {
	return o.data == nil
}

func (o *Optional[T]) Get() (T, bool) {
	if o.data != nil {
		return *(o.data), true
	}
	var t T
	return t, false
}

func (o *Optional[T]) Set(data T) {
	if o.data != nil {
		*(o.data) = data
	}
}

func OptionalFlatMap[T, U any](o *Optional[T], f func(T) *Optional[U]) *Optional[U] {
	if o.data != nil {
		return f(*(o.data))
	}
	return NewOptionalNone[U]()
}

func OptionalMap[T, U any](o *Optional[T], f func(T) U) *Optional[U] {
	return OptionalFlatMap(o, func(t T) *Optional[U] {
		return NewOptionalSome(f(t))
	})
}
