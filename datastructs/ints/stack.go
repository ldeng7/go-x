package ints

type ListStack struct {
	l List
}

func (s *ListStack) Init() *ListStack {
	s.l = List{}
	s.l.Init()
	return s
}

func (s *ListStack) Empty() bool {
	return 0 == s.l.Len()
}

func (s *ListStack) Top() (int, bool) {
	t := s.l.Tail()
	if nil != t {
		return t.Val, true
	}
	return 0, false
}

func (s *ListStack) Push(item int) {
	s.l.PushBack(&ListNode{Val: item})
}

func (s *ListStack) Pop() (int, bool) {
	t := s.l.PopBack()
	if nil != t {
		return t.Val, true
	}
	return 0, false
}

type ArrayStack struct {
	arr []int
}

func (s *ArrayStack) Init() *ArrayStack {
	s.arr = []int{}
	return s
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
