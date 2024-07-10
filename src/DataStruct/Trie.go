package DataStruct

type Trie struct {
	son []*Trie
	end bool
}

func NewTrie() *Trie {
	return &Trie{son: make([]*Trie, 26), end: false}
}

func (tr *Trie) Insert(s string) {
	cur := tr
	for _, c := range s {
		ne := c - 'a'
		if cur.son[ne] == nil {
			cur.son[ne] = NewTrie()
		}
		cur = cur.son[ne]
	}
	cur.end = true
}

func (tr *Trie) Search(s string) bool {
	cur := tr
	for _, c := range s {
		ne := c - 'a'
		if cur.son[ne] == nil {
			return false
		}
		cur = cur.son[ne]
	}
	return cur.end
}
