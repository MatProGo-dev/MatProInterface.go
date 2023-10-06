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

/*
TestScalarExpression_IsScalarExpression1
Description:

	Tests whether or not the IsScalarExpression function works on ScalarQuadraticExpression.
*/
func TestScalarExpression_IsScalarExpression1(t *testing.T) {
	// Constants
	m := optim.NewModel("testse_IsScalarExpression1")
	N := 10
	x := m.AddVariableVector(N)

	qe1 := optim.ScalarQuadraticExpression{
		Q: optim.Identity(N),
		X: x,
		L: optim.OnesVector(N),
		C: 3.14,
	}

	// Check
	if !optim.IsScalarExpression(qe1) {
		t.Errorf("ScalarQuadraticExpression is a scalar expression, but the toolbox does not think so!")
	}
}

/*
TestScalarExpression_ToScalarExpression1
Description:

	Tests how the conversion function ToScalarExpression() works on a
	scalar quadratic expression.
*/
func TestScalarExpression_ToScalarExpression1(t *testing.T) {
	// Constants
	m := optim.NewModel("testse_ToScalarExpression1")
	N := 10
	x := m.AddVariableVector(N)

	qe1 := optim.ScalarQuadraticExpression{
		Q: optim.Identity(N),
		X: x,
		L: optim.OnesVector(N),
		C: 3.14,
	}

	// Algorithm
	sqe2, err := optim.ToScalarExpression(qe1)
	if err != nil {
		t.Errorf("ToScalarExpression created an error: %v", err)
	}

	_, ok := sqe2.(optim.ScalarQuadraticExpression)
	if !ok {
		t.Errorf(
			"sqe2 should be a ScalarQuadraticExpression, but it was %T",
			sqe2,
		)
	}
}
