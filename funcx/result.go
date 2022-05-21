package funcx

type Result[T any] struct {
	data *T
	err  error
}

func NewResultOk[T any](data T) *Result[T] {
	d := data
	return &Result[T]{&d, nil}
}

func NewResultErr[T any](err error) *Result[T] {
	return &Result[T]{nil, err}
}

func NewResult[T any](data T, err error) *Result[T] {
	if err == nil {
		return NewResultOk(data)
	}
	return NewResultErr[T](err)
}

func (r *Result[T]) IsErr() bool {
	return r.data == nil
}

func (r *Result[T]) GetData() (T, bool) {
	if r.data != nil {
		return *(r.data), true
	}
	var t T
	return t, false
}

func (r *Result[T]) GetError() error {
	return r.err
}

func ResultFlatMap[T, U any](r *Result[T], f func(T) *Result[U]) *Result[U] {
	if r.data != nil {
		return f(*(r.data))
	}
	return NewResultErr[U](r.err)
}

func ResultFlatMapCascade2[T1, T2, T3 any](r *Result[T1], f1 func(T1) *Result[T2], f2 func(T2) *Result[T3]) *Result[T3] {
	r1 := ResultFlatMap(r, f1)
	return ResultFlatMap(r1, f2)
}

func ResultFlatMapCascade3[T1, T2, T3, T4 any](r *Result[T1], f1 func(T1) *Result[T2], f2 func(T2) *Result[T3],
	f3 func(T3) *Result[T4]) *Result[T4] {
	r1 := ResultFlatMap(r, f1)
	r2 := ResultFlatMap(r1, f2)
	return ResultFlatMap(r2, f3)
}

func ResultFlatMapCascade4[T1, T2, T3, T4, T5 any](r *Result[T1], f1 func(T1) *Result[T2], f2 func(T2) *Result[T3],
	f3 func(T3) *Result[T4], f4 func(T4) *Result[T5]) *Result[T5] {
	r1 := ResultFlatMap(r, f1)
	r2 := ResultFlatMap(r1, f2)
	r3 := ResultFlatMap(r2, f3)
	return ResultFlatMap(r3, f4)
}

func ResultFlatMapCascade5[T1, T2, T3, T4, T5, T6 any](r *Result[T1], f1 func(T1) *Result[T2], f2 func(T2) *Result[T3],
	f3 func(T3) *Result[T4], f4 func(T4) *Result[T5], f5 func(T5) *Result[T6]) *Result[T6] {
	r1 := ResultFlatMap(r, f1)
	r2 := ResultFlatMap(r1, f2)
	r3 := ResultFlatMap(r2, f3)
	r4 := ResultFlatMap(r3, f4)
	return ResultFlatMap(r4, f5)
}

func ResultFlatMapCascade6[T1, T2, T3, T4, T5, T6, T7 any](r *Result[T1], f1 func(T1) *Result[T2], f2 func(T2) *Result[T3],
	f3 func(T3) *Result[T4], f4 func(T4) *Result[T5], f5 func(T5) *Result[T6], f6 func(T6) *Result[T7]) *Result[T7] {
	r1 := ResultFlatMap(r, f1)
	r2 := ResultFlatMap(r1, f2)
	r3 := ResultFlatMap(r2, f3)
	r4 := ResultFlatMap(r3, f4)
	r5 := ResultFlatMap(r4, f5)
	return ResultFlatMap(r5, f6)
}

type ResultRawFunc[T, U any] func(T) (U, error)

func ResultFlatMapRaw[T, U any](r *Result[T], f ResultRawFunc[T, U]) *Result[U] {
	return ResultFlatMap(r, func(t T) *Result[U] {
		return NewResult(f(t))
	})
}

func ResultFlatMapRawCascade2[T1, T2, T3 any](r *Result[T1], f1 ResultRawFunc[T1, T2], f2 ResultRawFunc[T2, T3]) *Result[T3] {
	r1 := ResultFlatMapRaw(r, f1)
	return ResultFlatMapRaw(r1, f2)
}

func ResultFlatMapRawCascade3[T1, T2, T3, T4 any](r *Result[T1], f1 ResultRawFunc[T1, T2], f2 ResultRawFunc[T2, T3],
	f3 ResultRawFunc[T3, T4]) *Result[T4] {
	r1 := ResultFlatMapRaw(r, f1)
	r2 := ResultFlatMapRaw(r1, f2)
	return ResultFlatMapRaw(r2, f3)
}

func ResultFlatMapRawCascade4[T1, T2, T3, T4, T5 any](r *Result[T1], f1 ResultRawFunc[T1, T2], f2 ResultRawFunc[T2, T3],
	f3 ResultRawFunc[T3, T4], f4 ResultRawFunc[T4, T5]) *Result[T5] {
	r1 := ResultFlatMapRaw(r, f1)
	r2 := ResultFlatMapRaw(r1, f2)
	r3 := ResultFlatMapRaw(r2, f3)
	return ResultFlatMapRaw(r3, f4)
}

func ResultFlatMapRawCascade5[T1, T2, T3, T4, T5, T6 any](r *Result[T1], f1 ResultRawFunc[T1, T2], f2 ResultRawFunc[T2, T3],
	f3 ResultRawFunc[T3, T4], f4 ResultRawFunc[T4, T5], f5 ResultRawFunc[T5, T6]) *Result[T6] {
	r1 := ResultFlatMapRaw(r, f1)
	r2 := ResultFlatMapRaw(r1, f2)
	r3 := ResultFlatMapRaw(r2, f3)
	r4 := ResultFlatMapRaw(r3, f4)
	return ResultFlatMapRaw(r4, f5)
}

func ResultFlatMapRawCascade6[T1, T2, T3, T4, T5, T6, T7 any](r *Result[T1], f1 ResultRawFunc[T1, T2], f2 ResultRawFunc[T2, T3],
	f3 ResultRawFunc[T3, T4], f4 ResultRawFunc[T4, T5], f5 ResultRawFunc[T5, T6], f6 ResultRawFunc[T6, T7]) *Result[T7] {
	r1 := ResultFlatMapRaw(r, f1)
	r2 := ResultFlatMapRaw(r1, f2)
	r3 := ResultFlatMapRaw(r2, f3)
	r4 := ResultFlatMapRaw(r3, f4)
	r5 := ResultFlatMapRaw(r4, f5)
	return ResultFlatMapRaw(r5, f6)
}

func ResultMap[T, U any](r *Result[T], f func(T) U) *Result[U] {
	return ResultFlatMap(r, func(t T) *Result[U] {
		return NewResultOk(f(t))
	})
}
