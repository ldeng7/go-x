package ints

type ArrayLimitedQueue struct {
	arr        []int
	l, c, h, t int
}

func (lq *ArrayLimitedQueue) Init(capa int) *ArrayLimitedQueue {
	lq.arr = make([]int, capa)
	lq.l, lq.h = capa, capa-1
	return lq
}

func (lq *ArrayLimitedQueue) Empty() bool {
	return lq.c == 0
}

func (lq *ArrayLimitedQueue) Full() bool {
	return lq.c == lq.l
}

func (lq *ArrayLimitedQueue) Top() (int, bool) {
	if lq.c != 0 {
		i := lq.h + 1
		if i == lq.l {
			i = 0
		}
		return lq.arr[i], true
	}
	return 0, false
}

func (lq *ArrayLimitedQueue) Push(item int) bool {
	if lq.c != lq.l {
		lq.arr[lq.t] = item
		lq.t++
		if lq.t == lq.l {
			lq.t = 0
		}
		lq.c++
		return true
	}
	return false
}

func (lq *ArrayLimitedQueue) Pop() (int, bool) {
	if lq.c != 0 {
		lq.h++
		if lq.h == lq.l {
			lq.h = 0
		}
		lq.c--
		return lq.arr[lq.h], true
	}
	return 0, false
}
