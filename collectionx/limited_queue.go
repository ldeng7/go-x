package collectionx

type ArrayLimitedQueue[T any] struct {
	arr        []T
	capa, cnt  int
	head, tail int
}

func (alq *ArrayLimitedQueue[T]) Init(capa int) *ArrayLimitedQueue[T] {
	alq.arr = make([]T, capa)
	alq.capa, alq.cnt, alq.head, alq.tail = capa, 0, capa-1, 0
	return alq
}

func (alq *ArrayLimitedQueue[T]) Empty() bool {
	return alq.cnt == 0
}

func (alq *ArrayLimitedQueue[T]) Full() bool {
	return alq.cnt == alq.capa
}

func (alq *ArrayLimitedQueue[T]) Top() *T {
	if alq.cnt != 0 {
		i := alq.head + 1
		if i == alq.capa {
			i = 0
		}
		return &alq.arr[i]
	}
	return nil
}

func (alq *ArrayLimitedQueue[T]) Push(val T) bool {
	if alq.cnt != alq.capa {
		alq.arr[alq.tail] = val
		alq.tail++
		if alq.tail == alq.capa {
			alq.tail = 0
		}
		alq.cnt++
		return true
	}
	return false
}

func (alq *ArrayLimitedQueue[T]) Pop() *T {
	if alq.cnt != 0 {
		alq.head++
		if alq.head == alq.capa {
			alq.head = 0
		}
		alq.cnt--
		e := alq.arr[alq.head]
		return &e
	}
	return nil
}

func (alq *ArrayLimitedQueue[T]) Clear() {
	alq.Init(alq.capa)
}
