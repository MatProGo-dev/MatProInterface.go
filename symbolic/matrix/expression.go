package matrix

import (
	"gonum.org/v1/gonum/mat"
)

//================
// Type Definition
//================

type Expression interface {
	Dims() []int // Computes the dimensions of the input matrix

	// IDs
	// Returns the ids of any variables in the matrix expression
	IDs() []uint64

	NumVars() int // Returns the number of variables in the matrix expression
}

/*
IsMatrixExpression
Description:

	Returns true if and only if the input object is a matrix expression.
*/
func IsMatrixExpression(e interface{}) bool {
	// Check each type
	switch e.(type) {
	case mat.Dense:
		return true
	case Constant:
		return true
	default:
		return false
	}
}

/*
ToMatrixExpression
Description:

	Converts the input object into a valid type that implements the
	MatrixExpression interface.
*/
func ToMatrixExpression(e interface{}) (Expression, error) {
	// Input Processing
	if !IsMatrixExpression(e) {
		return Constant(Zeros(1, 1)), TypeError{e}
	}

	// Convert
	switch candidate := e.(type) {
	case mat.Dense:
		return Constant(candidate), nil
	case Constant:
		return candidate, nil
	default:
		return Constant(Zeros(1, 1)), TypeError{e}
	}
}
