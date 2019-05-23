package ints

import "sort"

type SortedArray []int

func (sa SortedArray) LowerBound(bound int) int {
	return sort.Search(len(sa), func(j int) bool {
		return sa[j] >= bound
	})
}

func (sa SortedArray) Index(item int) int {
	i := sa.LowerBound(item)
	if i == len(sa) || sa[i] != item {
		return -1
	}
	return i
}

func (sa SortedArray) Add(item int) SortedArray {
	i := sa.LowerBound(item)
	if i != len(sa) {
		sa = append(sa, 0)
		copy(sa[i+1:], sa[i:])
		sa[i] = item
	} else {
		sa = append(sa, item)
	}
	return sa
}

func (sa SortedArray) RemoveAt(index int) SortedArray {
	if index != len(sa)-1 {
		copy(sa[index:], sa[index+1:])
	}
	return sa[:len(sa)-1]
}
