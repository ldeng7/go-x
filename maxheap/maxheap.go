package maxheap

import (
	"container/heap"
)

type MaxHeap struct {
	Pairs   []interface{}
	SortLen int
	LessCb  func(ia, ib interface{}) bool
	inner   []interface{}
}

// heap.Interface

func (h *MaxHeap) Len() int {
	return len(h.inner)
}

func (h *MaxHeap) Less(i, j int) bool {
	return h.LessCb(h.inner[i], h.inner[j])
}

func (h *MaxHeap) Swap(i, j int) {
	h.inner[i], h.inner[j] = h.inner[j], h.inner[i]
}

func (h *MaxHeap) Push(x interface{}) {
	h.inner = append(h.inner, x)
}

func (h *MaxHeap) Pop() interface{} {
	inner := h.inner
	e := len(inner) - 1
	x := inner[e]
	h.inner = inner[:e]
	return x
}

// heap.Interface end

func (h *MaxHeap) Sort() []interface{} {
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
