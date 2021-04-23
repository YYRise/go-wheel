package trie

import (
	"fmt"
	"testing"
)

func TestTrie_SearchOne(t *testing.T) {
	words := []string{"我们", "大", "二狗"}
	txt := "aa,谁说大家是二狗"
	txtSl := []rune(txt)
	fmt.Println("txtSl = ", txtSl)
	tr := InitTrie(words)
	match := tr.SearchOne(txt)
	fmt.Println("match = ", match)
}

func TestInitTrie(t *testing.T) {
	txt := "aa,谁说大家是二,狗"
	for _, c := range txt{
		fmt.Println(c)
	}
}

