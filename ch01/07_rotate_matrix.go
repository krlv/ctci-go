package ch01

// RotateMatrix rotates given matrix m to 90 degrees
func RotateMatrix(m [][]int) [][]int {
	n := len(m)

	mx := make([][]int, n)
	for i := 0; i < n; i++ {
		mx[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mx[j][n-1-i] = m[i][j]
		}
	}

	return mx
}
