package collectionx

type TreeMap[K any, V any] struct {
	tree *RBTree[K, V]
}

func (tm *TreeMap[K, V]) Init(lessOp func(K, K) bool) *TreeMap[K, V] {
	return &TreeMap[K, V]{
		(&RBTree[K, V]{}).Init(lessOp),
	}
}

func (tm *TreeMap[K, V]) Len() int {
	return tm.tree.Len()
}

func (tm *TreeMap[K, V]) ForEach(f func(K, *V)) {
	tm.tree.ForEach(func(node *RBTreeNode[K, V]) {
		f(node.Key(), &node.Value)
	})
}

func (tm *TreeMap[K, V]) Get(key K) *V {
	_, node := tm.tree.Exist(key)
	if node != nil {
		return &node.Value
	}
	return nil
}

func (tm *TreeMap[K, V]) Add(key K, value V) bool {
	_, ok := tm.tree.AddUnique(key, value)
	return ok
}

func (tm *TreeMap[K, V]) Set(key K, value V) {
	if ok, node := tm.tree.Exist(key); ok {
		node.Value = value
	} else {
		tm.tree.Add(key, value)
	}
}

func (tm *TreeMap[K, V]) Remove(key K) {
	tm.tree.RemoveOne(key)
}

func (tm *TreeMap[K, V]) Clear() {
	tm.tree.Clear()
}
