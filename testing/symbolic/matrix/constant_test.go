package matrix_test

import (
	"github.com/MatProGo-dev/MatProInterface.go/symbolic"
	"github.com/MatProGo-dev/MatProInterface.go/symbolic/matrix"
	"testing"
)

/*
constant_test.go
Description:

*/

/*
TestConstant_NumVars1
Description:

	Tests that NumVars() returns 0 for a matrix.Constant
*/
func TestConstant_NumVars1(t *testing.T) {
	// Constants
	mat1 := symbolic.MatrixConstant(matrix.Zeros(2, 3))

	// Algorithm
	if mat1.NumVars() != 0 {
		t.Errorf(
			"mat1.NumVars() = %v =/= 0",
			mat1.NumVars(),
		)
	}
}

/*
TestConstant_IDs1
Description:

	Tests that IDs() returns an empty slice of []uint64{}
*/
func TestConstant_IDs1(t *testing.T) {
	// Constants
	mat1 := symbolic.MatrixConstant(matrix.Zeros(2, 3))

	// Algorithm
	ids1 := mat1.IDs()
	if len(ids1) != 0 {
		t.Errorf(
			"len(ids1) = %v =/= 0",
			len(ids1),
		)
	}
}
