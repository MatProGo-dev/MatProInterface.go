package optim

/*
model_test.go
Description:
	This script tests the model object.
*/

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"testing"
)

/*
TestModel_NewModel1
Description:

	Tests the new model was initialized in the proper way.
*/
func TestModel_NewModel1(t *testing.T) {
	// Constants
	m := optim.NewModel("test")

	// Algorithm
	if m.ShowLog {
		t.Errorf("ShowLog should be initialized as false.")
	}

	if m.Name != "test" {
		t.Errorf("Expected model's name to be %v; received %v", "test", m.Name)
	}
}

/*
TestModel_AddConstr1
Description:

	Tests that a simple constraint (scalarlinearconstraint) can be given to the model.
*/
func TestModel_AddConstr1(t *testing.T) {
	// Constants
	m := optim.NewModel("AddConstr1")

	// Create Constraint
	n := 3
	vv1 := m.AddVariableVector(n)
	sle1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(n),
		X: vv1,
		C: 0.0,
	}
	slc1, err := sle1.LessEq(optim.K(0.0))
	if err != nil {
		t.Errorf("There was an issue creating the desired scalar constraint: %v", err)
	}

	// Add Constraint to Model
	err = m.AddConstr(slc1)
	if err != nil {
		t.Errorf("There was an issue adding the constraint to the model: %v", err)
	}

	if len(m.Constraints) != 1 {
		t.Errorf("Expected for the updated model to contain 1 constraint; received %v", len(m.Constraints))
	}

}
