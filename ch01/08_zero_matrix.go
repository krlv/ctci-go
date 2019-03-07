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

// ZeroMatrixSpaceEfficient space efficient zero matrix
func ZeroMatrixSpaceEfficient(m [][]int) {
	rows := len(m)
	if rows < 1 {
		return
	}

	cols := len(m[0])
	if cols < 1 {
		return
	}

	zrow, zcol := false, false
	for j := 0; j < cols; j++ {
		if m[0][j] == 0 {
			zcol = true
			break
		}
	}

	for i := 0; i < rows; i++ {
		if m[i][0] == 0 {
			zrow = true
			break
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if m[i][j] == 0 {
				m[0][j] = 0
				m[i][0] = 0
			}
		}
	}

	for j := 1; j < cols; j++ {
		if m[0][j] != 0 {
			continue
		}

		for i := 0; i < rows; i++ {
			m[i][j] = 0
		}
	}

	for i := 1; i < rows; i++ {
		if m[i][0] != 0 {
			continue
		}

		for j := 0; j < cols; j++ {
			m[i][j] = 0
		}
	}

	if zcol {
		for j := 0; j < cols; j++ {
			m[0][j] = 0
		}
	}

	if zrow {
		for i := 0; i < rows; i++ {
			m[i][0] = 0
		}
	}
}
