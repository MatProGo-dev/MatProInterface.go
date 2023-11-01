package matrix

import "gonum.org/v1/gonum/mat"

/*
MatrixConstant
Description:

	Represents a constant matrix.
*/
type Constant mat.Dense

func (c Constant) NumVars() int {
	return 0
}

/*
IDs
Description:

	Returns the IDs of any variables in the expression.
*/
func (c Constant) IDs() []uint64 {
	return []uint64{}
}

func (c Constant) Dims() []int {
	dense := mat.Dense(c)
	nR, nC := dense.Dims()
	return []int{nR, nC}
}
