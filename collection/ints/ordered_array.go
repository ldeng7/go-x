package ints

import "sort"

type oaValType = int
type oaValCmpCb = func(oaValType, oaValType) bool

type OrderedArray struct {
	arr    []oaValType
	lessCb oaValCmpCb
	eqCb   oaValCmpCb
}

func (oa *OrderedArray) Init(arr []oaValType, sorted bool, lessCb, eqCb oaValCmpCb) *OrderedArray {
	oa.arr = arr
	if (!sorted) && len(arr) > 1 {
		sort.Slice(arr, func(i, j int) bool { return lessCb(arr[i], arr[j]) })
	}
	oa.lessCb = lessCb
	oa.eqCb = eqCb
	return oa
}

func (oa *OrderedArray) Len() int {
	return len(oa.arr)
}

func (oa *OrderedArray) Get(index int) oaValType {
	return oa.arr[index]
}

func (oa *OrderedArray) LowerBound(val oaValType) int {
	i, j := 0, len(oa.arr)
	for i < j {
		h := i + (j-i)>>1
		if oa.lessCb(oa.arr[h], val) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func (oa *OrderedArray) UpperBound(val oaValType) int {
	i, j := 0, len(oa.arr)
	for i < j {
		h := i + (j-i)>>1
		if !oa.lessCb(val, oa.arr[h]) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func (oa *OrderedArray) EqualRange(val oaValType) (int, int) {
	i := oa.LowerBound(val)
	if i == len(oa.arr) || !oa.eqCb(oa.arr[i], val) {
		return -1, -1
	}
	return i, oa.UpperBound(val)
}

func (oa *OrderedArray) Exist(val oaValType) (bool, int) {
	i := oa.LowerBound(val)
	return i != len(oa.arr) && !oa.eqCb(oa.arr[i], val), i
}

func (oa *OrderedArray) Add(val oaValType) {
	i := oa.LowerBound(val)
	if i != len(oa.arr) {
		oa.arr = append(oa.arr, 0)
		copy(oa.arr[i+1:], oa.arr[i:])
		oa.arr[i] = val
	} else {
		oa.arr = append(oa.arr, val)
	}
}

func (oa *OrderedArray) Upsert(val oaValType) {
	i := oa.LowerBound(val)
	if i != len(oa.arr) {
		if !oa.eqCb(oa.arr[i], val) {
			oa.arr = append(oa.arr, 0)
			copy(oa.arr[i+1:], oa.arr[i:])
		}
		oa.arr[i] = val
	} else {
		oa.arr = append(oa.arr, val)
	}
}

func (oa *OrderedArray) RemoveAt(index int) {
	if index != 0 {
		if index != len(oa.arr)-1 {
			copy(oa.arr[index:], oa.arr[index+1:])
		}
		oa.arr = oa.arr[:len(oa.arr)-1]
	} else {
		oa.arr = oa.arr[1:]
	}
}

func (oa *OrderedArray) RemoveOne(val oaValType) {
	i := oa.LowerBound(val)
	if i == len(oa.arr) || !oa.eqCb(oa.arr[i], val) {
		return
	}
	oa.RemoveAt(i)
}

func (oa *OrderedArray) RemoveRange(indexBegin, indexEnd int) {
	if indexBegin != 0 {
		if indexEnd != len(oa.arr) {
			copy(oa.arr[indexBegin:], oa.arr[indexEnd:])
		}
		oa.arr = oa.arr[:len(oa.arr)-(indexEnd-indexBegin)]
	} else {
		oa.arr = oa.arr[indexEnd:]
	}
}

func (oa *OrderedArray) Remove(val oaValType) int {
	i := oa.LowerBound(val)
	if i == len(oa.arr) || !oa.eqCb(oa.arr[i], val) {
		return 0
	}
	ie := oa.UpperBound(val)
	oa.RemoveRange(i, ie)
	return ie - i
}
