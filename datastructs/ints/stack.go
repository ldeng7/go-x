package ints

type Stack interface {
	Empty() bool
	Top() (int, bool)
	Push(item int)
	Pop() (int, bool)
}

type ListStack struct {
	l *List
}

func NewListStack() Stack {
	return &ListStack{&List{}}
}

func (s *ListStack) Empty() bool {
	return nil == s.l.head
}

func (s *ListStack) Top() (int, bool) {
	if nil != s.l.tail {
		return s.l.tail.Val, true
	}
	return 0, false
}

func (s *ListStack) Push(item int) {
	s.l.PushBack(&ListNode{Val: item})
}

func (s *ListStack) Pop() (int, bool) {
	tail := s.l.PopBack()
	if nil != tail {
		return tail.Val, true
	}
	return 0, false
}

type ArrayStack struct {
	arr []int
}

func NewArrayStack() Stack {
	return &ArrayStack{[]int{}}
}

func (s *ArrayStack) Empty() bool {
	return len(s.arr) == 0
}

func (s *ArrayStack) Top() (int, bool) {
	if len(s.arr) != 0 {
		return s.arr[len(s.arr)-1], true
	}
	return 0, false
}

func (s *ArrayStack) Push(item int) {
	s.arr = append(s.arr, item)
}

func (s *ArrayStack) Pop() (int, bool) {
	if len(s.arr) != 0 {
		item := s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
		return item, true
	}
	return 0, false
}
