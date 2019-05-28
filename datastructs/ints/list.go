package ints

type ListNode struct {
	Val  int
	prev *ListNode
	next *ListNode
}

func (n *ListNode) GetPrev() *ListNode {
	return n.Prev
}

func (n *ListNode) GetNext() *ListNode {
	return n.Next
}

type List struct {
	head *ListNode
	tail *ListNode
}

func (l *List) GetHead() *ListNode {
	return l.head
}

func (l *List) GetTail() *ListNode {
	return l.tail
}

func (l *List) PushFront(n *ListNode) {
	n.prev = nil
	if l.head != nil {
		l.head.prev = n
	} else {
		l.tail = n
	}
	n.next = l.head
	l.head = n
}

func (l *List) PushBack(n *ListNode) {
	n.next = nil
	if l.tail != nil {
		l.tail.next = n
	} else {
		l.head = n
	}
	n.prev = l.tail
	l.tail = n
}

func (l *List) InsertBefore(n, before *ListNode) {
	if before != l.head {
		n.prev, n.next = before.prev, before
		n.prev.next, before.prev = n, n
	} else {
		n.prev, n.next = nil, before
		before.prev, l.head = n, n
	}
}

func (l *List) InsertAfter(n, after *ListNode) {
	if after != l.tail {
		n.prev, n.next = after, after.next
		n.next.prev, after.next = n, n
	} else {
		n.prev, n.next = after, nil
		after.next, l.tail = n, n
	}
}

func (l *List) PopFront() *ListNode {
	if l.head == nil {
		return nil
	}
	h := l.head
	if h.next != nil {
		l.head = h.next
	} else {
		l.head, l.tail = nil, nil
	}
	return h
}

func (l *List) PopBack() *ListNode {
	if l.tail == nil {
		return nil
	}
	r := l.tail
	if r.prev != nil {
		l.tail = r.prev
	} else {
		l.head, l.tail = nil, nil
	}
	return r
}

func (l *List) Unlink(n *ListNode) {
	if l.head == n {
		l.PopFront()
		return
	} else if l.tail == n {
		l.PopBack()
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
}
