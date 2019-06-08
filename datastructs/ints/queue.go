package ints

type ListQueue struct {
	l List
}

func (q *ListQueue) Init() *ListQueue {
	q.l = List{}
	q.l.Init()
	return q
}

func (q *ListQueue) Empty() bool {
	return 0 == q.l.Len()
}

func (q *ListQueue) Top() (int, bool) {
	h := q.l.Head()
	if nil != h {
		return h.Val, true
	}
	return 0, false
}

func (q *ListQueue) Push(item int) {
	q.l.PushFront(&ListNode{Val: item})
}

func (q *ListQueue) Pop() (int, bool) {
	h := q.l.PopFront()
	if nil != h {
		return h.Val, true
	}
	return 0, false
}

type ArrayQueue struct {
	arr []int
	i   int
}

func (q *ArrayQueue) Init() *ArrayQueue {
	q.arr = []int{}
	return q
}

func (q *ArrayQueue) Empty() bool {
	return len(q.arr)-q.i == 0
}

func (q *ArrayQueue) Top() (int, bool) {
	if len(q.arr)-q.i != 0 {
		return q.arr[q.i], true
	}
	return 0, false
}

func (q *ArrayQueue) Push(item int) {
	if len(q.arr) <= 32 || q.i <= (len(q.arr)>>1) {
		q.arr = append(q.arr, item)
	} else {
		arr := make([]int, len(q.arr)-q.i+1)
		copy(arr, q.arr[q.i:])
		arr[len(arr)-1] = item
		q.arr = arr
		q.i = 0
	}
}

func (q *ArrayQueue) Pop() (int, bool) {
	if len(q.arr)-q.i != 0 {
		item := q.arr[q.i]
		q.i++
		return item, true
	}
	return 0, false
}
