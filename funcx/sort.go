package funcx

import (
	"github.com/ldeng7/go-x/collectionx"
	"github.com/ldeng7/go-x/common"
)

type LessOp[T any] func(T, T) bool

func greaterOp[T any](lessOp LessOp[T]) LessOp[T] {
	return func(a, b T) bool { return lessOp(b, a) }
}

// heap sort

func SliceHeapSort[T any](sl []T, heapified bool, lessOp LessOp[T]) {
	h := (&collectionx.BinaryHeap[T]{}).Init(sl, heapified, greaterOp(lessOp))
	l := len(sl)
	for i := l - 1; i > 0; i-- {
		e, _ := h.Pop()
		sl[i] = e
	}
}

func SliceHeapSortIntrincs[T common.IntrincsOrd](sl []T, heapified bool) {
	SliceHeapSort(sl, heapified, IntrincsLess[T])
}

type sortSolver[T any] struct {
	sl     []T
	lessOp LessOp[T]
}

// insert sort

func (ss *sortSolver[T]) insertionSort(i, j int) {
	sl := ss.sl
	for a := i + 1; a < j; a++ {
		for b := a; b > i && ss.lessOp(sl[b], sl[b-1]); b-- {
			sl[b], sl[b-1] = sl[b-1], sl[b]
		}
	}
}

func SliceInsertionSort[T any](sl []T, lessOp LessOp[T]) {
	ss := &sortSolver[T]{sl, lessOp}
	ss.insertionSort(0, len(sl))
}

func SliceInsertionSortIntrincs[T common.IntrincsOrd](sl []T) {
	SliceInsertionSort(sl, IntrincsLess[T])
}

// quick sort

func (ss *sortSolver[T]) medianOfThree(i1, i0, i2 int) {
	sl := ss.sl
	if ss.lessOp(sl[i1], sl[i0]) {
		sl[i1], sl[i0] = sl[i0], sl[i1]
	}
	if ss.lessOp(sl[i2], sl[i1]) {
		sl[i2], sl[i1] = sl[i1], sl[i2]
		if ss.lessOp(sl[i1], sl[i0]) {
			sl[i1], sl[i0] = sl[i0], sl[i1]
		}
	}
}

func (ss *sortSolver[T]) doPivot(i, j int) (int, int) {
	m := int(uint(i+j) >> 1)
	if j-i > 40 {
		s := (j - i) / 8
		ss.medianOfThree(i, i+s, i+2*s)
		ss.medianOfThree(m, m-s, m+s)
		ss.medianOfThree(j-1, j-1-s, j-1-2*s)
	}
	ss.medianOfThree(i, m, j-1)

	a, c := i+1, j-1
	sl := ss.sl
	for ; a < c && ss.lessOp(sl[a], sl[i]); a++ {
	}
	b := a
	for {
		pivot := sl[i]
		for ; b < c && !ss.lessOp(pivot, sl[b]); b++ {
		}
		for ; b < c && ss.lessOp(pivot, sl[c-1]); c-- {
		}
		if b >= c {
			break
		}
		sl[b], sl[c-1] = sl[c-1], sl[b]
		b, c = b+1, c-1
	}
	protect := j-c < 5
	if !protect && j-c < (j-i)/4 {
		dups := 0
		if !ss.lessOp(sl[i], sl[j-1]) {
			sl[c], sl[j-1] = sl[j-1], sl[c]
			c, dups = c+1, dups+1
		}
		if !ss.lessOp(sl[b-1], sl[i]) {
			b, dups = b-1, dups+1
		}
		if !ss.lessOp(sl[m], sl[i]) {
			sl[m], sl[b-1] = sl[b-1], sl[m]
			b, dups = b-1, dups+1
		}
		protect = dups > 1
	}
	if protect {
		for {
			pivot := sl[i]
			for ; a < b && !ss.lessOp(sl[b-1], pivot); b-- {
			}
			for ; a < b && ss.lessOp(sl[a], pivot); a++ {
			}
			if a >= b {
				break
			}
			sl[a], sl[b-1] = sl[b-1], sl[a]
			a, b = a+1, b-1
		}
	}
	sl[i], sl[b-1] = sl[b-1], sl[i]
	return b - 1, c
}

