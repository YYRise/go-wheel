package trie

type Node struct {
	isEnd    bool
	children map[rune]*Node
}

type Trie struct {
	root map[rune]*Node
}

func InitTrie(sl []string) (t *Trie) {
	t = &Trie{
		root: make(map[rune]*Node),
	}
	t.init(sl)
	return
}

func NewNode()(n *Node){
	n = &Node{
		children: make(map[rune]*Node),
	}
	return
}



func (t *Trie) init(sl []string) {
	for _, s := range sl {
		runes := []rune(s)
		l := len(runes)
		if l > 0 {
			w := runes[0]
			if _, ok := t.root[w]; !ok {
				t.root[w] = NewNode()
			}
			t.root[w].insert(runes[1:])
		}
	}
}

func (t *Trie) SearchOne(txt string) (match string) {
	words := []rune(txt)
	for i, w := range words {
		if tn, ok := t.root[w]; ok {
			if findIdx, ok := tn.search(words, i); ok {
				match = string(words[i : findIdx+1])
				return
			}
		}
	}
	return
}


func (tn *Node) insert(wl []rune) {
	node := tn
	for _, w := range wl {
		if _, ok := node.children[w]; !ok {
			node.children[w] = NewNode()
		}
		node = node.children[w]
	}
	node.isEnd = true
}

/*
	words：待匹配的文本
	idx: 搜索的起始位置
*/
func (tn *Node) search(words []rune, idx int) (findIdx int, ok bool) {
	findIdx = -1
	if tn.isEnd {
		return idx, true
	}
	idx++
	if idx >= len(words) {
		return
	}
	w := words[idx]
	if c, y := tn.children[w]; !y {
		return
	} else {
		findIdx, ok = c.search(words, idx)
	}
	return
}

func (tn *Node) remove(key []rune)(delKey bool){
	if len(key) == 0 {
		if tn.isEnd {
			tn.isEnd = false
		}
		if len(tn.children) == 0{
			delKey = true
			return
		}
	}

	k := key[0]
	if node, ok := tn.children[k]; !ok{
		return
	}else{
		delKey = node.remove(key[1:]) && len(node.children) == 1
		if delKey{
			delete(node.children, k)
		}
	}
	return
}

