package funcx

import "github.com/ldeng7/go-x/common"

func NewSliceFilling[T any](len int, cap int, elem T) []T {
	sl := make([]T, len, cap)
	for i := 0; i < len; i++ {
		sl[i] = elem
	}
	return sl
}

func NewSliceBy[T any](len int, cap int, f func(int) T) []T {
	sl := make([]T, len, cap)
	for i := 0; i < len; i++ {
		sl[i] = f(i)
	}
	return sl
}

func SliceClone[T any](sl []T, keepCap bool) []T {
	capNew := len(sl)
	if keepCap {
		capNew = cap(sl)
	}
	slNew := make([]T, len(sl), capNew)
	copy(slNew, sl)
	return slNew
}

func SliceFilter[T any](sl []T, f func(int, T) bool) []T {
	l := len(sl)
	slNew := make([]T, 0, l)
	for i := 0; i < l; i++ {
		if e := sl[i]; f(i, e) {
			slNew = append(slNew, e)
		}
	}
	return slNew
}

func SliceFlatMap[T, U any](sl []T, f func(int, T) []U) []U {
	l := len(sl)
	slNew := make([]U, 0)
	for i := 0; i < l; i++ {
		slNew = append(slNew, f(i, sl[i])...)
	}
	return slNew
}

func SliceForeach[T any](sl []T, f func(int, T)) {
	l := len(sl)
	for i := 0; i < l; i++ {
		f(i, sl[i])
	}
}

func SliceMap[T, U any](sl []T, f func(int, T) U) []U {
	l := len(sl)
	slNew := make([]U, l)
	for i := 0; i < l; i++ {
		slNew[i] = f(i, sl[i])
	}
	return slNew
}

func SliceMaxIntrincs[T common.IntrincsOrd](sl []T) T {
	return SliceMax(sl, IntrincsLess[T])
}

func SliceMax[T any](sl []T, lessOp func(T, T) bool) T {
	l := len(sl)
	r := sl[0]
	for i := 1; i < l; i++ {
		if e := sl[i]; lessOp(r, e) {
			r = e
		}
	}
	return r
}

func SliceMinIntrincs[T common.IntrincsOrd](sl []T) T {
	return SliceMin(sl, IntrincsLess[T])
}

func SliceMin[T any](sl []T, lessOp func(T, T) bool) T {
	l := len(sl)
	r := sl[0]
	for i := 1; i < l; i++ {
		if e := sl[i]; lessOp(e, r) {
			r = e
		}
	}
	return r
}

func SliceReduce[T, U any](sl []T, f func(U, int, T) U, init U) U {
	l := len(sl)
	acc := init
	for i := 0; i < l; i++ {
		acc = f(acc, i, sl[i])
	}
	return acc
}

func SliceReverse[T any](sl []T) {
	l := len(sl)
	for i := 0; i < l/2; i++ {
		j := l - i - 1
		sl[i], sl[j] = sl[j], sl[i]
	}
}

func SliceScan[T, U any](sl []T, f func(U, int, T) U, init U) []U {
	l := len(sl)
	slNew := make([]U, len(sl))
	acc := init
	for i := 0; i < l; i++ {
		acc = f(acc, i, sl[i])
		slNew[i] = acc
	}
	return slNew
}

func SliceSwapRange[T any](sl []T, i, j, n int) {
	for k := 0; k < n; k++ {
		sl[i+k], sl[j+k] = sl[j+k], sl[i+k]
	}
}

func SliceZip2[T, U any](slA []T, slB []U, f func(int, T, U) Tuple2[T, U]) []Tuple2[T, U] {
	return SliceZip2By(slA, slB, f)
}

func SliceZip2By[T, U, V any](slA []T, slB []U, f func(int, T, U) V) []V {
	l := IntrincsMin2(len(slA), len(slB))
	slNew := make([]V, l)
	for i := 0; i < l; i++ {
		slNew[i] = f(i, slA[i], slB[i])
	}
	return slNew
}

func SliceZip3[T, U, V any](slA []T, slB []U, slC []V, f func(int, T, U, V) Tuple3[T, U, V]) []Tuple3[T, U, V] {
	return SliceZip3By(slA, slB, slC, f)
}

func SliceZip3By[T, U, V, W any](slA []T, slB []U, slC []V, f func(int, T, U, V) W) []W {
	l := IntrincsMin(len(slA), len(slB), len(slC))
	slNew := make([]W, l)
	for i := 0; i < l; i++ {
		slNew[i] = f(i, slA[i], slB[i], slC[i])
	}
	return slNew
}
