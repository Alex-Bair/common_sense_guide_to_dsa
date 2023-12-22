package main 

import (
	"fmt"
)

// DoublyLinkedList struct and related functions

type Node struct {
	data *Vertex
	nextNode *Node
	previousNode *Node
}

type DoublyLinkedList struct {
	firstNode *Node
	lastNode *Node
}

func (l *DoublyLinkedList) removeFromFront() *Node {
	removedNode := l.firstNode
	l.firstNode = l.firstNode.nextNode
	return removedNode
}

func (l *DoublyLinkedList) insertAtEnd(value *Vertex) {
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

// Queue struct

type Queue struct {
	data *DoublyLinkedList
}

func (q *Queue) initialize() {
	doublyLinkedList := &DoublyLinkedList{}
	q.data = doublyLinkedList
}

func (q *Queue) enqueue(element *Vertex) {
	q.data.insertAtEnd(element)
}

func (q *Queue) dequeue() *Vertex {
	removedNode := q.data.removeFromFront()
	return removedNode.data
}

func (q *Queue) read() *Vertex {
	if q.data.firstNode == nil {
		return nil
	}

	return q.data.firstNode.data
}

// Graph struct and related functions

type Vertex struct {
	id int
	value string
	adjacentVertices []*Vertex
}

func (v *Vertex) hasAdjacent(vertex *Vertex) bool {
	for _, currentVertex := range v.adjacentVertices {
		if currentVertex.id == vertex.id {
			return true
		}
	}

	return false
}

func (v *Vertex) addAdjacentVertex(vertex *Vertex) {
	if v.hasAdjacent(vertex) {
		return
	}

	v.adjacentVertices = append(v.adjacentVertices, vertex)
	vertex.adjacentVertices = append(vertex.adjacentVertices, v)
}

func (v *Vertex) addAdjacentVertices(vertices ...*Vertex) {
	for _, vertex := range vertices {
		v.addAdjacentVertex(vertex)
	}
}

// Breadth-First Traversal and Search

func bfsTraverse(startingVertex *Vertex) {
	queue := &Queue{}
	queue.initialize()

	visitedVertices := make(map[*Vertex]bool)
	visitedVertices[startingVertex] = true
	queue.enqueue(startingVertex)

	for {
		if queue.read() == nil { 
			break
		}

		currentVertex := queue.dequeue()
		fmt.Println(currentVertex.id, currentVertex.value)

		for _, adjacentVertex := range currentVertex.adjacentVertices {
			if visitedVertices[adjacentVertex] == false {
				visitedVertices[adjacentVertex] = true
				queue.enqueue(adjacentVertex)
			}
		}
	}
}

func bfsSearch(startingVertex *Vertex, searchValue string) *Vertex {
	queue := &Queue{}
	queue.initialize()

	visitedVertices := make(map[*Vertex]bool)
	visitedVertices[startingVertex] = true
	queue.enqueue(startingVertex)

	for {
		if queue.read() == nil { 
			break
		}

		currentVertex := queue.dequeue()
		if currentVertex.value == searchValue {
			return currentVertex
		}

		for _, adjacentVertex := range currentVertex.adjacentVertices {
			if visitedVertices[adjacentVertex] == false {
				visitedVertices[adjacentVertex] = true
				queue.enqueue(adjacentVertex)
			}
		}
	}

	return &Vertex{id: 0, value: "Value not found"}
}

func setupVertices() *Vertex {
	luffy := &Vertex{id: 1, value: "Luffy"}
	zoro:= &Vertex{id: 2, value: "Zoro"}
	nami := &Vertex{id: 3, value: "Nami"}
	usopp := &Vertex{id: 4, value: "Usopp"}
	sanji := &Vertex{id: 5, value: "Sanji"}
	chopper := &Vertex{id: 6, value: "Chopper"}
	robin := &Vertex{id: 7, value: "Robin"}
	franky := &Vertex{id: 8, value: "Franky"}
	brook := &Vertex{id: 9, value: "Brook"}
	ace := &Vertex{id: 10, value: "Ace"}
	hancock := &Vertex{id: 11, value: "Hancock"}
	buggy := &Vertex{id: 12, value: "Buggy"}
	shanks := &Vertex{id: 13, value: "Shanks"}
	perona := &Vertex{id: 14, value: "Perona"}
	law := &Vertex{id: 15, value: "Law"}
	whitebeard := &Vertex{id: 16, value: "Whitebeard"}

	luffy.addAdjacentVertices(zoro, nami, usopp, ace, hancock, shanks, buggy, law)
	zoro.addAdjacentVertices(sanji, robin, perona, brook)
	nami.addAdjacentVertices(sanji, robin, chopper)
	chopper.addAdjacentVertices(usopp)
	ace.addAdjacentVertices(luffy, whitebeard)
	whitebeard.addAdjacentVertices(ace, shanks)
	franky.addAdjacentVertices(chopper, robin)

	return luffy
}

func main() {
	// setup vertices
	luffy := setupVertices()

	vertex := bfsSearch(luffy, "Franky")
	fmt.Println(vertex.value)
}