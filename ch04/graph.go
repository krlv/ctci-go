package ch04

import (
	"fmt"
	"strings"
)

// GraphNode is a node (vertex) in a graph
type GraphNode[T comparable] struct {
	data     T
	adjacent []*GraphNode[T]
}

// String returns the string representation of the graph node
func (v *GraphNode[T]) String() string {
	return fmt.Sprintf("%v", v.data)
}

// Graph is a collection of nodes (vertices) with edges between them
type Graph[T comparable] struct {
	nodes []*GraphNode[T]
}

// NewGraph creates a new empty graph
func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{nodes: make([]*GraphNode[T], 0)}
}

// Node returns the node (vertex) with the given data
func (g *Graph[T]) Node(data T) (*GraphNode[T], error) {
	for _, v := range g.nodes {
		if v.data == data {
			return v, nil
		}
	}

	return nil, fmt.Errorf("node %v not found", data)
}

// AddNode adds a new node (vertex) to the graph
func (g *Graph[T]) AddNode(data T) (*GraphNode[T], error) {
	if _, err := g.Node(data); err == nil {
		return nil, fmt.Errorf("node %v already exists", data)
	}

	v := &GraphNode[T]{data: data}
	g.nodes = append(g.nodes, v)

	return v, nil
}

// AddEdge adds an edge between two nodes
func (g *Graph[T]) AddEdge(from, to T) error {
	if from == to {
		return fmt.Errorf("cannot add edge from %v to itself", from)
	}

	fromVertex, err := g.Node(from)
	if err != nil {
		return err
	}

	toVertex, err := g.Node(to)
	if err != nil {
		return err
	}

	for _, v := range fromVertex.adjacent {
		if v == toVertex {
			return fmt.Errorf("edge from %v to %v already exists", from, to)
		}
	}

	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)

	return nil
}

// String returns the string representation of the graph
func (g *Graph[T]) String() string {
	var sb strings.Builder

	for _, v := range g.nodes {
		sb.WriteString(v.String())
		for _, a := range v.adjacent {
			sb.WriteString(" -> ")
			sb.WriteString(a.String())
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
