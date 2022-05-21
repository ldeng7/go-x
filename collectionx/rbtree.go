package collectionx

type RBTreeNode[K any, V any] struct {
	Value    V
	key      K
	tree     *RBTree[K, V]
	isBlack  bool
	parent   *RBTreeNode[K, V]
	children [2]*RBTreeNode[K, V]
}

func (n *RBTreeNode[K, V]) Key() K {
	return n.key
}

func (n *RBTreeNode[K, V]) Prev() *RBTreeNode[K, V] {
	node := n
	if nil == node.tree || node == node.tree.special.children[0] {
		return nil
	} else if nil != node.children[0] {
		child := node.children[0]
		for nil != child.children[1] {
			child = child.children[1]
		}
		node = child
	} else {
		parent := node.parent
		for node == parent.children[0] {
			node, parent = parent, parent.parent
		}
		node = parent
	}
	return node
}

func (n *RBTreeNode[K, V]) Next() *RBTreeNode[K, V] {
	node := n
	if nil == node.tree || node == node.tree.special.children[1] {
		return nil
	} else if nil != node.children[1] {
		node = node.children[1]
		for nil != node.children[0] {
			node = node.children[0]
		}
	} else {
		parent := node.parent
		for node == parent.children[1] {
			node, parent = parent, parent.parent
		}
		if node.children[1] != parent {
			node = parent
		}
	}
	return node
}

type RBTree[K any, V any] struct {
	special RBTreeNode[K, V]
	cnt     int
	lessOp  func(K, K) bool
	eqOp    func(K, K) bool
}

func (t *RBTree[K, V]) Init(lessOp func(K, K) bool) *RBTree[K, V] {
	t.lessOp = lessOp
	t.eqOp = func(a, b K) bool {
		return (!lessOp(a, b)) && (!lessOp(b, a))
	}
	t.special = RBTreeNode[K, V]{}
	sp := &t.special
	t.special.tree, t.special.children[0], t.special.children[1] = t, sp, sp
	return t
}

func (t *RBTree[K, V]) Len() int {
	return t.cnt
}

func (t *RBTree[K, V]) Head() *RBTreeNode[K, V] {
	if t.cnt != 0 {
		return t.special.children[0]
	}
	return nil
}

func (t *RBTree[K, V]) Tail() *RBTreeNode[K, V] {
	if t.cnt != 0 {
		return t.special.children[1]
	}
	return nil
}

func (t *RBTree[K, V]) ForEach(f func(*RBTreeNode[K, V])) {
	for node := t.Head(); node != nil; node = node.Next() {
		f(node)
	}
}

func (t *RBTree[K, V]) LowerBound(key K) *RBTreeNode[K, V] {
	node, parent := t.special.parent, &t.special
	for nil != node {
		if t.lessOp(node.key, key) {
			node = node.children[1]
		} else {
			parent, node = node, node.children[0]
		}
	}
	if parent != &t.special {
		return parent
	}
	return nil
}

func (t *RBTree[K, V]) UpperBound(key K) *RBTreeNode[K, V] {
	node, parent := t.special.parent, &t.special
	for nil != node {
		if !t.lessOp(key, node.key) {
			node = node.children[1]
		} else {
			parent, node = node, node.children[0]
		}
	}
	if parent != &t.special {
		return parent
	}
	return nil
}

func (t *RBTree[K, V]) Exist(key K) (bool, *RBTreeNode[K, V]) {
	node := t.LowerBound(key)
	if node == nil || t.lessOp(key, node.key) {
		return false, nil
	}
	return true, node
}

func (t *RBTree[K, V]) EqualRange(key K) (*RBTreeNode[K, V], *RBTreeNode[K, V]) {
	ok, node := t.Exist(key)
	if !ok {
		return nil, nil
	}
	return node, t.UpperBound(key)
}

func (t *RBTree[K, V]) Count(key K) int {
	ok, node := t.Exist(key)
	if !ok {
		return 0
	}
	node, cnt := node.Next(), 1
	for nodeEnd := t.UpperBound(key); node != nodeEnd; node, cnt = node.Next(), cnt+1 {
	}
	return cnt
}

func (t *RBTree[K, V]) rotate(node *RBTreeNode[K, V], dir uint8) {
	oppDir := dir ^ 1
	child := node.children[oppDir]
	node.children[oppDir] = child.children[dir]
	if nil != child.children[dir] {
		child.children[dir].parent = node
	}
	child.parent = node.parent
	if node == t.special.parent {
		t.special.parent = child
	} else if node == node.parent.children[dir] {
		node.parent.children[dir] = child
	} else {
		node.parent.children[oppDir] = child
	}
	child.children[dir], node.parent = node, child
}

func (t *RBTree[K, V]) add(node, pos, parent *RBTreeNode[K, V]) {
	sp := &t.special
	if parent == sp || nil != pos || t.lessOp(node.key, parent.key) {
		parent.children[0] = node
		if parent == sp {
			sp.parent, sp.children[1] = node, node
		} else if parent == sp.children[0] {
			sp.children[0] = node
		}
	} else {
		parent.children[1] = node
		if parent == sp.children[1] {
			sp.children[1] = node
		}
	}
	t.cnt++
	node.isBlack, node.parent, node.children[0], node.children[1] = false, parent, nil, nil

	for node != sp.parent && !parent.isBlack {
		var dir uint8
		if parent == parent.parent.children[0] {
			dir = 1
		}
		uncle := parent.parent.children[dir]
		if nil != uncle && !uncle.isBlack {
			parent.isBlack, uncle.isBlack = true, true
			parent.parent.isBlack = false
			node = node.parent.parent
		} else {
			if node == parent.children[dir] {
				node = parent
				t.rotate(node, dir^1)
			}
			node.parent.isBlack = true
			node.parent.parent.isBlack = false
			t.rotate(node.parent.parent, dir)
		}
	}
	sp.parent.isBlack = true
}

