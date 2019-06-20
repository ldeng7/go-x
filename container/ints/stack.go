package ints

type stackElemType = int

type Stack interface {
	Empty() bool
	Top() *stackElemType
	Push(stackElemType)
	Pop() *stackElemType
}

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

func (s *ListStack) Top() *stackElemType {
	t := s.l.Tail()
	if nil != t {
		return &t.Val
	}
	return nil
}

func (s *ListStack) Push(item stackElemType) {
	s.l.PushBack(&ListNode{Val: item})
}

func (s *ListStack) Pop() *stackElemType {
	t := s.l.PopBack()
	if nil != t {
		return &t.Val
	}
	return nil
}

type ArrayStack struct {
	arr []stackElemType
}

func (s *ArrayStack) Init() *ArrayStack {
	s.arr = []stackElemType{}
	return s
}

func (s *ArrayStack) Empty() bool {
	return len(s.arr) == 0
}

func (s *ArrayStack) Top() *stackElemType {
	if len(s.arr) != 0 {
		return &s.arr[len(s.arr)-1]
	}
	return nil
}

func (s *ArrayStack) Push(item stackElemType) {
	s.arr = append(s.arr, item)
}

func (s *ArrayStack) Pop() *stackElemType {
	if len(s.arr) != 0 {
		e := s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
		return &e
	}
	return nil
}
