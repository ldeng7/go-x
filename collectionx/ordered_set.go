package collectionx

type TreeSet[T any] struct {
	tree *RBTree[T, struct{}]
}

func (ts *TreeSet[T]) Init(lessOp func(T, T) bool) *TreeSet[T] {
	return &TreeSet[T]{
		(&RBTree[T, struct{}]{}).Init(lessOp),
	}
}

func (ts *TreeSet[T]) Len() int {
	return ts.tree.Len()
}

func (ts *TreeSet[T]) ForEach(f func(T)) {
	ts.tree.ForEach(func(node *RBTreeNode[T, struct{}]) {
		f(node.Key())
	})
}

func (ts *TreeSet[T]) Exist(elem T) bool {
	ok, _ := ts.tree.Exist(elem)
	return ok
}

func (ts *TreeSet[T]) Add(elem T) {
	ts.tree.AddUnique(elem, struct{}{})
}

func (ts *TreeSet[T]) Remove(elem T) {
	ts.tree.RemoveOne(elem)
}

func (ts *TreeSet[T]) Clear() {
	ts.tree.Clear()
}
