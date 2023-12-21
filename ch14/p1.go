package main

import "fmt"

type Node struct {
	data int
	nextNode *Node
}

type LinkedList struct {
	firstNode *Node
}

func (l *LinkedList) read(index int) int {
	currentNode := l.firstNode

	for currentIndex := 0; currentIndex < index; currentIndex++ {
		currentNode = currentNode.nextNode

		if currentNode == nil {return -1}
	}

	return currentNode.data
}

func (l *LinkedList) indexOf(value int) int {
	currentNode := l.firstNode

	for currentIndex := 0; currentNode != nil; currentIndex++ {
		if currentNode.data == value { return currentIndex }
		currentNode = currentNode.nextNode
	}

	return -1
}

func (l *LinkedList) insertAtIndex(index, value int) {
	newNode := &Node{data: value}

	if index == 0 {
		newNode.nextNode = l.firstNode
		l.firstNode = newNode
		return
	}

	currentNode := l.firstNode
	
	for currentIndex := 0; currentIndex < index - 1; currentIndex++ {
		currentNode = currentNode.nextNode
	}
	
	newNode.nextNode = currentNode.nextNode
	currentNode.nextNode = newNode
}

func (l *LinkedList) deleteAtIndex(index int) {
	if index == 0 {
		l.firstNode = l.firstNode.nextNode
		return
	}

	currentNode := l.firstNode

	for currentIndex := 0; currentIndex < index - 1; currentIndex++ {
		currentNode = currentNode.nextNode
	}

	nodeAfterDeletedNode := currentNode.nextNode.nextNode
	currentNode.nextNode = nodeAfterDeletedNode
}

func (l *LinkedList) printAll() {
	currentNode := l.firstNode
	for {
		if currentNode == nil { break }
		fmt.Println(currentNode.data)
		currentNode = currentNode.nextNode
	}
}

func main() {
	node1 := &Node{data: 43}
	node2 := &Node{data: 90}
	node3 := &Node{data: 158}
	node4 := &Node{data: 8}
	node1.nextNode = node2
	node2.nextNode = node3
	node3.nextNode = node4
	
	list := &LinkedList{firstNode: node1}

	list.printAll()
}