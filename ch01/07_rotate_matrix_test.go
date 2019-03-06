package ch01

import (
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2},
		{10, 11, 12},
		{20, 21, 22},
	}
	expected := [][]int{
		{20, 10, 0},
		{21, 11, 1},
		{22, 12, 2},
	}
	actual := RotateMatrix(matrix)

	for i, row := range matrix {
		for j := range row {
			if actual[i][j] != expected[i][j] {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		}
	}

	matrix = [][]int{
		{0, 1, 2, 3, 4, 5},
		{10, 11, 12, 13, 14, 15},
		{20, 21, 22, 23, 24, 25},
		{30, 31, 32, 33, 34, 35},
		{40, 41, 42, 43, 44, 45},
		{50, 51, 52, 53, 54, 55},
	}
	expected = [][]int{
		{50, 40, 30, 20, 10, 0},
		{51, 41, 31, 21, 11, 1},
		{52, 42, 32, 22, 12, 2},
		{53, 43, 33, 23, 13, 3},
		{54, 44, 34, 24, 14, 4},
		{55, 45, 35, 25, 15, 5},
	}
	actual = RotateMatrix(matrix)

	for i, row := range matrix {
		for j := range row {
			if actual[i][j] != expected[i][j] {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		}
	}
}

func TestRotateMatrixEfficient(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2},
		{10, 11, 12},
		{20, 21, 22},
	}
	expected := [][]int{
		{20, 10, 0},
		{21, 11, 1},
		{22, 12, 2},
	}
	actual := RotateMatrixEfficient(matrix)

	for i, row := range matrix {
		for j := range row {
			if actual[i][j] != expected[i][j] {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		}
	}

	matrix = [][]int{
		{0, 1, 2, 3, 4, 5},
		{10, 11, 12, 13, 14, 15},
		{20, 21, 22, 23, 24, 25},
		{30, 31, 32, 33, 34, 35},
		{40, 41, 42, 43, 44, 45},
		{50, 51, 52, 53, 54, 55},
	}
	expected = [][]int{
		{50, 40, 30, 20, 10, 0},
		{51, 41, 31, 21, 11, 1},
		{52, 42, 32, 22, 12, 2},
		{53, 43, 33, 23, 13, 3},
		{54, 44, 34, 24, 14, 4},
		{55, 45, 35, 25, 15, 5},
	}
	actual = RotateMatrixEfficient(matrix)

	for i, row := range matrix {
		for j := range row {
			if actual[i][j] != expected[i][j] {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		}
	}
}
