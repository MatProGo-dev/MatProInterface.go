package optim_test

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"testing"
)

/*
constraint_test.go
Description:
	Tests for all functions and objects defined in the constraint.go file.
*/

/*
TestConstraint_IsConstraint1
Description:

	This test verifies if a scalar constraint is properly detected by IsConstraint.
*/
func TestConstraint_IsConstraint1(t *testing.T) {
	// Constants
	m := optim.NewModel("IsConstraint1")

	// Create a scalar constraint.

	lhs0 := optim.One
	x := m.AddBinaryVariable()

	scalarConstr0, err := optim.Eq(lhs0, x)
	if err != nil {
		t.Errorf("An error occurred constructing the equality constraint: %v", err)
	}

	if !optim.IsConstraint(scalarConstr0) {
		t.Errorf("The scalar constraint is not implementing a Constraint() interface!")
	}
}

/*
TestConstraint_IsConstraint2
Description:

	This test verifies if a vector constraint is properly detected by IsConstraint.
*/
func TestConstraint_IsConstraint2(t *testing.T) {
	// Constants
	m := optim.NewModel("IsConstraint2")

	// Create a scalar constraint.

	lhs0 := optim.OnesVector(4)
	x := m.AddVariableClassic(0, 3.0, optim.Continuous)
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, x, x, x},
	}

	vectorConstr0, err := optim.Eq(lhs0, vv1)
	if err != nil {
		t.Errorf("An error occurred constructing the equality constraint: %v", err)
	}

	if !optim.IsConstraint(vectorConstr0) {
		t.Errorf("The scalar constraint is not implementing a Constraint() interface!")
	}
}

/*
TestConstraint_IsConstraint3
Description:

	This verifies that a float is not a constraint.
*/
func TestConstraint_IsConstraint3(t *testing.T) {
	// Constants
	f1 := 7.5

	// Algorithm
	if optim.IsConstraint(f1) {
		t.Errorf("The float was not properly detected as NOT BEING a constant.")
	}
}
