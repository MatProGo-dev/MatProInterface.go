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
TestModel_AddRealVariable1
Description:
*/
func TestModel_AddRealVariable1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-addvariable1")

	// Algorithm
	if len(m.Variables) != 0 {
		t.Errorf(
			"The uninitialized model has %v variables; expected 0!",
			len(m.Variables),
		)
	}

	v := m.AddRealVariable()
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
	if v.Vtype != optim.Continuous {
		t.Errorf(
			"The type of the created variable was %v; not optim.Continuous",
			v.Vtype,
		)
	}
}

/*
TestModel_AddBinaryVariable1
Description:
*/
func TestModel_AddBinaryVariable1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-addbinaryvariable1")

	// Algorithm
	if len(m.Variables) != 0 {
		t.Errorf(
			"The uninitialized model has %v variables; expected 0!",
			len(m.Variables),
		)
	}

	v := m.AddBinaryVariable()
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
	if v.Lower != 0.0 {
		t.Errorf(
			"The lower bound for the binary variable was %v; not 0.0",
			v.Lower,
		)
	}
	if v.Upper != 1.0 {
		t.Errorf(
			"The upper bound for the binary variable was %v; not 1.0",
			v.Upper,
		)
	}

	if v.Vtype != optim.Binary {
		t.Errorf(
			"Unexpected variable type. Expected Binary, received %v",
			v.Vtype,
		)
	}
}

/*
TestModel_AddVariableClassic1
Description:
*/
func TestModel_AddVariableClassic1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-addvariable-classic1")

	// Algorithm
	if len(m.Variables) != 0 {
		t.Errorf(
			"The uninitialized model has %v variables; expected 0!",
			len(m.Variables),
		)
	}

	lower0 := -10.0
	upper0 := 100.0
	v := m.AddVariableClassic(lower0, upper0, optim.Continuous)
	if len(m.Variables) != 1 {
		t.Errorf(
			"After adding one variable to the model expected for slice to have one element; received %v",
			len(m.Variables),
		)
	}
	if v.Lower != lower0 {
		t.Errorf(
			"Expected lower bound to be %v; received %v",
			lower0,
			v.Lower,
		)
	}
	if v.Upper != upper0 {
		t.Errorf(
			"Expected upper bound to be %v; received %v",
			upper0,
			v.Upper,
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
TestModel_AddVariableVector1
Description:

	This test verifies that the AddVariableVector function correctly adds
	a vector of variables.
*/
func TestModel_AddVariableVector1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-addbinaryvariable1")
	nVars := 4

	// Algorithm
	if len(m.Variables) != 0 {
		t.Errorf(
			"The uninitialized model has %v variables; expected 0!",
			len(m.Variables),
		)
	}

	m.AddVariableVector(nVars)
	if len(m.Variables) != nVars {
		t.Errorf(
			"After adding one variable to the model expected for slice to have one element; received %v",
			len(m.Variables),
		)
	}
	for _, tempVar := range m.Variables {
		if tempVar.ID < 0 {
			t.Errorf(
				"The ID of the new variable was %v; expected > 0",
				tempVar.ID,
			)
		}
		if tempVar.Vtype != optim.Continuous {
			t.Errorf(
				"Unexpected variable type. Expected Continuous, received %v",
				tempVar.Vtype,
			)
		}
	}
}

/*
TestModel_AddVariableMatrix1
Description:

	This test will verify that the appropriate number of variables are created by
	AddVariableMatrix.
*/
func TestModel_AddVariableMatrix1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-addvariablematrix1")
	nVars := 4

	// Algorithm
	if len(m.Variables) != 0 {
		t.Errorf(
			"The uninitialized model has %v variables; expected 0!",
			len(m.Variables),
		)
	}

	m.AddVariableMatrix(nVars, nVars, -optim.INFINITY, optim.INFINITY, optim.Continuous)
	if len(m.Variables) != nVars*nVars {
		t.Errorf(
			"After adding one variable to the model expected for slice to have one element; received %v",
			len(m.Variables),
		)
	}
	for _, tempVar := range m.Variables {
		if tempVar.ID < 0 {
			t.Errorf(
				"The ID of the new variable was %v; expected > 0",
				tempVar.ID,
			)
		}
		if tempVar.ID > uint64(nVars*nVars) {
			t.Errorf(
				"The ID of the new variable was %v; expected < %v",
				tempVar.ID,
				nVars*nVars,
			)
		}
		if tempVar.Vtype != optim.Continuous {
			t.Errorf(
				"Unexpected variable type. Expected Continuous, received %v",
				tempVar.Vtype,
			)
		}
	}
}

