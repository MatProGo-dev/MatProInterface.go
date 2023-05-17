package optim

/*
model_test.go
Description:
	This script tests the model object.
*/

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"testing"
	"time"
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
TestModel_SetTimeLimit1
Description:

	Tests the ability of this function to set the time limit for the model solution.
*/
func TestModel_SetTimeLimit1(t *testing.T) {
	// Constants
	t1 := time.Now()
	dur1 := time.Since(t1)

	m := optim.NewModel("test-settimelimit1")

	// Algorithm
	m.SetTimeLimit(dur1)

	if m.TimeLimit != dur1 {
		t.Errorf(
			"Expected time limit to be %v; received %v",
			m.TimeLimit,
			dur1,
		)
	}

}

/*
TestModel_AddVariable1
Description:
*/
func TestModel_AddVariable1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-addvariable1")

	// Algorithm
	if len(m.Variables) != 0 {
		t.Errorf(
			"The uninitialized model has %v variables; expected 0!",
			len(m.Variables),
		)
	}

	v := m.AddVariable()
	if len(m.Variables) != 1 {
		t.Errorf(
			"After adding one variable to the model expected for slice to have one element; received %v",
			len(m.Variables),
		)
	}
	if v.ID != 0 {
		t.Errorf(
			"The ID of the new variable was %v; expected %v",
			v.ID, 0,
		)
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

/*
TestModel_AddConstr2
Description:

	Tests that a simple constraint (VectorConstraint) can be given to the model.
*/
func TestModel_AddConstr2(t *testing.T) {
	// Constants
	m := optim.NewModel("AddConstr1")

	// Create Constraint
	n := 3
	vv1 := m.AddVariableVector(n)
	kv1 := optim.OnesVector(n)
	slc1, err := vv1.LessEq(kv1)
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

/*
TestModel_AddConstr3
Description:

	Tests that a simple constraint (VectorConstraint) can be given to the model.
*/
func TestModel_AddConstr3(t *testing.T) {
	// Constants
	m := optim.NewModel("AddConstr1")

	// Create Constraint
	n := 3
	vv1 := m.AddVariableVector(n)
	kv1 := optim.OnesVector(n)
	L1 := optim.Identity(n)
	L1.Scale(2.0, &L1)
	vle1 := optim.VectorLinearExpr{
		L: L1,
		X: vv1,
		C: kv1,
	}

	slc1, err := vle1.LessEq(optim.ZerosVector(n))
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
