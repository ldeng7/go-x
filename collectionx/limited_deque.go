package collectionx

type ArrayLimitedDeque[T any] struct {
	arr        []T
	capa, cnt  int
	head, tail int
}

func (ald *ArrayLimitedDeque[T]) Init(capa int) *ArrayLimitedDeque[T] {
	ald.arr = make([]T, capa)
	ald.capa, ald.cnt, ald.head, ald.tail = capa, 0, capa-1, 0
	return ald
}

func (ald *ArrayLimitedDeque[T]) Empty() bool {
	return ald.cnt == 0
}

func (ald *ArrayLimitedDeque[T]) Full() bool {
	return ald.cnt == ald.capa
}

func (ald *ArrayLimitedDeque[T]) Head() *T {
	if ald.cnt != 0 {
		i := ald.head + 1
		if i == ald.capa {
			i = 0
		}
		return &ald.arr[i]
	}
	return nil
}

func (ald *ArrayLimitedDeque[T]) Tail() *T {
	if ald.cnt != 0 {
		i := ald.tail - 1
		if i == -1 {
			i = ald.capa - 1
		}
		return &ald.arr[i]
	}
	return nil
}

func (ald *ArrayLimitedDeque[T]) PushFront(val T) bool {
	if ald.cnt != ald.capa {
		ald.arr[ald.head] = val
		ald.head--
		if ald.head == -1 {
			ald.head = ald.capa - 1
		}
		ald.cnt++
		return true
	}
	return false
}

func (ald *ArrayLimitedDeque[T]) PushBack(val T) bool {
	if ald.cnt != ald.capa {
		ald.arr[ald.tail] = val
		ald.tail++
		if ald.tail == ald.capa {
			ald.tail = 0
		}
		ald.cnt++
		return true
	}
	return false
}

func (ald *ArrayLimitedDeque[T]) PopFront() *T {
	if ald.cnt != 0 {
		ald.head++
		if ald.head == ald.capa {
			ald.head = 0
		}
		ald.cnt--
		e := ald.arr[ald.head]
		return &e
	}
	return nil
}

func (ald *ArrayLimitedDeque[T]) PopBack() *T {
	if ald.cnt != 0 {
		ald.tail--
		if ald.tail == -1 {
			ald.tail = ald.capa - 1
		}
		ald.cnt--
		e := ald.arr[ald.tail]
		return &e
	}
	return nil
}

func (ald *ArrayLimitedDeque[T]) Clear() {
	ald.Init(ald.capa)
}
