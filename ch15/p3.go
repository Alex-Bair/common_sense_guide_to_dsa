package main

import "fmt"

type Node struct {
	value int
	leftChild *Node
	rightChild *Node
}

type BinarySearchTree struct {
	root *Node
}

func (tree *BinarySearchTree) insert(value int) {
	if tree.root == nil {
		tree.root = &Node{value: value}
	} else {
		insert(value, tree.root)
	}
}

func insert(value int, node *Node) {
	if value < node.value {
		if node.leftChild == nil {
			node.leftChild = &Node{value: value}
		} else {
			insert(value, node.leftChild)
		}	
	} else if value > node.value {
		if node.rightChild == nil {
			node.rightChild = &Node{value: value}
		} else {
			insert(value, node.rightChild)
		}
	}
}

func (tree *BinarySearchTree) traverseAndPrint() {
	if tree.root == nil { return }
	traverseAndPrint(tree.root)
}

func traverseAndPrint(node *Node) {
	if node == nil { return }
	traverseAndPrint(node.leftChild)
	fmt.Println(node.value)
	traverseAndPrint(node.rightChild)
}

func (tree *BinarySearchTree) max() int {
	if tree.root == nil { return 0 }
	return max(tree.root)
}

func max(node *Node) int {
	if node.rightChild == nil {
		return node.value
	} else {
		return max(node.rightChild)
	}
}

func main() {
	tree := &BinarySearchTree{}
	values := []int{7, 2, 8, 5, 8, 3, 12, 7, 45, 2, 0, 1, 100}

	for _, val := range values {
		tree.insert(val)
	}

	fmt.Println(tree.max())
}