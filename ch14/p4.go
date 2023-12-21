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

func (l *LinkedList) last() int {
	currentNode := l.firstNode
	for {
		if currentNode.nextNode == nil { break }
		currentNode = currentNode.nextNode
	}

	return currentNode.data
}

func (l *LinkedList) reverse() {
	if l.firstNode.nextNode == nil { return }

	var previousNode *Node = nil
	var nextNode *Node
	
	for currentNode := l.firstNode; currentNode != nil; currentNode = nextNode {
		nextNode = currentNode.nextNode
		currentNode.nextNode = previousNode
		previousNode = currentNode
	}

	l.firstNode = previousNode
}

func main() {
	list := &LinkedList{}

	list.insertAtIndex(0, 8)
	list.insertAtIndex(0, 158)
	list.insertAtIndex(0, 90)
	list.insertAtIndex(0, 43)

	list.printAll()
	list.reverse()
	fmt.Println("--------")
	list.printAll()
}