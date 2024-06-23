package trie

import (
	"fmt"
)

type Trie interface {
	Insert(word string)
	Search(word string) bool
}

func NewTrieSlice() *NodeSlice {
	return &NodeSlice{
		children: [26]*NodeSlice{},
	}
}

func NewTrieMap() *NodeMap {
	return &NodeMap{
		children: make(map[rune]*NodeMap),
	}
}

func NewTrieMapNoAlloc() *NodeMapNoAlloc {
	return &NodeMapNoAlloc{
		children: make(map[rune]NodeMapNoAlloc),
	}
}

type NodeSlice struct {
	children [26]*NodeSlice
	wordEnd  bool
}

func (ns *NodeSlice) Print() {
	ns.print("")
}

func (ns *NodeSlice) print(acc string) {
	if ns.wordEnd {
		fmt.Println(acc)
	}
	for i, child := range ns.children {
		if child != nil {
			child.print(acc + string(byte('a'+i)))
		}
	}
}

func (ns *NodeSlice) Insert(word string) {
	cur := ns

	for _, c := range word {
		idx := c - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = &NodeSlice{
				children: [26]*NodeSlice{},
			}
		}
		cur = cur.children[idx]
	}

	cur.wordEnd = true
}

func (ns *NodeSlice) Search(word string) bool {
	cur := ns

	for _, c := range word {
		idx := c - 'a'
		if cur.children[idx] == nil {
			return false
		}
		cur = cur.children[idx]
	}

	return cur.wordEnd
}

type NodeMap struct {
	children map[rune]*NodeMap
	wordEnd  bool
}

func (nm NodeMap) Print() {
	nm.print("")
}

func (nm NodeMap) print(acc string) {
	if nm.wordEnd {
		fmt.Println(acc)
	}
	for c, child := range nm.children {
		child.print(acc + string(byte(c)))
	}
}

func (nm *NodeMap) Insert(word string) {
	cur := nm

	// var lastRune rune

	for _, c := range word {
		child, found := cur.children[c]
		if !found {
			child = &NodeMap{
				children: make(map[rune]*NodeMap),
			}
			cur.children[c] = child
		}
		cur = child
		// last = c
	}

	cur.wordEnd = true

}

func (nm *NodeMap) Search(word string) bool {
	cur := nm

	for _, c := range word {
		child, found := cur.children[c]
		if !found {
			return false
		}
		cur = child
	}

	return cur.wordEnd
}

type NodeMapNoAlloc struct {
	children map[rune]NodeMapNoAlloc
	wordEnd  bool
}

func (nm NodeMapNoAlloc) Print() {
	nm.print("")
}

func (nm NodeMapNoAlloc) print(acc string) {
	if nm.wordEnd {
		fmt.Println(acc)
	}
	for c, child := range nm.children {
		child.print(acc + string(byte(c)))
	}
}

func (nm *NodeMapNoAlloc) Insert(word string) {
	*nm = insert(*nm, []rune(word))
}

func insert(node NodeMapNoAlloc, word []rune) NodeMapNoAlloc {
	char := word[0]

	child, found := node.children[char]
	if !found {
		child = NodeMapNoAlloc{
			children: make(map[rune]NodeMapNoAlloc),
		}
	}

	if len(word) > 1 {
		node.children[char] = insert(child, word[1:])
		return node
	} else {
		child.wordEnd = true
		node.children[char] = child
		return node
	}
}

func (nm NodeMapNoAlloc) Search(word string) bool {
	cur := nm

	for _, c := range word {
		child, found := cur.children[c]
		if !found {
			return false
		}
		cur = child
	}

	return cur.wordEnd
}
