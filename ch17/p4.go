package main

import "fmt"

type TrieNode struct {
	children map[string]*TrieNode
}

type Trie struct {
	root *TrieNode
}

type Words struct {
	words []string
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
	if currentNode == nil {
		return
	}

	for key, child := range currentNode.children {
		fmt.Println(key)
		printAllKeys(child)
	}
}

func (t *Trie) collectAllWords(currentNode *TrieNode) []string {
	collection := &Words{
		words: make([]string, 0),
	}

	if currentNode == nil {
		currentNode = t.root
	}

	collectAllWords(currentNode, "", collection)

	return collection.words
}

func collectAllWords(currentNode *TrieNode, word string, collection *Words) {
	for key, childNode := range currentNode.children {
		if key == "*" {
			collection.words = append(collection.words, word)
		} else {
			collectAllWords(childNode, word + key, collection)
		}
	}
}

func (t *Trie) autocomplete(prefix string) []string {
	currentNode := t.search(prefix)
	if currentNode == nil {
		return nil
	}

	return t.collectAllWords(currentNode)
}

func (t *Trie) autocorrect(word string) string {
	currentNode := t.root

	wordSoFar := ""

	for _, runeValue := range word {
		char := string(runeValue)
		child, present := currentNode.children[char]
		
		if present {
			wordSoFar += char
			currentNode = child
		} else {
			return wordSoFar + t.collectAllWords(currentNode)[0]
		}
	}

	return word
}

func main() {
	root := &TrieNode{children: make(map[string]*TrieNode)}
	trie := &Trie{root: root}

	words := []string{"cat", "catnap", "catnip"}
	for _, word := range words {
		trie.insert(word)
	}

	fmt.Println(trie.autocorrect("catnar"))
	fmt.Println(trie.autocorrect("caxasfdij"))
	fmt.Println(trie.autocorrect("ca"))
}