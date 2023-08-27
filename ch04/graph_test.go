package ch04_test

import (
	"reflect"
	"testing"

	"github.com/krlv/ctci-go/ch04"
)

func TestGraph_Vertex(t *testing.T) {
	emptyGraph := ch04.NewGraph[string]()

	graphWithA := ch04.NewGraph[string]()
	vertexA, _ := graphWithA.AddVertex("A")

	type args[T comparable] struct {
		data T
	}
	type testCase[T comparable] struct {
		name    string
		g       *ch04.Graph[T]
		args    args[T]
		want    *ch04.Vertex[T]
		wantErr bool
	}
	tests := []testCase[string]{
		{
			name:    "vertex exists",
			g:       graphWithA,
			args:    args[string]{data: "A"},
			want:    vertexA,
			wantErr: false,
		},
		{
			name:    "vertex does not exist",
			g:       emptyGraph,
			args:    args[string]{data: "A"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.Vertex(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Vertex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vertex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_AddVertex(t *testing.T) {
	emptyGraph := ch04.NewGraph[string]()

	graphWithA := ch04.NewGraph[string]()
	graphWithA.AddVertex("A")

	type args[T comparable] struct {
		data T
	}
	type testCase[T comparable] struct {
		name    string
		g       *ch04.Graph[T]
		args    args[T]
		want    string
		wantErr bool
	}
	tests := []testCase[string]{
		{
			name:    "add vertex to empty graph",
			g:       emptyGraph,
			args:    args[string]{data: "A"},
			want:    "A",
			wantErr: false,
		},
		{
			name:    "fail to add vertex that already exists",
			g:       graphWithA,
			args:    args[string]{data: "A"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.AddVertex(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddVertex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var vs string
			if got != nil {
				vs = got.String()
			}

			if !reflect.DeepEqual(vs, tt.want) {
				t.Errorf("AddVertex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_AddEdge(t *testing.T) {
	emptyGraph := ch04.NewGraph[string]()

	graphWithA := ch04.NewGraph[string]()
	graphWithA.AddVertex("A")

	graphWithAB := ch04.NewGraph[string]()
	graphWithAB.AddVertex("A")
	graphWithAB.AddVertex("B")

	graphWithEdge := ch04.NewGraph[string]()
	graphWithEdge.AddVertex("A")
	graphWithEdge.AddVertex("B")
	graphWithEdge.AddEdge("A", "B")

	type args[T comparable] struct {
		from T
		to   T
	}
	type testCase[T comparable] struct {
		name    string
		g       *ch04.Graph[T]
		args    args[T]
		wantErr bool
	}
	tests := []testCase[string]{
		{
			name:    "add edge to empty graph",
			g:       emptyGraph,
			args:    args[string]{from: "A", to: "B"},
			wantErr: true,
		},
		{
			name:    "add edge to graph with one vertex",
			g:       graphWithA,
			args:    args[string]{from: "A", to: "B"},
			wantErr: true,
		},
		{
			name:    "add edge to graph",
			g:       graphWithAB,
			args:    args[string]{from: "A", to: "B"},
			wantErr: false,
		},
		{
			name:    "add duplicate edge to graph",
			g:       graphWithEdge,
			args:    args[string]{from: "A", to: "B"},
			wantErr: true,
		},
		{
			name:    "add edge to graph with self-loop",
			g:       graphWithA,
			args:    args[string]{from: "A", to: "A"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.AddEdge(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("AddEdge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGraph_String(t *testing.T) {
	emptyGraph := ch04.NewGraph[string]()

	graphWithA := ch04.NewGraph[string]()
	graphWithA.AddVertex("A")

	graphWithAB := ch04.NewGraph[string]()
	graphWithAB.AddVertex("A")
	graphWithAB.AddVertex("B")

	graphWithEdge := ch04.NewGraph[string]()
	graphWithEdge.AddVertex("A")
	graphWithEdge.AddVertex("B")
	graphWithEdge.AddEdge("A", "B")
	graphWithEdge.AddVertex("C")
	graphWithEdge.AddVertex("D")
	graphWithEdge.AddEdge("C", "D")
	graphWithEdge.AddEdge("D", "A")

	type testCase[T comparable] struct {
		name string
		g    *ch04.Graph[T]
		want string
	}
	tests := []testCase[string]{
		{
			name: "empty graph",
			g:    emptyGraph,
			want: "",
		},
		{
			name: "graph with one vertex",
			g:    graphWithA,
			want: "A\n",
		},
		{
			name: "graph with two vertices",
			g:    graphWithAB,
			want: "A\nB\n",
		},
		{
			name: "graph with edges",
			g:    graphWithEdge,
			want: "A -> B\nB\nC -> D\nD -> A\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