func (ss *sortSolver[T]) quickSort(i, j, maxDepth int) {
	sl := ss.sl
	for j-i > 12 {
		if maxDepth == 0 {
			SliceHeapSort(sl[i:j], false, ss.lessOp)
			return
		}
		maxDepth--
		mlo, mhi := ss.doPivot(i, j)
		if mlo-i < j-mhi {
			ss.quickSort(i, mlo, maxDepth)
			i = mhi
		} else {
			ss.quickSort(mhi, j, maxDepth)
			j = mlo
		}
	}
	if j-i > 1 {
		for k := i + 6; k < j; k++ {
			if a, b := sl[k], sl[k-6]; ss.lessOp(a, b) {
				sl[k], sl[k-6] = b, a
			}
		}
		ss.insertionSort(i, j)
	}
}

func SliceQuickSort[T any](sl []T, lessOp LessOp[T]) {
	l, maxDepth := len(sl), 0
	for i := l; i != 0; i /= 2 {
		maxDepth++
	}
	maxDepth *= 2
	ss := &sortSolver[T]{sl, lessOp}
	ss.quickSort(0, l, maxDepth)
}

func SliceQuickSortIntrincs[T common.IntrincsOrd](sl []T) {
	SliceQuickSort(sl, IntrincsLess[T])
}

// stable sort

func (ss *sortSolver[T]) symMerge(i, m, j int) {
	sl := ss.sl
	if m-i == 1 {
		a, b := m, j
		for a < b {
			h := int(uint(a+b) >> 1)
			if ss.lessOp(sl[h], sl[i]) {
				a = h + 1
			} else {
				b = h
			}
		}
		for k := i; k < a-1; k++ {
			sl[k], sl[k+1] = sl[k+1], sl[k]
		}
		return
	}

	if j-m == 1 {
		a, b := i, m
		for a < b {
			h := int(uint(a+b) >> 1)
			if !ss.lessOp(sl[m], sl[h]) {
				a = h + 1
			} else {
				b = h
			}
		}
		for k := m; k > a; k-- {
			sl[k], sl[k-1] = sl[k-1], sl[k]
		}
		return
	}

	h := int(uint(i+j) >> 1)
	n := h + m
	var start, r int
	if m > h {
		start, r = n-j, h
	} else {
		start, r = i, m
	}
	p := n - 1
	for start < r {
		h := int(uint(start+r) >> 1)
		if !ss.lessOp(sl[p-h], sl[h]) {
			start = h + 1
		} else {
			r = h
		}
	}

	end := n - start
	if start < m && m < end {
		a, b := m-start, end-m
		for a != b {
			if a > b {
				SliceSwapRange(sl, m-a, m, b)
				a -= b
			} else {
				SliceSwapRange(sl, m-a, m+b-a, a)
				b -= a
			}
		}
		SliceSwapRange(sl, m-a, m, a)
	}
	if i < start && start < h {
		ss.symMerge(i, start, h)
	}
	if h < end && end < j {
		ss.symMerge(h, end, j)
	}
}

func SliceStableSort[T any](sl []T, lessOp LessOp[T]) {
	l := len(sl)
	ss := &sortSolver[T]{sl, lessOp}
	blockSize := 20
	a, b := 0, blockSize
	for b <= l {
		ss.insertionSort(a, b)
		a, b = b, b+blockSize
	}
	ss.insertionSort(a, l)

	for blockSize < l {
		a, b = 0, 2*blockSize
		for b <= l {
			ss.symMerge(a, a+blockSize, b)
			a, b = b, b+2*blockSize
		}
		if m := a + blockSize; m < l {
			ss.symMerge(a, m, l)
		}
		blockSize *= 2
	}
}

