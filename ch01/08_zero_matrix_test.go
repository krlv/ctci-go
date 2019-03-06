package ch01

import "testing"

func TestZeroMatrix(t *testing.T) {
	example := [][]int{}
	expected := [][]int{}
	ZeroMatrix(example)
	for i, row := range example {
		for j, col := range row {
			if expected[i][j] != col {
				t.Errorf("Expected zero matrix to be %v, got %v", expected, example)
			}
		}
	}

	example = [][]int{
		{},
	}
	expected = [][]int{
		{},
	}
	ZeroMatrix(example)
	for i, row := range example {
		for j, col := range row {
			if expected[i][j] != col {
				t.Errorf("Expected zero matrix to be %v, got %v", expected, example)
			}
		}
	}

	example = [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 0, 5, 6},
		{1, 2, 3, 4, 5, 6},
	}
	expected = [][]int{
		{1, 2, 3, 0, 5, 6},
		{0, 0, 0, 0, 0, 0},
		{1, 2, 3, 0, 5, 6},
	}
	ZeroMatrix(example)
	for i, row := range example {
		for j, col := range row {
			if expected[i][j] != col {
				t.Errorf("Expected zero matrix to be %v, got %v", expected, example)
			}
		}
	}

	example = [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, 3, 0, 5, 0, 7},
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, 3, 4, 5, 6, 0},
	}
	expected = [][]int{
		{1, 2, 3, 0, 5, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 2, 3, 0, 5, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	ZeroMatrix(example)
	for i, row := range example {
		for j, col := range row {
			if expected[i][j] != col {
				t.Errorf("Expected zero matrix to be %v, got %v", expected, example)
			}
		}
	}
}
