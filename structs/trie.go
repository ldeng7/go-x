package structs

type TrieNode struct {
	children [26]*TrieNode
	isWord   bool
}

type Trie struct {
	root TrieNode
}

func (t *Trie) Add(word string) {
	n := &t.root
	for i := 0; i < len(word); i++ {
		j := word[i] - 'a'
		if nil == n.children[j] {
			n.children[j] = &TrieNode{}
		}
		n = n.children[j]
	}
	n.isWord = true
}

func (t *Trie) find(key string) *TrieNode {
	n := &t.root
	for i := 0; i < len(key); i++ {
		j := key[i] - 'a'
		if nil == n.children[j] {
			return nil
		}
		n = n.children[j]
	}
	return n
}

func (t *Trie) Search(word string) bool {
	if n := t.find(word); nil != n {
		return n.isWord
	}
	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	return nil != t.find(prefix)
}
