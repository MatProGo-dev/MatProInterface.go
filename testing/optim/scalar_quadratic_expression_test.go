package optim

/*
scalar_quadratic_expression_test.go
Description:
	Defines tests on the scalar quadratic expression tests.
*/

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"testing"
)

func TestScalarQudraticExpression_NewQuadraticExpr_qb0_1(t *testing.T) {
	// Constants
	m := optim.NewModel("NewQuadraticExpr_qb0-test1")

	n := 4
	Q0 := optim.Identity(n)

	X0 := m.AddVariableVector(n)

	// Algorithm
	sqe, err := optim.NewQuadraticExpr_qb0(Q0, X0)
	if err != nil {
		t.Errorf("There was an issue creating the quadratic expression: %v", err)
	}

	for rowIndex := 0; rowIndex < n; rowIndex++ {
		for colIndex := 0; colIndex < n; colIndex++ {
			// Compare row, col indicies
			if Q0.At(rowIndex, colIndex) != sqe.Q.At(rowIndex, colIndex) {
				t.Errorf(
					"Expected sqe.Q.At(%v,%v) to be %v; received %v",
					rowIndex, colIndex,
					Q0.At(rowIndex, colIndex),
					sqe.Q.At(rowIndex, colIndex),
				)
			}
		}
	}

	// Compare linear vector
	for rowIndex := 0; rowIndex < n; rowIndex++ {
		if sqe.L.AtVec(rowIndex) != 0.0 {
			t.Errorf(
				"Expected sqe.L.AtVec(%v) to be 0.0; received %v",
				rowIndex,
				sqe.L.AtVec(rowIndex),
			)
		}
	}

}
