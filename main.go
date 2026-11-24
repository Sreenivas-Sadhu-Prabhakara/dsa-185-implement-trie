package main

import "fmt"

// Trie stores lowercase words for exact-word and prefix lookups.
type Trie struct {
	children [26]*Trie
	isWord   bool
}

// NewTrie returns an empty trie.
func NewTrie() *Trie { return &Trie{} }

// Insert adds a word (lowercase a-z).
func (t *Trie) Insert(word string) {
	n := t
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if n.children[c] == nil {
			n.children[c] = &Trie{}
		}
		n = n.children[c]
	}
	n.isWord = true
}

// walk descends along s; nil when the path doesn't exist.
func (t *Trie) walk(s string) *Trie {
	n := t
	for i := 0; i < len(s); i++ {
		n = n.children[s[i]-'a']
		if n == nil {
			return nil
		}
	}
	return n
}

// Search reports whether the exact word was inserted.
func (t *Trie) Search(word string) bool {
	n := t.walk(word)
	return n != nil && n.isWord
}

// StartsWith reports whether any inserted word begins with prefix.
func (t *Trie) StartsWith(prefix string) bool {
	return t.walk(prefix) != nil
}

func main() {
	tr := NewTrie()
	tr.Insert("apple")
	fmt.Println(tr.Search("apple"), tr.Search("app"), tr.StartsWith("app"))
}