/*
TestModel_AddBinaryVariableMatrix1
Description:

	This test will verify that the appropriate number of variables are created by
	AddVariableMatrix.
*/
func TestModel_AddBinaryVariableMatrix1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-addbinaryvariablematrix1")
	nVars := 4

	// Algorithm
	if len(m.Variables) != 0 {
		t.Errorf(
			"The uninitialized model has %v variables; expected 0!",
			len(m.Variables),
		)
	}

	m.AddBinaryVariableMatrix(nVars, nVars)
	if len(m.Variables) != nVars*nVars {
		t.Errorf(
			"After adding one variable to the model expected for slice to have one element; received %v",
			len(m.Variables),
		)
	}
	for _, tempVar := range m.Variables {
		if tempVar.ID < 0 {
			t.Errorf(
				"The ID of the new variable was %v; expected > 0",
				tempVar.ID,
			)
		}
		if tempVar.ID > uint64(nVars*nVars) {
			t.Errorf(
				"The ID of the new variable was %v; expected < %v",
				tempVar.ID,
				nVars*nVars,
			)
		}
		if tempVar.Vtype != optim.Binary {
			t.Errorf(
				"Unexpected variable type. Expected Continuous, received %v",
				tempVar.Vtype,
			)
		}
	}
}

/*
TestModel_SetObjective1
Description:

	This test verifies that the AddVariableVector function and some other functions
	can be used to set the objective properly.
*/
func TestModel_SetObjective1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-setobjective1")
	nVars := 4

	// Algorithm
	if len(m.Variables) != 0 {
		t.Errorf(
			"The uninitialized model has %v variables; expected 0!",
			len(m.Variables),
		)
	}

	vv := m.AddVariableVector(nVars)
	sle1 := optim.ScalarLinearExpr{
		X: vv,
		L: optim.OnesVector(nVars),
		C: 3.14,
	}

	m.SetObjective(sle1, optim.SenseMinimize)

	if m.Obj.Sense != optim.SenseMinimize {
		t.Errorf("The sense of the objective should be minimize; received %v", m.Obj.Sense)
	}

	if _, ok := m.Obj.ScalarExpression.(optim.ScalarLinearExpr); !ok {
		t.Errorf(
			"Expected for objective to be a ScalarLinearExpression; received %T",
			m.Obj.ScalarExpression,
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
	err = m.AddConstraint(slc1)
	if err != nil {
		t.Errorf("There was an issue adding the constraint to the model: %v", err)
	}

	if len(m.Constraints) != 1 {
		t.Errorf("Expected for the updated model to contain 1 constraint; received %v", len(m.Constraints))
	}

}

/*
TestModel_AddConstraint2
Description:

	Tests that a simple constraint (VectorConstraint) can be given to the model.
*/
func TestModel_AddConstraint2(t *testing.T) {
	// Constants
	m := optim.NewModel("AddConstraint1")

	// Create Constraint
	n := 3
	vv1 := m.AddVariableVector(n)
	kv1 := optim.OnesVector(n)
	slc1, err := vv1.LessEq(kv1)
	if err != nil {
		t.Errorf("There was an issue creating the desired scalar constraint: %v", err)
	}

	// Add Constraint to Model
	err = m.AddConstraint(slc1)
	if err != nil {
		t.Errorf("There was an issue adding the constraint to the model: %v", err)
	}

	if len(m.Constraints) != 1 {
		t.Errorf("Expected for the updated model to contain 1 constraint; received %v", len(m.Constraints))
	}

}

/*
TestModel_AddConstraint3
Description:

	Tests that a simple constraint (VectorConstraint) can be given to the model.
*/
func TestModel_AddConstraint3(t *testing.T) {
	// Constants
	m := optim.NewModel("AddConstraint1")

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
	err = m.AddConstraint(slc1)
	if err != nil {
		t.Errorf("There was an issue adding the constraint to the model: %v", err)
	}

	if len(m.Constraints) != 1 {
		t.Errorf("Expected for the updated model to contain 1 constraint; received %v", len(m.Constraints))
	}

}
