package ch04_test

import (
	"testing"

	"github.com/krlv/ctci-go/ch04"
)

func TestHasRouteBetweenNodes(t *testing.T) {
	graphWithPath := ch04.NewGraph[string]()
	vertexPathA, _ := graphWithPath.AddNode("A")
	graphWithPath.AddNode("B")
	graphWithPath.AddNode("C")
	graphWithPath.AddNode("D")
	graphWithPath.AddNode("E")
	graphWithPath.AddNode("F")
	vertexPathG, _ := graphWithPath.AddNode("G")
	graphWithPath.AddEdge("A", "B")
	graphWithPath.AddEdge("A", "C")
	graphWithPath.AddEdge("B", "D")
	graphWithPath.AddEdge("B", "E")
	graphWithPath.AddEdge("C", "F")
	graphWithPath.AddEdge("E", "G")

	graphWithoutPath := ch04.NewGraph[string]()
	vertexNoPathA, _ := graphWithoutPath.AddNode("A")
	graphWithoutPath.AddNode("B")
	graphWithoutPath.AddNode("C")
	graphWithoutPath.AddNode("D")
	graphWithoutPath.AddNode("E")
	graphWithoutPath.AddNode("F")
	vertexNoPathG, _ := graphWithoutPath.AddNode("G")
	graphWithoutPath.AddEdge("A", "B")
	graphWithoutPath.AddEdge("A", "C")
	graphWithoutPath.AddEdge("B", "D")
	graphWithoutPath.AddEdge("B", "E")
	graphWithoutPath.AddEdge("C", "F")

	type args[T comparable] struct {
		g     *ch04.Graph[T]
		start *ch04.GraphNode[T]
		end   *ch04.GraphNode[T]
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[string]{
		{
			name: "path exists",
			args: args[string]{
				g:     graphWithPath,
				start: vertexPathA,
				end:   vertexPathG,
			},
			want: true,
		},
		{
			name: "path does not exist",
			args: args[string]{
				g:     graphWithoutPath,
				start: vertexNoPathA,
				end:   vertexNoPathG,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ch04.HasRouteBetweenNodes(tt.args.g, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("HasRouteBetweenNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
