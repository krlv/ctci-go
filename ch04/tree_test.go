package ch04_test

import (
	"cmp"
	"fmt"
	"reflect"
	"testing"

	"github.com/krlv/ctci-go/ch04"
)

func TestInOrderTraversal(t *testing.T) {
	type walker[T cmp.Ordered] struct {
		walked []T
	}

	visitFunc := func(w *walker[int]) ch04.WalkFunc[int] {
		return func(n int) error {
			w.walked = append(w.walked, n)
			return nil
		}
	}

	visitFuncErr := func(w *walker[int]) ch04.WalkFunc[int] {
		return func(n int) error {
			w.walked = append(w.walked, n)
			return fmt.Errorf("error")
		}
	}

	type args[T cmp.Ordered] struct {
		n         *ch04.BinaryTreeNode[T]
		walker    *walker[T]
		visitFunc func(w *walker[T]) ch04.WalkFunc[T]
	}
	type testCase[T cmp.Ordered] struct {
		name    string
		args    args[T]
		want    []T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name: "empty tree (no tree)",
			args: args[int]{
				n:         nil,
				walker:    &walker[int]{},
				visitFunc: visitFunc,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "single root node",
			args: args[int]{
				n:         &ch04.BinaryTreeNode[int]{},
				walker:    &walker[int]{},
				visitFunc: visitFunc,
			},
			want:    []int{0},
			wantErr: false,
		},
		{
			name: "visit all nodes",
			args: args[int]{
				n: &ch04.BinaryTreeNode[int]{
					Data: 4,
					Left: &ch04.BinaryTreeNode[int]{
						Data:  2,
						Left:  &ch04.BinaryTreeNode[int]{Data: 1},
						Right: &ch04.BinaryTreeNode[int]{Data: 3},
					},
					Right: &ch04.BinaryTreeNode[int]{
						Data:  6,
						Left:  &ch04.BinaryTreeNode[int]{Data: 5},
						Right: &ch04.BinaryTreeNode[int]{Data: 7},
					},
				},
				walker:    &walker[int]{},
				visitFunc: visitFunc,
			},
			want:    []int{1, 2, 3, 4, 5, 6, 7},
			wantErr: false,
		},
		{
			name: "visit nodes with error",
			args: args[int]{
				n: &ch04.BinaryTreeNode[int]{
					Data: 4,
					Left: &ch04.BinaryTreeNode[int]{
						Data:  2,
						Left:  &ch04.BinaryTreeNode[int]{Data: 1},
						Right: &ch04.BinaryTreeNode[int]{Data: 3},
					},
					Right: &ch04.BinaryTreeNode[int]{
						Data:  6,
						Left:  &ch04.BinaryTreeNode[int]{Data: 5},
						Right: &ch04.BinaryTreeNode[int]{Data: 7},
					},
				},
				walker:    &walker[int]{},
				visitFunc: visitFuncErr,
			},
			want:    []int{1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ch04.InOrderTraversal(tt.args.n, tt.args.visitFunc(tt.args.walker)); (err != nil) != tt.wantErr {
				t.Errorf("InOrderTraversal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.want, tt.args.walker.walked) {
				t.Errorf("InOrderTraversal() = %v, want %v", tt.args.walker.walked, tt.want)
			}
		})
	}
}