func (t *RBTree[K, V]) AddUnique(key K, value V) (*RBTreeNode[K, V], bool) {
	pos, parent := t.special.parent, &t.special
	less := true
	for nil != pos {
		parent = pos
		less = t.lessOp(key, pos.key)
		if less {
			pos = pos.children[0]
		} else {
			pos = pos.children[1]
		}
	}
	node := &RBTreeNode[K, V]{Value: value, key: key, tree: t}
	p1 := parent
	if less {
		if p1 == t.special.children[0] {
			t.add(node, pos, parent)
			return node, true
		} else {
			p1 = p1.Prev()
		}
	}
	if t.lessOp(p1.key, key) {
		t.add(node, pos, parent)
		return node, true
	}
	node.tree = nil
	return p1, false
}

func (t *RBTree[K, V]) Add(key K, value V) *RBTreeNode[K, V] {
	pos, parent := t.special.parent, &t.special
	for nil != pos {
		parent = pos
		if t.lessOp(key, pos.key) {
			pos = pos.children[0]
		} else {
			pos = pos.children[1]
		}
	}
	node := &RBTreeNode[K, V]{Value: value, key: key, tree: t}
	t.add(node, pos, parent)
	return node
}

func (t *RBTree[K, V]) remove(node *RBTreeNode[K, V]) *RBTreeNode[K, V] {
	n0 := node
	sp := &t.special
	var c, cp *RBTreeNode[K, V]
	if nil == node.children[0] {
		c = node.children[1]
	} else if nil == node.children[1] {
		c = node.children[0]
	} else {
		node = node.children[1]
		for nil != node.children[0] {
			node = node.children[0]
		}
		c = node.children[1]
	}

	if node != n0 {
		n0.children[0].parent = node
		node.children[0] = n0.children[0]
		if node != n0.children[1] {
			cp = node.parent
			if nil != c {
				c.parent = node.parent
			}
			node.parent.children[0] = c
			node.children[1] = n0.children[1]
			n0.children[1].parent = node
		} else {
			cp = node
		}
		if sp.parent == n0 {
			sp.parent = node
		} else if n0.parent.children[0] == n0 {
			n0.parent.children[0] = node
		} else {
			n0.parent.children[1] = node
		}
		node.parent = n0.parent
		node.isBlack, n0.isBlack = n0.isBlack, node.isBlack
		node = n0
	} else {
		cp = node.parent
		if nil != c {
			c.parent = node.parent
		}
		if sp.parent == n0 {
			sp.parent = c
		} else {
			if n0.parent.children[0] == n0 {
				n0.parent.children[0] = c
			} else {
				n0.parent.children[1] = c
			}
		}
		for _, dir := range [2]uint8{0, 1} {
			if sp.children[dir] == n0 {
				if nil == n0.children[dir^1] {
					sp.children[dir] = n0.parent
				} else {
					c1 := c
					for nil != c1.children[dir] {
						c1 = c1.children[dir]
					}
					sp.children[dir] = c1
				}
			}
		}
	}

	if node.isBlack {
		for c != sp.parent && (nil == c || c.isBlack) {
			var dir uint8
			if c != cp.children[0] {
				dir = 1
			}
			oppDir := dir ^ 1
			c1 := cp.children[oppDir]
			if !c1.isBlack {
				c1.isBlack, cp.isBlack = true, false
				t.rotate(cp, dir)
				c1 = cp.children[oppDir]
			}
			if (nil == c1.children[0] || c1.children[0].isBlack) &&
				(nil == c1.children[1] || c1.children[1].isBlack) {
				c1.isBlack = false
				c, cp = cp, cp.parent
			} else {
				if nil == c1.children[oppDir] || c1.children[oppDir].isBlack {
					if nil != c1.children[oppDir] {
						c1.children[oppDir].isBlack = true
					}
					c1.isBlack = false
					t.rotate(c1, oppDir)
					c1 = cp.children[oppDir]
				}
				c1.isBlack, cp.isBlack = cp.isBlack, true
				if nil != c1.children[oppDir] {
					c1.children[oppDir].isBlack = true
				}
				t.rotate(cp, dir)
				break
			}
		}
		if nil != c {
			c.isBlack = true
		}
	}
	return node
}

func (t *RBTree[K, V]) RemoveAt(node *RBTreeNode[K, V]) {
	if node.tree == t {
		n := t.remove(node)
		n.tree, n.parent, n.children[0], n.children[1] = nil, nil, nil, nil
		t.cnt--
	}
}

func (t *RBTree[K, V]) RemoveRange(nodeBegin, nodeEnd *RBTreeNode[K, V]) int {
	cnt := t.cnt
	for nodeBegin != nodeEnd {
		next := nodeBegin.Next()
		t.RemoveAt(nodeBegin)
		nodeBegin = next
	}
	return cnt - t.cnt
}

func (t *RBTree[K, V]) RemoveOne(key K) {
	ok, node := t.Exist(key)
	if !ok {
		return
	}
	t.RemoveAt(node)
}

func (t *RBTree[K, V]) Remove(key K) int {
	ok, node := t.Exist(key)
	if !ok {
		return 0
	}
	return t.RemoveRange(node, t.UpperBound(key))
}

func (t *RBTree[K, V]) Clear() {
	t.Init(t.lessOp)
}
