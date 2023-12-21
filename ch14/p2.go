package main

import "fmt"

type Node struct {
	data int
	nextNode *Node
	previousNode *Node
}

type DoublyLinkedList struct {
	firstNode *Node
	lastNode *Node
}

func (l *DoublyLinkedList) read(index int) int {
	currentNode := l.firstNode

	for currentIndex := 0; currentIndex < index; currentIndex++ {
		currentNode = currentNode.nextNode

		if currentNode == nil {return -1}
	}

	return currentNode.data
}

func (l *DoublyLinkedList) indexOf(value int) int {
	currentNode := l.firstNode

	for currentIndex := 0; currentNode != nil; currentIndex++ {
		if currentNode.data == value { return currentIndex }
		currentNode = currentNode.nextNode
	}

	return -1
}

func (l *DoublyLinkedList) insertAtIndex(index, value int) {
	newNode := &Node{data: value}

	if index == 0 {
		l.firstNode.previousNode = newNode
		newNode.nextNode = l.firstNode
		l.firstNode = newNode
		return
	}

	currentNode := l.firstNode
	
	for currentIndex := 0; currentIndex < index - 1; currentIndex++ {
		currentNode = currentNode.nextNode
	}
	
	newNode.previousNode = currentNode
	newNode.nextNode = currentNode.nextNode
	currentNode.nextNode = newNode

	if currentNode == l.lastNode {
		l.lastNode = newNode
	}
}

func (l *DoublyLinkedList) deleteAtIndex(index int) {
	if index == 0 {
		l.firstNode = l.firstNode.nextNode
		l.firstNode.previousNode = nil
		return
	}

	currentNode := l.firstNode

	for currentIndex := 0; currentIndex < index - 1; currentIndex++ {
		currentNode = currentNode.nextNode
	}

	if currentNode.nextNode == l.lastNode {
		l.lastNode = currentNode
	}

	nodeAfterDeletedNode := currentNode.nextNode.nextNode
	currentNode.nextNode = nodeAfterDeletedNode

	if nodeAfterDeletedNode != nil {
		nodeAfterDeletedNode.previousNode = currentNode
	}
}

func (l *DoublyLinkedList) printAll() {
	currentNode := l.firstNode
	for {
		if currentNode == nil { break }
		fmt.Println(currentNode.data)
		currentNode = currentNode.nextNode
	}
}

func (l *DoublyLinkedList) reversePrintAll() {
	currentNode := l.lastNode
	for {
		if currentNode == nil { break }
		fmt.Println(currentNode.data)
		currentNode = currentNode.previousNode
	}
}

func (l *DoublyLinkedList) insertAtEnd(value int) {
	newNode := &Node{data: value}

	if l.firstNode == nil {
		l.firstNode = newNode
		l.lastNode = newNode
	} else {
		newNode.previousNode = l.lastNode
		l.lastNode.nextNode = newNode
		l.lastNode = newNode
	}
}

func main() {
	list := &DoublyLinkedList{}
	list.insertAtEnd(43)
	list.insertAtEnd(8)
	list.insertAtEnd(27)
	list.insertAtEnd(986)

	list.printAll()
	fmt.Println("----")
	list.reversePrintAll()
}