package algo

type TrieNode struct {
	children []*TrieNode
	complete bool
}

type Trie struct {
	nChild      int
	charToIndex func(byte) int
	root        TrieNode
}

func (t *Trie) Init(nChild int, charToIndex func(byte) int) *Trie {
	t.nChild = nChild
	t.charToIndex = charToIndex
	t.root.children = make([]*TrieNode, nChild)
	return t
}

func (t *Trie) Add(word string) {
	n := &t.root
	for i := 0; i < len(word); i++ {
		j := t.charToIndex(word[i])
		if nil == n.children[j] {
			n.children[j] = &TrieNode{children: make([]*TrieNode, t.nChild)}
		}
		n = n.children[j]
	}
	n.complete = true
}

func (t *Trie) find(key string) *TrieNode {
	n := &t.root
	for i := 0; i < len(key); i++ {
		j := t.charToIndex(key[i])
		if nil == n.children[j] {
			return nil
		}
		n = n.children[j]
	}
	return n
}

func (t *Trie) Search(word string) bool {
	if n := t.find(word); nil != n {
		return n.complete
	}
	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	return nil != t.find(prefix)
}
