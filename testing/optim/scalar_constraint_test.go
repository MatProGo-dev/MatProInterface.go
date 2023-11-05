package optim

/*
scalar_constraint_test.go
Description:
	Creates the scalar constraint object.
*/

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestScalarConstraint_ScalarConstraint1(t *testing.T) {
	// Constants
	m := optim.NewModel("scalar-constraint-test1")
	v1 := m.AddVariable()
	k1 := optim.K(2.8)

	// Algorithm
	sc1 := optim.ScalarConstraint{
		LeftHandSide:  v1,
		RightHandSide: k1,
		Sense:         optim.SenseEqual,
	}

	if sc1.Sense != optim.SenseEqual {
		t.Errorf(
			"Scalar Constraint does not contain equality sense; received %v",
			sc1.Sense,
		)
	}
}

/*
TestScalarConstraint_IsLinear1
Description:

	Detects whether a simple inequality between
	a variable and a constant is a linear constraint.
*/
func TestScalarConstraint_IsLinear1(t *testing.T) {
	// Constants
	m := optim.NewModel("scalar-constraint-test1")
	v1 := m.AddVariable()
	k1 := optim.K(2.8)

	// Algorithm
	sc1 := optim.ScalarConstraint{
		LeftHandSide:  v1,
		RightHandSide: k1,
		Sense:         optim.SenseEqual,
	}

	tf, err := sc1.IsLinear()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !tf {
		t.Errorf("sc1 is linear, but function claims it is not!")
	}

}

/*
TestScalarConstraint_IsLinear1
Description:

	Detects whether a simple inequality between
	a variable and a constant is a linear constraint.
*/
func TestScalarConstraint_IsLinear2(t *testing.T) {
	// Constants
	m := optim.NewModel("scalar-constraint-test1")
	v1 := m.AddVariable()
	sqe2 := optim.ScalarQuadraticExpression{
		L: optim.OnesVector(1),
		Q: *mat.NewDense(1, 1, []float64{3.14}),
		X: optim.VarVector{Elements: []optim.Variable{v1}},
	}

	k1 := optim.K(2.8)

	// Algorithm
	sc1 := optim.ScalarConstraint{
		LeftHandSide:  sqe2,
		RightHandSide: k1,
		Sense:         optim.SenseEqual,
	}

	tf, err := sc1.IsLinear()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if tf {
		t.Errorf("sc1 is not linear, but function claims it is!")
	}

}
