package ints

type listValType = int

type ListNode struct {
	Val  listValType
	list *List
	prev *ListNode
	next *ListNode
}

func (n *ListNode) Prev() *ListNode {
	if n1 := n.prev; nil != n.list && n1 != &n.list.root {
		return n1
	}
	return nil
}

func (n *ListNode) Next() *ListNode {
	if n1 := n.next; nil != n.list && n1 != &n.list.root {
		return n1
	}
	return nil
}

type List struct {
	root ListNode
	cnt  int
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	return l
}

func (l *List) Len() int {
	return l.cnt
}

func (l *List) Head() *ListNode {
	if l.cnt != 0 {
		return l.root.next
	}
	return nil
}

func (l *List) Tail() *ListNode {
	if l.cnt != 0 {
		return l.root.prev
	}
	return nil
}

func (l *List) insertAfter(n, mark *ListNode) {
	next := mark.next
	n.prev, n.next, n.list = mark, next, l
	mark.next, next.prev = n, n
	l.cnt++
}

func (l *List) InsertBefore(n, mark *ListNode) {
	if nil == n.list || mark.list == l {
		l.insertAfter(n, mark.prev)
	}
}

func (l *List) InsertAfter(n, mark *ListNode) {
	if nil == n.list || mark.list == l {
		l.insertAfter(n, mark)
	}
}

func (l *List) PushFront(n *ListNode) {
	if nil == n.list {
		l.insertAfter(n, &l.root)
	}
}

func (l *List) PushBack(n *ListNode) {
	if nil == n.list {
		l.insertAfter(n, l.root.prev)
	}
}

func (l *List) unlink(n *ListNode) {
	n.prev.next, n.next.prev = n.next, n.prev
	n.next, n.prev, n.list = nil, nil, nil
	l.cnt--
}

func (l *List) Unlink(n *ListNode) {
	if n.list == l {
		l.unlink(n)
	}
}

func (l *List) PopFront() *ListNode {
	if l.cnt != 0 {
		n := l.root.next
		l.unlink(n)
		return n
	}
	return nil
}

func (l *List) PopBack() *ListNode {
	if l.cnt != 0 {
		n := l.root.prev
		l.unlink(n)
		return n
	}
	return nil
}
