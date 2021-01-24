package ints

import "sort"

type oaElemType = int
type oaElemCmpCb = func(oaElemType, oaElemType) bool

type OrderedArray struct {
	arr    []oaElemType
	lessCb oaElemCmpCb
	eqCb   oaElemCmpCb
}

func (oa *OrderedArray) Init(arr []oaElemType, sorted bool, lessCb, eqCb oaElemCmpCb) *OrderedArray {
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

func (oa *OrderedArray) Get(index int) oaElemType {
	return oa.arr[index]
}

func (oa *OrderedArray) LowerBound(item oaElemType) int {
	i, j := 0, len(oa.arr)
	for i < j {
		h := i + (j-i)>>1
		if oa.lessCb(oa.arr[h], item) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func (oa *OrderedArray) UpperBound(item oaElemType) int {
	i, j := 0, len(oa.arr)
	for i < j {
		h := i + (j-i)>>1
		if !oa.lessCb(item, oa.arr[h]) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func (oa *OrderedArray) EqualRange(item oaElemType) (int, int) {
	i := oa.LowerBound(item)
	if i == len(oa.arr) || !oa.eqCb(oa.arr[i], item) {
		return -1, -1
	}
	return i, oa.UpperBound(item)
}

func (oa *OrderedArray) Exist(item oaElemType) (bool, int) {
	i := oa.LowerBound(item)
	return i != len(oa.arr) && !oa.eqCb(oa.arr[i], item), i
}

func (oa *OrderedArray) Add(item oaElemType) {
	i := oa.LowerBound(item)
	if i != len(oa.arr) {
		oa.arr = append(oa.arr, 0)
		copy(oa.arr[i+1:], oa.arr[i:])
		oa.arr[i] = item
	} else {
		oa.arr = append(oa.arr, item)
	}
}

func (oa *OrderedArray) Upsert(item oaElemType) {
	i := oa.LowerBound(item)
	if i != len(oa.arr) {
		if !oa.eqCb(oa.arr[i], item) {
			oa.arr = append(oa.arr, 0)
			copy(oa.arr[i+1:], oa.arr[i:])
		}
		oa.arr[i] = item
	} else {
		oa.arr = append(oa.arr, item)
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

func (oa *OrderedArray) RemoveOne(item oaElemType) {
	i := oa.LowerBound(item)
	if i == len(oa.arr) || !oa.eqCb(oa.arr[i], item) {
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

func (oa *OrderedArray) Remove(item oaElemType) int {
	i := oa.LowerBound(item)
	if i == len(oa.arr) || !oa.eqCb(oa.arr[i], item) {
		return 0
	}
	ie := oa.UpperBound(item)
	oa.RemoveRange(i, ie)
	return ie - i
}
