package matrix_test

import (
	"github.com/MatProGo-dev/MatProInterface.go/symbolic/matrix"
	"testing"
)

/*
TestUtils_Zeros1
Description:
	Tests that the zeros util function works.
*/

func TestUtils_Zeros1(t *testing.T) {
	// Constant
	mat1 := matrix.Zeros(2, 3)
	mat2 := matrix.Constant(mat1)

	// Algorithm
	dims := mat2.Dims()
	for rowIndex := 0; rowIndex < dims[0]; rowIndex++ {
		for colIndex := 0; colIndex < dims[1]; colIndex++ {
			if mat1.At(rowIndex, colIndex) != 0.0 {
				t.Errorf(
					"mat1[%v,%v] = %v =/= 0.0",
					rowIndex, colIndex,
					mat1.At(rowIndex, colIndex),
				)
			}
		}
	}

}
