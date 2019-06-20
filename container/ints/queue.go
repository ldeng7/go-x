package ints

type queueElemType = int

type Queue interface {
	Empty() bool
	Top() *queueElemType
	Push(queueElemType)
	Pop() *queueElemType
}

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

func (q *ListQueue) Top() *queueElemType {
	h := q.l.Head()
	if nil != h {
		return &h.Val
	}
	return nil
}

func (q *ListQueue) Push(item queueElemType) {
	q.l.PushFront(&ListNode{Val: item})
}

func (q *ListQueue) Pop() *queueElemType {
	h := q.l.PopFront()
	if nil != h {
		return &h.Val
	}
	return nil
}

type ArrayQueue struct {
	arr []queueElemType
	i   int
}

func (q *ArrayQueue) Init() *ArrayQueue {
	q.arr = []queueElemType{}
	return q
}

func (q *ArrayQueue) Empty() bool {
	return len(q.arr)-q.i == 0
}

func (q *ArrayQueue) Top() *queueElemType {
	if len(q.arr)-q.i != 0 {
		return &q.arr[q.i]
	}
	return nil
}

func (q *ArrayQueue) Push(item queueElemType) {
	if len(q.arr) <= 32 || q.i <= (len(q.arr)>>1) {
		q.arr = append(q.arr, item)
	} else {
		arr := make([]queueElemType, len(q.arr)-q.i+1)
		copy(arr, q.arr[q.i:])
		arr[len(arr)-1] = item
		q.arr = arr
		q.i = 0
	}
}

func (q *ArrayQueue) Pop() *queueElemType {
	if len(q.arr)-q.i != 0 {
		e := q.arr[q.i]
		q.i++
		return &e
	}
	return nil
}
