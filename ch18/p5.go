package main 

// Ended up mimicing the Dijkstra's algorithm solution and just giving each edge a weight of 1. 
// Book's solution is more elegant and can return early when the ending vertex is found by using Breadth-First search principles.

import (
	"fmt"
	"slices"
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

// shortest path function and helpers

func delete(slice []*Vertex, vertexToDelete *Vertex) []*Vertex {
	return slices.DeleteFunc(slice, func(vertex *Vertex) bool {
		return vertex.id == vertexToDelete.id
	})
}

func nextVertex(unvisitedVertices []*Vertex, shortestDistanceMap map[string]int) (next *Vertex) {
	if len(unvisitedVertices) == 0 {
		return nil
	}

	next = unvisitedVertices[0]

	for _, vertex := range unvisitedVertices[1:] {
		if shortestDistanceMap[next.value] > shortestDistanceMap[vertex.value] {
			next = vertex
		}
	}

	return
}

func shortestPath(startingVertex, endingVertex *Vertex) (shortestPath []string) {
	shortestDistanceMap := make(map[string]int)
	shortestPreviousVertexMap := make(map[string]string)

	unvisitedVertices := make([]*Vertex, 0)
	visitedVertices := make(map[string]bool)

	shortestDistanceMap[startingVertex.value] = 0
	currentVertex := startingVertex

	for {
		if currentVertex == nil {
			break
		}

		visitedVertices[currentVertex.value] = true
		unvisitedVertices = delete(unvisitedVertices, currentVertex)

		for _, adjacentVertex := range currentVertex.adjacentVertices {
			if visitedVertices[adjacentVertex.value] == false {
				unvisitedVertices = append(unvisitedVertices, adjacentVertex)
			}

			distanceThroughCurrentVertex := shortestDistanceMap[currentVertex.value] + 1

			if shortestDistanceMap[adjacentVertex.value] == 0 || distanceThroughCurrentVertex < shortestDistanceMap[adjacentVertex.value] {
				shortestDistanceMap[adjacentVertex.value] = distanceThroughCurrentVertex
				shortestPreviousVertexMap[adjacentVertex.value] = currentVertex.value
			}
		}

		currentVertex = nextVertex(unvisitedVertices, shortestDistanceMap)
	}

	currentVertexValue := endingVertex.value

	for {
		if currentVertexValue == startingVertex.value {
			break
		}

		shortestPath = append(shortestPath, currentVertexValue)
		currentVertexValue = shortestPreviousVertexMap[currentVertexValue]
	}

	shortestPath = append(shortestPath, startingVertex.value)

	slices.Reverse(shortestPath)

	return shortestPath
}

// vertices setup

func setupVertices() []*Vertex {
	idris := &Vertex{id: 1, value: "Idris"}
	kamil:= &Vertex{id: 2, value: "Kamil"}
	talia := &Vertex{id: 3, value: "Talia"}
	lina := &Vertex{id: 4, value: "Lina"}
	ken := &Vertex{id: 5, value: "Ken"}
	marco := &Vertex{id: 6, value: "Marco"}
	sasha := &Vertex{id: 7, value: "Sasha"}

	idris.addAdjacentVertices(kamil, talia)
	lina.addAdjacentVertices(kamil, sasha)
	marco.addAdjacentVertices(sasha, ken)
	talia.addAdjacentVertices(idris, ken)

	return []*Vertex{ idris, lina }
}

func main() {
	// setup vertices
	vertices := setupVertices()
	idris := vertices[0]
	lina := vertices[1]

	// invoke shortestPath function to find shortest path between idris and lina
	shortestPath := shortestPath(idris, lina)
	fmt.Println(shortestPath)
}