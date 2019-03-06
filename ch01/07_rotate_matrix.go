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

// RotateMatrixEfficient rotates given matrix m to 90 degrees
func RotateMatrixEfficient(m [][]int) [][]int {
	n := len(m)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			m[j][n-1-i], m[i][j], m[n-1-j][i], m[n-1-i][n-1-j] = m[i][j], m[n-1-j][i], m[n-1-i][n-1-j], m[j][n-1-i]
		}
	}

	return m
}
