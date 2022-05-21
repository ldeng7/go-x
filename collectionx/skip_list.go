package collectionx

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

type SkipListNode[K any, V any] struct {
	Value V
	key   K
	next  []*SkipListNode[K, V]
}

func (sln *SkipListNode[K, V]) Key() K {
	return sln.key
}

type SkipList[K any, V any] struct {
	root      SkipListNode[K, V]
	length    int
	randSrc   rand.Source
	prevNodes []*SkipListNode[K, V]
	lessOp    func(K, K) bool
	eqOp      func(K, K) bool
}

func (sl *SkipList[K, V]) Init(lessOp func(K, K) bool) *SkipList[K, V] {
	sl.root.next = make([]*SkipListNode[K, V], sklMaxLevel)
	sl.length = 0
	sl.randSrc = rand.NewSource(time.Now().UnixNano())
	sl.prevNodes = make([]*SkipListNode[K, V], sklMaxLevel)
	sl.lessOp = lessOp
	sl.eqOp = func(a, b K) bool {
		return (!lessOp(a, b)) && (!lessOp(b, a))
	}
	return sl
}

func (sl *SkipList[K, V]) Len() int {
	return sl.length
}

func (sl *SkipList[K, V]) Get(key K) *SkipListNode[K, V] {
	prev := &sl.root
	var node *SkipListNode[K, V]
	for i := sklMaxLevel - 1; i >= 0; i-- {
		node = prev.next[i]
		for node != nil && sl.lessOp(node.key, key) {
			prev, node = node, node.next[i]
		}
	}
	if node != nil && sl.eqOp(node.key, key) {
		return node
	}
	return nil
}

func (sl *SkipList[K, V]) getPrevNodes(key K) []*SkipListNode[K, V] {
	prev := &sl.root
	prevs := sl.prevNodes
	for i := sklMaxLevel - 1; i >= 0; i-- {
		node := prev.next[i]
		for node != nil && sl.lessOp(node.key, key) {
			prev, node = node, node.next[i]
		}
		prevs[i] = prev
	}
	return prevs
}

func (sl *SkipList[K, V]) Add(key K, value V) (*SkipListNode[K, V], bool) {
	prevs := sl.getPrevNodes(key)
	if e := prevs[0].next[0]; e != nil && sl.eqOp(e.key, key) {
		return e, false
	}
	r := sl.randSrc.Int63()
	level := 1
	for ; level < sklMaxLevel && r < sklProbs[level-1]; level++ {
	}
	node := &SkipListNode[K, V]{value, key, make([]*SkipListNode[K, V], level)}
	for i := 0; i < level; i++ {
		node.next[i] = prevs[i].next[i]
		prevs[i].next[i] = node
	}
	sl.length++
	return node, true
}

func (sl *SkipList[K, V]) Remove(key K) *SkipListNode[K, V] {
	prevs := sl.getPrevNodes(key)
	if e := prevs[0].next[0]; e != nil && sl.eqOp(e.key, key) {
		for i, node := range e.next {
			prevs[i].next[i] = node
		}
		sl.length--
		return e
	}
	return nil
}

func (sl *SkipList[K, V]) Clear() {
	sl.Init(sl.lessOp)
}
