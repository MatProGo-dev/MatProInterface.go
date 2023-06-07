package optim

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"testing"
)

/*
scalar_expression_test.go
Description:
	Tests whether or not the functions designed for the ScalarExpression interace
	works well.
*/

/*
TestScalarExpression_NewScalarExpression1
Description:

	Tests how well the algorithm detects a small scalar expression created with
	NewScalarExpression.
*/
func TestScalarExpression_NewScalarExpression1(t *testing.T) {
	// Constants
	se := optim.NewScalarExpression(2.1)

	seAsSLE, ok1 := se.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("se was not optim.ScalarLinearExpr, but instead %T", se)
	}

	if seAsSLE.C != 2.1 {
		t.Errorf(
			"Expected offset to be 2.1; received %v",
			seAsSLE.C,
		)
	}
}
