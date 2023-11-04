package matrix

import "gonum.org/v1/gonum/mat"

/*
Zeros
Description:

	Returns a dense matrix of all zeros.
*/
func Zeros(nR, nC int) mat.Dense {
	// Create empty slice
	elts := make([]float64, nR*nC)
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			elts[rowIndex*nC+colIndex] = 0.0
		}
	}

	return *mat.NewDense(nR, nC, elts)
}

/*
Identity
Description:

	Returns a symmetric matrix that is the identity matrix.
	Note: this function assumes lengthIn is a positive number.
*/
func Identity(dim int) mat.Dense {
	// Create the empty matrix.
	zeroBase := Zeros(dim, dim)

	// Populate Diagonal
	for rowIndex := 0; rowIndex < dim; rowIndex++ {
		zeroBase.Set(rowIndex, rowIndex, 1.0)
	}

	return zeroBase
}
