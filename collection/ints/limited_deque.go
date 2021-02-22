package ints

type ldValType = int

type ArrayLimitedDeque struct {
	arr        []ldValType
	capa, cnt  int
	head, tail int
}

func (ld *ArrayLimitedDeque) Init(capa int) *ArrayLimitedDeque {
	ld.arr = make([]ldValType, capa)
	ld.capa, ld.head = capa, capa-1
	return ld
}

func (ld *ArrayLimitedDeque) Empty() bool {
	return ld.cnt == 0
}

func (ld *ArrayLimitedDeque) Full() bool {
	return ld.cnt == ld.capa
}

func (ld *ArrayLimitedDeque) Head() *ldValType {
	if ld.cnt != 0 {
		i := ld.head + 1
		if i == ld.capa {
			i = 0
		}
		return &ld.arr[i]
	}
	return nil
}

func (ld *ArrayLimitedDeque) Tail() *ldValType {
	if ld.cnt != 0 {
		i := ld.tail - 1
		if i == -1 {
			i = ld.capa - 1
		}
		return &ld.arr[i]
	}
	return nil
}

func (ld *ArrayLimitedDeque) PushFront(val ldValType) bool {
	if ld.cnt != ld.capa {
		ld.arr[ld.head] = val
		ld.head--
		if ld.head == -1 {
			ld.head = ld.capa - 1
		}
		ld.cnt++
		return true
	}
	return false
}

func (ld *ArrayLimitedDeque) PushBack(val ldValType) bool {
	if ld.cnt != ld.capa {
		ld.arr[ld.tail] = val
		ld.tail++
		if ld.tail == ld.capa {
			ld.tail = 0
		}
		ld.cnt++
		return true
	}
	return false
}

func (ld *ArrayLimitedDeque) PopFront() *ldValType {
	if ld.cnt != 0 {
		ld.head++
		if ld.head == ld.capa {
			ld.head = 0
		}
		ld.cnt--
		e := ld.arr[ld.head]
		return &e
	}
	return nil
}

func (ld *ArrayLimitedDeque) PopBack() *ldValType {
	if ld.cnt != 0 {
		ld.tail--
		if ld.tail == -1 {
			ld.tail = ld.capa - 1
		}
		ld.cnt--
		e := ld.arr[ld.tail]
		return &e
	}
	return nil
}
