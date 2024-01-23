package ntt

func NTT4Step(xn []int, n1 int) []int {
	n2 := len(xn) / n1
	if len(xn) != n1*n2 {
		panic("n1 * n2 is not equal to len(xn)")
	}
	gnn := GNN(n1 * n2)

	// PRE TRANSPOSE
	A := transpose(rearrage(xn, n1, n2))

	// STEP1
	ap := matrix(n2, n1)
	for i := 0; i < n2; i++ {
		ap[i] = NTT(A[i])
	}

	// STEP2
	app := matrix(n2, n1)
	for i := 0; i < n2; i++ {
		for j := 0; j < n1; j++ {
			app[i][j] = multiply(ap[i][j], gnn[i*j])
		}
	}

	// STEP3
	appp := transpose(app)

	// STEP4
	apppp := matrix(n1, n2)
	for i := 0; i < n1; i++ {
		apppp[i] = NTT(appp[i])
	}

	// POST TRANSPOSE
	a := transpose(apppp)
	var res []int
	for i := range a {
		res = append(res, a[i]...)
	}

	return res
}

func matrix(rows, cols int) [][]int {
	a := make([][]int, rows)
	for i := 0; i < rows; i++ {
		a[i] = make([]int, cols)
	}
	return a
}

func rearrage(xn []int, rows, cols int) [][]int {
	if rows*cols != len(xn) {
		panic("rearrage failed")
	}
	res := make([][]int, 0, rows)
	for i := 0; i < rows; i++ {
		res = append(res, xn[i*cols:i*cols+cols])
	}
	return res
}

func transpose(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	result := make([][]int, cols)
	for i := 0; i < cols; i++ {
		result[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}

// NTT4StepOpt uses array to represent matrix
func NTT4StepOpt(xn []int, n1 int) []int {
	n2 := len(xn) / n1
	if len(xn) != n1*n2 {
		panic("n1 * n2 is not equal to len(xn)")
	}
	gnn := GNN(n1 * n2)

	// PRE TRANSPOSE
	A := transposeInSlice(xn, n2)
	// STEP1
	ap := sliceMatrix(n2, n1)
	for i := 0; i < n2; i++ {
		copy(ap[i*n1:], NTT(A[i*n1:(i+1)*n1]))
	}
	// STEP2
	app := sliceMatrix(n2, n1)
	for i := 0; i < n2; i++ {
		for j := 0; j < n1; j++ {
			app[i*n1+j] = multiply(ap[i*n1+j], gnn[i*j])
		}
	}

	// STEP3
	appp := transposeInSlice(app, n1)
	// STEP4
	apppp := sliceMatrix(n1, n2)
	for i := 0; i < n1; i++ {
		copy(apppp[i*n2:], NTT(appp[i*n2:(i+1)*n2]))
	}

	// POST TRANSPOSE
	a := transposeInSlice(apppp, n2)
	return a
}

// sliceMatrix uses array to represent matrix
func sliceMatrix(n1, n2 int) []int {
	return make([]int, n1*n2)
}

// transposeInSlice transposes a matrix represented in slice
func transposeInSlice(matrix []int, cols int) []int {
	if len(matrix)%cols != 0 {
		panic("invalid matrix")
	}
	rows := len(matrix) / cols
	matrix2 := make([]int, len(matrix))
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix2[i+j*rows] = matrix[i*cols+j]
		}
	}
	return matrix2
}