func SliceStableSortIntrincs[T common.IntrincsOrd](sl []T) {
	SliceStableSort(sl, IntrincsLess[T])
}

// partial sort

func SlicePartialSort[T any](sl []T, n int, lessOp LessOp[T]) {
	h := (&collectionx.BinaryHeap[T]{}).Init(sl[:n], false, greaterOp(lessOp))
	l := len(sl)
	for i := n; i < l; i++ {
		if e, t := sl[i], sl[0]; lessOp(e, t) {
			h.Set(0, e)
			sl[i] = t
		}
	}
	SliceHeapSort(sl[:n], true, lessOp)
}

func SlicePartialSortIntrincs[T common.IntrincsOrd](sl []T, n int) {
	SlicePartialSort(sl, n, IntrincsLess[T])
}

// nth element

func (ss *sortSolver[T]) unguardedPartition(i, j int) (int, int) {
	m := int(uint(i+j) >> 1)
	if d := j - 1 - i; d > 40 {
		s := d >> 3
		j1 := j - 1 - s
		ss.medianOfThree(i, i+s, i+s*2)
		ss.medianOfThree(m-s, m, m+s)
		ss.medianOfThree(j1-s, j1, j-1)
		ss.medianOfThree(i+s, m, j1)
	} else {
		ss.medianOfThree(i, m, j-1)
	}
	a, b := m, m+1
	sl := ss.sl
	for ; a > i && !ss.lessOp(sl[a-1], sl[a]) && !ss.lessOp(sl[a], sl[a-1]); a-- {
	}
	for ea := sl[a]; b < j && !ss.lessOp(sl[b], ea) && !ss.lessOp(ea, sl[b]); b++ {
	}
	a1, b1 := a, b

	for {
		for ; b1 < j; b1++ {
			ea, eb1 := sl[a], sl[b1]
			if ss.lessOp(ea, eb1) {
				continue
			} else if ss.lessOp(eb1, ea) {
				break
			} else {
				sl[b], sl[b1] = eb1, sl[b]
				b++
			}
		}
		for ; i < a1; a1-- {
			ea, ea1p := sl[a], sl[a1-1]
			if ss.lessOp(ea1p, ea) {
				continue
			} else if ss.lessOp(ea, ea1p) {
				break
			} else {
				a--
				sl[a], sl[a1-1] = ea1p, sl[a]
			}
		}
		if a1 == i && b1 == j {
			return a, b
		}

		if a1 == i {
			if b != b1 {
				sl[a], sl[b] = sl[b], sl[a]
			}
			sl[a], sl[b1] = sl[b1], sl[a]
			a, b, b1 = a+1, b+1, b1+1
		} else if b1 == j {
			a, a1, b = a-1, a1-1, b-1
			if a1 != a {
				sl[a], sl[a1] = sl[a1], sl[a]
			}
			sl[a], sl[b] = sl[b], sl[a]
		} else {
			a1--
			sl[a1], sl[b1] = sl[b1], sl[a1]
			b1++
		}
	}
}

func SliceNthElement[T any](sl []T, n int, lessOp LessOp[T]) {
	ss := &sortSolver[T]{sl, lessOp}
	i, j := 0, len(sl)
	for j-i > 32 {
		i1, j1 := ss.unguardedPartition(i, j)
		if j1 <= n {
			i = j1
		} else if i1 <= n {
			return
		} else {
			j = i1
		}
	}
	ss.insertionSort(i, j)
}

func SliceNthElementIntrincs[T common.IntrincsOrd](sl []T, n int) {
	SliceNthElement(sl, n, IntrincsLess[T])
}

func SliceIsSorted[T any](sl []T, lessOp LessOp[T]) bool {
	l := len(sl)
	for i := l - 1; i > 0; i-- {
		if lessOp(sl[i], sl[i-1]) {
			return false
		}
	}
	return true
}

func SliceIsSortedIntrincs[T common.IntrincsOrd](sl []T, lessOp LessOp[T]) bool {
	return SliceIsSorted(sl, IntrincsLess[T])
}
