package ch04

import (
	"fmt"
	"strings"
)

// Vertex is a node in a graph
type Vertex[T comparable] struct {
	data     T
	adjacent []*Vertex[T]
}

// String returns the string representation of the vertex
func (v *Vertex[T]) String() string {
	return fmt.Sprintf("%v", v.data)
}

// Graph is a collection of vertices with edges between them
type Graph[T comparable] struct {
	vertices []*Vertex[T]
}

// NewGraph creates a new empty graph
func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{vertices: make([]*Vertex[T], 0)}
}

// Vertex returns the vertex with the given data
func (g *Graph[T]) Vertex(data T) (*Vertex[T], error) {
	for _, v := range g.vertices {
		if v.data == data {
			return v, nil
		}
	}

	return nil, fmt.Errorf("vertex %v not found", data)
}

// AddVertex adds a new vertex to the graph
func (g *Graph[T]) AddVertex(data T) (*Vertex[T], error) {
	if _, err := g.Vertex(data); err == nil {
		return nil, fmt.Errorf("vertex %v already exists", data)
	}

	v := &Vertex[T]{data: data}
	g.vertices = append(g.vertices, v)

	return v, nil
}

// AddEdge adds an edge between two vertices
func (g *Graph[T]) AddEdge(from, to T) error {
	if from == to {
		return fmt.Errorf("cannot add edge from %v to itself", from)
	}

	fromVertex, err := g.Vertex(from)
	if err != nil {
		return err
	}

	toVertex, err := g.Vertex(to)
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

	for _, v := range g.vertices {
		sb.WriteString(v.String())
		for _, a := range v.adjacent {
			sb.WriteString(" -> ")
			sb.WriteString(a.String())
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
