package collectionx

type BinaryHeap[T any] struct {
	arr    []T
	lessOp func(T, T) bool
}

func (bh *BinaryHeap[T]) up(i int) {
	for {
		j := (i - 1) / 2
		if j == i || !bh.lessOp(bh.arr[i], bh.arr[j]) {
			break
		}
		bh.arr[j], bh.arr[i] = bh.arr[i], bh.arr[j]
		i = j
	}
}

func (bh *BinaryHeap[T]) down(i0 int) bool {
	i, l := i0, len(bh.arr)
	for {
		j := i*2 + 1
		if j >= l || j < 0 {
			break
		}
		if j+1 < l && bh.lessOp(bh.arr[j+1], bh.arr[j]) {
			j++
		}
		if !bh.lessOp(bh.arr[j], bh.arr[i]) {
			break
		}
		bh.arr[i], bh.arr[j] = bh.arr[j], bh.arr[i]
		i = j
	}
	return i > i0
}

func (bh *BinaryHeap[T]) Init(arr []T, sorted bool, lessOp func(T, T) bool) *BinaryHeap[T] {
	bh.arr = arr
	bh.lessOp = lessOp
	if !sorted {
		l := len(bh.arr)
		for i := l/2 - 1; i >= 0; i-- {
			bh.down(i)
		}
	}
	return bh
}

func (bh *BinaryHeap[T]) Len() int {
	return len(bh.arr)
}

func (bh *BinaryHeap[T]) Get(i int) T {
	return bh.arr[i]
}

func (bh *BinaryHeap[T]) Top() (T, bool) {
	if len(bh.arr) != 0 {
		return bh.arr[0], true
	}
	var t T
	return t, false
}

func (bh *BinaryHeap[T]) Push(val T) {
	bh.arr = append(bh.arr, val)
	bh.up(len(bh.arr) - 1)
}

func (bh *BinaryHeap[T]) Pop() (T, bool) {
	j := len(bh.arr) - 1
	if j >= 0 {
		e, ej := bh.arr[0], bh.arr[j]
		bh.arr = bh.arr[:j]
		if j != 0 {
			bh.SetTop(ej)
		}
		return e, true
	}
	var t T
	return t, false
}

func (bh *BinaryHeap[T]) RemoveAt(index int) T {
	j := len(bh.arr) - 1
	e, ej := bh.arr[index], bh.arr[j]
	bh.arr = bh.arr[:j]
	if j != index {
		bh.Set(index, ej)
	}
	return e
}

func (bh *BinaryHeap[T]) SetTop(val T) {
	bh.arr[0] = val
	bh.down(0)
}

func (bh *BinaryHeap[T]) Set(index int, val T) {
	bh.arr[index] = val
	if !bh.down(index) {
		bh.up(index)
	}
}

func (bh *BinaryHeap[T]) Clear(keepCap bool) {
	capa := cap(bh.arr)
	if !keepCap {
		capa = 0
	}
	bh.arr = make([]T, 0, capa)
}
