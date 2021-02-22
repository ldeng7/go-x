package ints

type lqValType = int

type ArrayLimitedQueue struct {
	arr        []lqValType
	capa, cnt  int
	head, tail int
}

func (lq *ArrayLimitedQueue) Init(capa int) *ArrayLimitedQueue {
	lq.arr = make([]lqValType, capa)
	lq.capa, lq.head = capa, capa-1
	return lq
}

func (lq *ArrayLimitedQueue) Empty() bool {
	return lq.cnt == 0
}

func (lq *ArrayLimitedQueue) Full() bool {
	return lq.cnt == lq.capa
}

func (lq *ArrayLimitedQueue) Top() *lqValType {
	if lq.cnt != 0 {
		i := lq.head + 1
		if i == lq.capa {
			i = 0
		}
		return &lq.arr[i]
	}
	return nil
}

func (lq *ArrayLimitedQueue) Push(val lqValType) bool {
	if lq.cnt != lq.capa {
		lq.arr[lq.tail] = val
		lq.tail++
		if lq.tail == lq.capa {
			lq.tail = 0
		}
		lq.cnt++
		return true
	}
	return false
}

func (lq *ArrayLimitedQueue) Pop() *lqValType {
	if lq.cnt != 0 {
		lq.head++
		if lq.head == lq.capa {
			lq.head = 0
		}
		lq.cnt--
		e := lq.arr[lq.head]
		return &e
	}
	return nil
}
