package minheap

import (
	"container/heap"
)

type MinHeap struct {
	Pairs   []interface{}
	SortLen int
	LessCb  func(ia, ib interface{}) bool
	inner   []interface{}
}

// heap.Interface

func (h *MinHeap) Len() int {
	return len(h.inner)
}

func (h *MinHeap) Less(i, j int) bool {
	return h.LessCb(h.inner[i], h.inner[j])
}

func (h *MinHeap) Swap(i, j int) {
	h.inner[i], h.inner[j] = h.inner[j], h.inner[i]
}

func (h *MinHeap) Push(x interface{}) {
	h.inner = append(h.inner, x)
}

func (h *MinHeap) Pop() interface{} {
	inner := h.inner
	e := len(inner) - 1
	x := inner[e]
	h.inner = inner[:e]
	return x
}

// heap.Interface end

func (h *MinHeap) Sort() []interface{} {
	l := h.SortLen
	h.inner = h.Pairs[:l]
	heap.Init(h)
	for i := l; i < len(h.Pairs); i++ {
		heap.Push(h, h.Pairs[i])
		heap.Pop(h)
	}
	out := make([]interface{}, l)
	for i := l - 1; i >= 0 && len(h.inner) > 0; i-- {
		out[i] = heap.Pop(h)
	}
	return out
}
