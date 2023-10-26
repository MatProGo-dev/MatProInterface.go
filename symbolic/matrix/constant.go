package matrix

import "gonum.org/v1/gonum/mat"

/*
MatrixConstant
Description:

	Represents a constant matrix.
*/
type Constant mat.Dense

func (c Constant) Dims() []uint {
	dense := mat.Dense(c)
	nR, nC := dense.Dims()
	return []uint{uint(nR), uint(nC)}
}
