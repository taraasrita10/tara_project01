package trie

import (
	"slices"
)

type Node struct {
	Children map[rune]*Node
	IsEnd    bool
}

func NewNode() *Node {
	return &Node{
		Children: make(map[rune]*Node),
		IsEnd:    false,
	}
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

func (t *Trie) Add(word string) {
	currNode := t.root
	for _, char := range word {
		_, exists := currNode.Children[char]
		if !exists {
			currNode.Children[char] = NewNode()
		}
		currNode = currNode.Children[char]
	}
	currNode.IsEnd = true
}

func (t *Trie) Check(word string) bool {
	currNode := t.root
	for _, char := range word {
		_, exists := currNode.Children[char]
		if !exists {
			return false
		}
		currNode = currNode.Children[char]
	}
	return currNode.IsEnd
}

func (t *Trie) RemoveHelper(currNode *Node, word string, index int) bool {
	if index == len(word) {
		if !currNode.IsEnd {
			return false
		}
		currNode.IsEnd = false
		return len(currNode.Children) == 0
	}
	char := rune(word[index])
	nextNode, exists := currNode.Children[char]
	if !exists {
		return false
	}
	deleteChild := t.RemoveHelper(nextNode, word, index+1)
	if deleteChild {
		delete(currNode.Children, char)
		return len(currNode.Children) == 0 && currNode.IsEnd
	}

	return false
}

func (t *Trie) Remove(word string) {
	t.RemoveHelper(t.root, word, 0)
}

func (t *Trie) ListHelper(currNode *Node, prefix []rune, results *[]string) {
	if currNode.IsEnd {
		*results = append(*results, string(prefix))
	}
	for char, nextNode := range currNode.Children {
		newPrefix := append(prefix, char)
		t.ListHelper(nextNode, newPrefix, results)
	}
}

func (t *Trie) List() []string {
	results := []string{}
	t.ListHelper(t.root, []rune{}, &results)
	slices.Sort(results)
	return results
}
