package ints

type ldElemType = int

type ArrayLimitedDeque struct {
	arr        []ldElemType
	capa, cnt  int
	head, tail int
}

func (ld *ArrayLimitedDeque) Init(capa int) *ArrayLimitedDeque {
	ld.arr = make([]ldElemType, capa)
	ld.capa, ld.head = capa, capa-1
	return ld
}

func (ld *ArrayLimitedDeque) Empty() bool {
	return ld.cnt == 0
}

func (ld *ArrayLimitedDeque) Full() bool {
	return ld.cnt == ld.capa
}

func (ld *ArrayLimitedDeque) Head() *ldElemType {
	if ld.cnt != 0 {
		i := ld.head + 1
		if i == ld.capa {
			i = 0
		}
		return &ld.arr[i]
	}
	return nil
}

func (ld *ArrayLimitedDeque) Tail() *ldElemType {
	if ld.cnt != 0 {
		i := ld.tail - 1
		if i == -1 {
			i = ld.capa - 1
		}
		return &ld.arr[i]
	}
	return nil
}

func (ld *ArrayLimitedDeque) PushFront(item ldElemType) bool {
	if ld.cnt != ld.capa {
		ld.arr[ld.head] = item
		ld.head--
		if ld.head == -1 {
			ld.head = ld.capa - 1
		}
		ld.cnt++
		return true
	}
	return false
}

func (ld *ArrayLimitedDeque) PushBack(item ldElemType) bool {
	if ld.cnt != ld.capa {
		ld.arr[ld.tail] = item
		ld.tail++
		if ld.tail == ld.capa {
			ld.tail = 0
		}
		ld.cnt++
		return true
	}
	return false
}

func (ld *ArrayLimitedDeque) PopFront() *ldElemType {
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

func (ld *ArrayLimitedDeque) PopBack() *ldElemType {
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
