package ch01

import "testing"

func TestZeroMatrix(t *testing.T) {
	actual := [][]int{}
	expected := [][]int{}
	ZeroMatrix(actual)
	if !isEqualMatrix(actual, expected) {
		t.Errorf("Expected zero matrix to be %v, got %v", expected, actual)
	}

	actual = [][]int{
		{},
	}
	expected = [][]int{
		{},
	}
	ZeroMatrix(actual)
	if !isEqualMatrix(actual, expected) {
		t.Errorf("Expected zero matrix to be %v, got %v", expected, actual)
	}

	actual = [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 0, 5, 6},
		{1, 2, 3, 4, 5, 6},
	}
	expected = [][]int{
		{1, 2, 3, 0, 5, 6},
		{0, 0, 0, 0, 0, 0},
		{1, 2, 3, 0, 5, 6},
	}
	ZeroMatrix(actual)
	if !isEqualMatrix(actual, expected) {
		t.Errorf("Expected zero matrix to be %v, got %v", expected, actual)
	}

	actual = [][]int{
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
	ZeroMatrix(actual)
	if !isEqualMatrix(actual, expected) {
		t.Errorf("Expected zero matrix to be %v, got %v", expected, actual)
	}
}

func TestZeroMatrixSpaceEfficient(t *testing.T) {
	actual := [][]int{}
	expected := [][]int{}
	ZeroMatrixSpaceEfficient(actual)
	if !isEqualMatrix(actual, expected) {
		t.Errorf("Expected zero matrix to be %v, got %v", expected, actual)
	}

	actual = [][]int{
		{},
	}
	expected = [][]int{
		{},
	}
	ZeroMatrixSpaceEfficient(actual)
	if !isEqualMatrix(actual, expected) {
		t.Errorf("Expected zero matrix to be %v, got %v", expected, actual)
	}

	actual = [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 0, 5, 6},
		{1, 2, 3, 4, 5, 6},
	}
	expected = [][]int{
		{1, 2, 3, 0, 5, 6},
		{0, 0, 0, 0, 0, 0},
		{1, 2, 3, 0, 5, 6},
	}
	ZeroMatrixSpaceEfficient(actual)
	if !isEqualMatrix(actual, expected) {
		t.Errorf("Expected zero matrix to be %v, got %v", expected, actual)
	}

	actual = [][]int{
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
	ZeroMatrixSpaceEfficient(actual)
	if !isEqualMatrix(actual, expected) {
		t.Errorf("Expected zero matrix to be %v, got %v", expected, actual)
	}
}

func isEqualMatrix(actual [][]int, expected [][]int) bool {
	for i, row := range actual {
		for j, col := range row {
			if expected[i][j] != col {
				return false
			}
		}
	}
	return true
}
