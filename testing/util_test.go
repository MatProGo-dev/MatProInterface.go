package testing

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"testing"
)

func TestSumVars(t *testing.T) {
	numVars := 3
	m := optim.NewModel("TestSumVars")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()
	z := m.AddBinaryVariable()
	expr := optim.SumVars(x, y, z)

	// t.Errorf("%v", expr.(optim.ScalarLinearExpr))

	for _, coeff := range expr.Coeffs() {
		if coeff != 1 {
			t.Errorf("Coeff mismatch: %v != 1", coeff)
		}
	}

	if expr.NumVars() != numVars {
		t.Errorf("NumVars mismatch: %v != %v", expr.NumVars(), numVars)
	}

	if expr.Constant() != 0 {
		t.Errorf("Constant mismatch: %v != 0", expr.Constant())
	}
}

/*
TestUtil_Identity1
Description:

	Create identity matrix of dimension 1 (scalar?).
*/
func TestUtil_Identity1(t *testing.T) {
	// Constants
	n := 1
	// Algorithm
	identMat1 := optim.Identity(n)

	nX, nY := identMat1.Dims()
	if (nX != n) || (nY != n) {
		t.Errorf("The identity matrix created has dimension %v x %v; Expected %v x %v.",
			nX, nY,
			n, n,
		)
	}

}

/*
TestUtil_Identity2
Description:

	Create identity matrix of dimension 10.
*/
func TestUtil_Identity2(t *testing.T) {
	// Constants
	n := 10
	// Algorithm
	identMat1 := optim.Identity(n)

	nX, nY := identMat1.Dims()
	if (nX != n) || (nY != n) {
		t.Errorf("The identity matrix created has dimension %v x %v; Expected %v x %v.",
			nX, nY,
			n, n,
		)
	}

}
