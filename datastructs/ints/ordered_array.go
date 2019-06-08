package ints

import "sort"

type OrderedArray struct {
	Arr    []int
	lessCb func(a, b int) bool
}

func (oa *OrderedArray) Len() int           { return len(oa.Arr) }
func (oa *OrderedArray) Less(i, j int) bool { return oa.lessCb(oa.Arr[i], oa.Arr[j]) }
func (oa *OrderedArray) Swap(i, j int)      { oa.Arr[i], oa.Arr[j] = oa.Arr[j], oa.Arr[i] }

func (oa *OrderedArray) Init(arr []int, lessCb func(int, int) bool) *OrderedArray {
	oa.Arr = arr
	if len(arr) > 1 {
		sort.Sort(oa)
	}
	oa.lessCb = lessCb
	if nil == lessCb {
		oa.lessCb = func(a, b int) bool { return a < b }
	}
	return oa
}

func (oa *OrderedArray) binSearch(item int) int {
	return sort.Search(len(oa.Arr), func(index int) bool {
		return !oa.lessCb(oa.Arr[index], item)
	})
}

func (oa *OrderedArray) Index(item int) int {
	i := oa.binSearch(item)
	if i == len(oa.Arr) || oa.Arr[i] != item {
		return -1
	}
	return i
}

func (oa *OrderedArray) Add(item int) {
	i := oa.binSearch(item)
	if i != len(oa.Arr) {
		oa.Arr = append(oa.Arr, 0)
		copy(oa.Arr[i+1:], oa.Arr[i:])
		oa.Arr[i] = item
	} else {
		oa.Arr = append(oa.Arr, item)
	}
}

func (oa *OrderedArray) RemoveAt(index int) {
	if index != len(oa.Arr)-1 {
		copy(oa.Arr[index:], oa.Arr[index+1:])
	}
	oa.Arr = oa.Arr[:len(oa.Arr)-1]
}

func (oa *OrderedArray) Fix() {
	sort.Sort(oa)
}

func (oa *OrderedArray) InitAsReversed(arr []int) *OrderedArray {
	return oa.Init(arr, func(a, b int) bool { return a > b })
}
