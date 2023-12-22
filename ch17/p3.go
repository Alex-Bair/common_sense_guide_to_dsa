package main

import "fmt"

type TrieNode struct {
	children map[string]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) search(word string) *TrieNode {
	return search(t.root, word)
}

func search(currentNode *TrieNode, word string) *TrieNode {
	for _, runeValue := range word {
		children, present := currentNode.children[string(runeValue)]
		if present {
			currentNode = children
		} else {
			return nil
		}
	}

	return currentNode
}

func (t *Trie) insert(word string) {
	insert(t.root, word)
}

func insert(currentNode *TrieNode, word string) {
	for _, runeValue := range word {
		children, present := currentNode.children[string(runeValue)]
		if present {
			currentNode = children
		} else {
			newNode := &TrieNode{children: make(map[string]*TrieNode)}
			currentNode.children[string(runeValue)] = newNode
			currentNode = newNode
		}
	}

	currentNode.children["*"] = nil
}

func (t *Trie) printAllKeys() {
	printAllKeys(t.root)
}

func printAllKeys(currentNode *TrieNode) {
	for key, child := range currentNode.children {
		fmt.Println(key)
		if key != "*" {
			printAllKeys(child)
		}
	}
}

func main() {
	root := &TrieNode{children: make(map[string]*TrieNode)}
	trie := &Trie{root: root}

	trie.insert("apple")
	trie.insert("bad")
	trie.insert("bat")
	trie.insert("cat")
	trie.insert("can")

	trie.printAllKeys()

}