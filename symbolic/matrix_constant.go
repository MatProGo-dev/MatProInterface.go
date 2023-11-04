package symbolic

import (
	"fmt"
	"github.com/MatProGo-dev/MatProInterface.go/symbolic/matrix"
	"gonum.org/v1/gonum/mat"
)

/*
MatrixConstant
Description:

	Represents a constant matrix.
*/
type MatrixConstant mat.Dense

func (mc MatrixConstant) NumVars() int {
	return 0
}

/*
IDs
Description:

	Returns the IDs of any variables in the expression.
*/
func (mc MatrixConstant) IDs() []uint64 {
	return []uint64{}
}

func (mc MatrixConstant) Dims() []int {
	dense := mat.Dense(mc)
	nR, nC := dense.Dims()
	return []int{nR, nC}
}

/*
Multiply
Description:
*/
func (mc MatrixConstant) Multiply(rightIn interface{}, errors ...error) (Expression, error) {
	return mc, fmt.Errorf("not implemented yet...")
}

/*
Plus
Description:

	Sums this matrix with something else.
*/
func (mc MatrixConstant) Plus(rightIn interface{}, errors ...error) (Expression, error) {
	return mc, fmt.Errorf("not implemented yet...")
}

func (mc MatrixConstant) Transpose() Expression {
	dims := mc.Dims()
	transposed := matrix.Zeros(dims[1], dims[0])

	mcAsD := mat.Dense(mc)

	for rowIndex := 0; rowIndex < dims[1]; rowIndex++ {
		for colIndex := 0; colIndex < dims[0]; colIndex++ {
			transposed.Set(
				rowIndex, colIndex,
				(&mcAsD).At(colIndex, rowIndex),
			)
		}
	}
	return MatrixConstant(transposed)
}
