package optim

/*
scalar_constraint_test.go
Description:
	Creates the scalar constraint object.
*/

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
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
