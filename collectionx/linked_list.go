package collectionx

type LinkedListNode[T any] struct {
	Val  T
	list *LinkedList[T]
	prev *LinkedListNode[T]
	next *LinkedListNode[T]
}

func (lln *LinkedListNode[T]) Prev() *LinkedListNode[T] {
	if n1 := lln.prev; nil != lln.list && n1 != &lln.list.root {
		return n1
	}
	return nil
}

func (lln *LinkedListNode[T]) Next() *LinkedListNode[T] {
	if n1 := lln.next; nil != lln.list && n1 != &lln.list.root {
		return n1
	}
	return nil
}

type LinkedList[T any] struct {
	root LinkedListNode[T]
	cnt  int
}

func (ll *LinkedList[T]) Init() *LinkedList[T] {
	ll.root.next = &ll.root
	ll.root.prev = &ll.root
	ll.cnt = 0
	return ll
}

func (ll *LinkedList[T]) Len() int {
	return ll.cnt
}

func (ll *LinkedList[T]) Head() *LinkedListNode[T] {
	if ll.cnt != 0 {
		return ll.root.next
	}
	return nil
}

func (ll *LinkedList[T]) Tail() *LinkedListNode[T] {
	if ll.cnt != 0 {
		return ll.root.prev
	}
	return nil
}

func (ll *LinkedList[T]) insertAfter(node, mark *LinkedListNode[T]) {
	next := mark.next
	node.prev, node.next, node.list = mark, next, ll
	mark.next, next.prev = node, node
	ll.cnt++
}

func (ll *LinkedList[T]) InsertBefore(node, mark *LinkedListNode[T]) {
	if nil == node.list || mark.list == ll {
		ll.insertAfter(node, mark.prev)
	}
}

func (ll *LinkedList[T]) InsertAfter(node, mark *LinkedListNode[T]) {
	if nil == node.list || mark.list == ll {
		ll.insertAfter(node, mark)
	}
}

func (ll *LinkedList[T]) PushFront(node *LinkedListNode[T]) {
	if nil == node.list {
		ll.insertAfter(node, &ll.root)
	}
}

func (ll *LinkedList[T]) PushBack(node *LinkedListNode[T]) {
	if nil == node.list {
		ll.insertAfter(node, ll.root.prev)
	}
}

func (ll *LinkedList[T]) unlink(node *LinkedListNode[T]) {
	node.prev.next, node.next.prev = node.next, node.prev
	node.next, node.prev, node.list = nil, nil, nil
	ll.cnt--
}

func (ll *LinkedList[T]) Unlink(node *LinkedListNode[T]) {
	if node.list == ll {
		ll.unlink(node)
	}
}

func (ll *LinkedList[T]) PopFront() *LinkedListNode[T] {
	if ll.cnt != 0 {
		node := ll.root.next
		ll.unlink(node)
		return node
	}
	return nil
}

func (ll *LinkedList[T]) PopBack() *LinkedListNode[T] {
	if ll.cnt != 0 {
		node := ll.root.prev
		ll.unlink(node)
		return node
	}
	return nil
}

func (ll *LinkedList[T]) Clear() {
	ll.Init()
}
