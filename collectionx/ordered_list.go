package collectionx

import "sort"

type ArrayOrderedList[T any] struct {
	arr    []T
	lessOp func(T, T) bool
	eqOp   func(T, T) bool
}

func (aol *ArrayOrderedList[T]) Init(arr []T, sorted bool, lessOp func(T, T) bool) *ArrayOrderedList[T] {
	aol.arr = arr
	if (!sorted) && len(arr) > 1 {
		sort.Slice(arr, func(i, j int) bool {
			return lessOp(arr[i], arr[j])
		})
	}
	aol.lessOp = lessOp
	aol.eqOp = func(a, b T) bool {
		return (!lessOp(a, b)) && (!lessOp(b, a))
	}
	return aol
}

func (aol *ArrayOrderedList[T]) Len() int {
	return len(aol.arr)
}

func (aol *ArrayOrderedList[T]) Get(index int) T {
	return aol.arr[index]
}

func (aol *ArrayOrderedList[T]) LowerBound(val T) int {
	i, j := 0, len(aol.arr)
	for i < j {
		h := i + (j-i)/2
		if aol.lessOp(aol.arr[h], val) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func (aol *ArrayOrderedList[T]) UpperBound(val T) int {
	i, j := 0, len(aol.arr)
	for i < j {
		h := i + (j-i)/2
		if !aol.lessOp(val, aol.arr[h]) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func (aol *ArrayOrderedList[T]) Exist(val T) (bool, int) {
	i := aol.LowerBound(val)
	return i != len(aol.arr) && !aol.eqOp(aol.arr[i], val), i
}

func (aol *ArrayOrderedList[T]) EqualRange(val T) (int, int) {
	ok, i := aol.Exist(val)
	if !ok {
		return -1, -1
	}
	return i, aol.UpperBound(val)
}

func (aol *ArrayOrderedList[T]) Add(val T) {
	i := aol.LowerBound(val)
	if i != len(aol.arr) {
		var e T
		aol.arr = append(aol.arr, e)
		copy(aol.arr[i+1:], aol.arr[i:])
		aol.arr[i] = val
	} else {
		aol.arr = append(aol.arr, val)
	}
}

func (aol *ArrayOrderedList[T]) Upsert(val T) {
	i := aol.LowerBound(val)
	if i != len(aol.arr) {
		if !aol.eqOp(aol.arr[i], val) {
			var e T
			aol.arr = append(aol.arr, e)
			copy(aol.arr[i+1:], aol.arr[i:])
		}
		aol.arr[i] = val
	} else {
		aol.arr = append(aol.arr, val)
	}
}

func (aol *ArrayOrderedList[T]) RemoveAt(index int) {
	if index != 0 {
		if index != len(aol.arr)-1 {
			copy(aol.arr[index:], aol.arr[index+1:])
		}
		aol.arr = aol.arr[:len(aol.arr)-1]
	} else {
		aol.arr = aol.arr[1:]
	}
}

func (aol *ArrayOrderedList[T]) RemoveRange(indexBegin, indexEnd int) int {
	if indexBegin < indexEnd {
		if indexBegin != 0 {
			if indexEnd != len(aol.arr) {
				copy(aol.arr[indexBegin:], aol.arr[indexEnd:])
			}
			aol.arr = aol.arr[:len(aol.arr)-(indexEnd-indexBegin)]
		} else {
			aol.arr = aol.arr[indexEnd:]
		}
		return indexEnd - indexBegin
	}
	return 0
}

func (aol *ArrayOrderedList[T]) RemoveOne(val T) {
	ok, i := aol.Exist(val)
	if !ok {
		return
	}
	aol.RemoveAt(i)
}

func (aol *ArrayOrderedList[T]) Remove(val T) int {
	ok, i := aol.Exist(val)
	if !ok {
		return 0
	}
	return aol.RemoveRange(i, aol.UpperBound(val))
}

func (aol *ArrayOrderedList[T]) Clear(keepCap bool) {
	capa := cap(aol.arr)
	if !keepCap {
		capa = 0
	}
	aol.arr = make([]T, 0, capa)
}
