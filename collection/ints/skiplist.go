package ints

import (
	"math/rand"
	"time"
)

const (
	sklMaxLevel int = 18
)

var sklProbs = [sklMaxLevel - 1]int64{
	3393088950634442752, 1248247667004394496, 459204654181133312, 168931951563480736,
	62146591937174464, 22862453512557408, 8410626622007697, 3094096621605848,
	1138254536086807, 418740442646473, 154046000036667, 56670356408185,
	20847859046429, 7669498735621, 2821450908925, 1037953783668, 381841857897,
}

type sklKeyType = int
type sklValType = int
type sklKeyCmpCb = func(sklKeyType, sklKeyType) bool

type SkipListNode struct {
	Value sklValType
	key   sklKeyType
	next  []*SkipListNode
}

func (n *SkipListNode) Key() sklKeyType {
	return n.key
}

type SkipList struct {
	root      SkipListNode
	length    int
	randSrc   rand.Source
	prevNodes []*SkipListNode
	eqCb      sklKeyCmpCb
	lessCb    sklKeyCmpCb
}

func (l *SkipList) Init(eqCb, lessCb sklKeyCmpCb) *SkipList {
	l.root.next = make([]*SkipListNode, sklMaxLevel)
	l.randSrc = rand.NewSource(time.Now().UnixNano())
	l.prevNodes = make([]*SkipListNode, sklMaxLevel)
	l.eqCb, l.lessCb = eqCb, lessCb
	return l
}

func (l *SkipList) Len() int {
	return l.length
}

func (l *SkipList) Get(key sklKeyType) *SkipListNode {
	p := &l.root
	var n *SkipListNode
	for i := sklMaxLevel - 1; i >= 0; i-- {
		n = p.next[i]
		for n != nil && l.lessCb(n.key, key) {
			p, n = n, n.next[i]
		}
	}
	if n != nil && l.eqCb(n.key, key) {
		return n
	}
	return nil
}

func (l *SkipList) getPrevNodes(key sklKeyType) []*SkipListNode {
	p := &l.root
	prevs := l.prevNodes
	for i := sklMaxLevel - 1; i >= 0; i-- {
		n := p.next[i]
		for n != nil && l.lessCb(n.key, key) {
			p, n = n, n.next[i]
		}
		prevs[i] = p
	}
	return prevs
}

func (l *SkipList) Add(key sklKeyType, value sklValType) (*SkipListNode, bool) {
	prevs := l.getPrevNodes(key)
	if e := prevs[0].next[0]; e != nil && l.eqCb(e.key, key) {
		return e, false
	}
	r := l.randSrc.Int63()
	level := 1
	for ; level < sklMaxLevel && r < sklProbs[level-1]; level++ {
	}
	node := &SkipListNode{value, key, make([]*SkipListNode, level)}
	for i := 0; i < level; i++ {
		node.next[i] = prevs[i].next[i]
		prevs[i].next[i] = node
	}
	l.length++
	return node, true
}

func (l *SkipList) Remove(key sklKeyType) *SkipListNode {
	prevs := l.getPrevNodes(key)
	if e := prevs[0].next[0]; e != nil && l.eqCb(e.key, key) {
		for i, n := range e.next {
			prevs[i].next[i] = n
		}
		l.length--
		return e
	}
	return nil
}
