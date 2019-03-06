package ch01

// ZeroMatrix will turn entire row and column into zeros if an element is zero
func ZeroMatrix(m [][]int) {
	rows := len(m)
	if rows < 1 {
		return
	}

	cols := len(m[0])
	if cols < 1 {
		return
	}

	var zrows, zcols []int

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if m[i][j] == 0 {
				zrows = append(zrows, i)
				zcols = append(zcols, j)
			}
		}
	}

	for _, row := range zrows {
		for j := 0; j < cols; j++ {
			m[row][j] = 0
		}
	}

	for _, col := range zcols {
		for i := 0; i < rows; i++ {
			m[i][col] = 0
		}
	}
}
