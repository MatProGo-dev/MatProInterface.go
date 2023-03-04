package optim_test

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"testing"
)

/*
vars_test.go
Description:
	Testing functions relevant to the Var() object. (Scalar Variable)
*/

/*
TestVar_Plus1
Description:

	Tests the approach of performing addition of a var with a constant.
*/
func TestVar_Plus1(t *testing.T) {
	// Constants
	m := optim.NewModel()
	x := m.AddVariable()

	k1 := optim.K(1.0)

	// Algorithm
	tempSum, err := x.Plus(k1)
	if err != nil {
		t.Errorf("There was an issue computing sum: %v", err)
	}

	sumAsSLE, ok1 := tempSum.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("The sum is expected to be a ScalarLinearExpression but it was not! (Type = %T)", tempSum)
	}

	if sumAsSLE.X.Len() != 1 {
		t.Errorf("Expected sum to have scalar variable, but found %v items.", sumAsSLE.X.Len())
	}

	if sumAsSLE.X.AtVec(0).(optim.Variable).ID != x.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(0).(optim.Variable).ID,
			x.ID,
		)
	}

	if sumAsSLE.C != float64(k1) {
		t.Errorf("Expected sum's constant to be %v; received %v.", k1, sumAsSLE.C)
	}
}

/*
TestVar_Plus2
Description:

	Tests the approach of performing addition of a var with a var.
*/
func TestVar_Plus2(t *testing.T) {
	// Constants
	m := optim.NewModel()
	x := m.AddVariable()
	y := m.AddVariable()

	// Algorithm
	tempSum, err := x.Plus(y)
	if err != nil {
		t.Errorf("There was an issue computing sum: %v", err)
	}

	sumAsSLE, ok1 := tempSum.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("The sum is expected to be a ScalarLinearExpression but it was not! (Type = %T)", tempSum)
	}

	if sumAsSLE.X.Len() != 2 {
		t.Errorf("Expected sum to have scalar variable, but found %v items.", sumAsSLE.X.Len())
	}

	if sumAsSLE.X.AtVec(0).(optim.Variable).ID != x.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(0).(optim.Variable).ID,
			x.ID,
		)
	}

	if sumAsSLE.X.AtVec(1).(optim.Variable).ID != y.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(1).(optim.Variable).ID,
			y.ID,
		)
	}

	if sumAsSLE.C != 0.0 {
		t.Errorf("Expected sum's constant to be %v; received %v.", 0.0, sumAsSLE.C)
	}
}

/*
TestVar_Plus3
Description:

	Tests the approach of performing addition of a var with a scalar linear expression.
*/
func TestVar_Plus3(t *testing.T) {
	// Constants
	m := optim.NewModel()
	x := m.AddVariable()
	y := m.AddVariable()
	z := m.AddVariable()

	vv := optim.VarVector{[]optim.Variable{y, z}}
	L := optim.OnesVector(2)
	C := 3.0
	sle1 := optim.ScalarLinearExpr{
		X: vv,
		L: L,
		C: C,
	}

	// Algorithm
	tempSum, err := x.Plus(sle1)
	if err != nil {
		t.Errorf("There was an issue computing sum: %v", err)
	}

	sumAsSLE, ok1 := tempSum.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("The sum is expected to be a ScalarLinearExpression but it was not! (Type = %T)", tempSum)
	}

	if sumAsSLE.X.Len() != 3 {
		t.Errorf("Expected sum to have scalar variable, but found %v items.", sumAsSLE.X.Len())
	}

	if sumAsSLE.X.AtVec(0).(optim.Variable).ID != x.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(0).(optim.Variable).ID,
			x.ID,
		)
	}

	if sumAsSLE.X.AtVec(1).(optim.Variable).ID != y.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(1).(optim.Variable).ID,
			y.ID,
		)
	}

	if sumAsSLE.X.AtVec(2).(optim.Variable).ID != z.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(2).(optim.Variable).ID,
			z.ID,
		)
	}

	// Checking linear terms
	if sumAsSLE.L.AtVec(0) != 1.0 {
		t.Errorf(
			"Expected first element of L (%v) to have value 1.0.",
			sumAsSLE.L.AtVec(2),
		)
	}

	if sumAsSLE.L.AtVec(1) != sle1.L.AtVec(0) {
		t.Errorf(
			"Expected second element of L (%v) to have value %v.",
			sumAsSLE.L.AtVec(1),
			sle1.L.AtVec(0),
		)
	}

	if sumAsSLE.L.AtVec(2) != sle1.L.AtVec(1) {
		t.Errorf(
			"Expected first element of L (%v) to have value %v.",
			sumAsSLE.L.AtVec(2),
			sle1.L.AtVec(1),
		)
	}

	// Checking Constant terms
	if sumAsSLE.C != sle1.C {
		t.Errorf("Expected sum's constant to be %v; received %v.", sle1.C, sumAsSLE.C)
	}
}

/*
TestVar_Plus4
Description:

	Tests the approach of performing addition of a var with a scalar quadratic expression.
*/
func TestVar_Plus4(t *testing.T) {
	// Constants
	m := optim.NewModel()
	x := m.AddVariable()
	y := m.AddVariable()
	z := m.AddVariable()

	vv := optim.VarVector{[]optim.Variable{y, z}}
	Q := optim.Identity(2)
	L := optim.OnesVector(2)
	C := 3.0
	qe1 := optim.ScalarQuadraticExpression{
		Q: Q,
		X: vv,
		L: L,
		C: C,
	}

	// Algorithm
	tempSum, err := x.Plus(qe1)
	if err != nil {
		t.Errorf("There was an issue computing sum: %v", err)
	}

	sumAsSLE, ok1 := tempSum.(optim.ScalarQuadraticExpression)
	if !ok1 {
		t.Errorf("The sum is expected to be a ScalarLinearExpression but it was not! (Type = %T)", tempSum)
	}

	if sumAsSLE.X.Len() != 3 {
		t.Errorf("Expected sum to have scalar variable, but found %v items.", sumAsSLE.X.Len())
	}

	if sumAsSLE.X.AtVec(0).(optim.Variable).ID != x.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(0).(optim.Variable).ID,
			x.ID,
		)
	}

	if sumAsSLE.X.AtVec(1).(optim.Variable).ID != y.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(1).(optim.Variable).ID,
			y.ID,
		)
	}

	if sumAsSLE.X.AtVec(2).(optim.Variable).ID != z.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(2).(optim.Variable).ID,
			z.ID,
		)
	}

	// Checking linear terms
	if sumAsSLE.L.AtVec(0) != 1.0 {
		t.Errorf(
			"Expected first element of L (%v) to have value 1.0.",
			sumAsSLE.L.AtVec(2),
		)
	}

	if sumAsSLE.L.AtVec(1) != qe1.L.AtVec(0) {
		t.Errorf(
			"Expected second element of L (%v) to have value %v.",
			sumAsSLE.L.AtVec(1),
			qe1.L.AtVec(0),
		)
	}

	if sumAsSLE.L.AtVec(2) != qe1.L.AtVec(1) {
		t.Errorf(
			"Expected first element of L (%v) to have value %v.",
			sumAsSLE.L.AtVec(2),
			qe1.L.AtVec(1),
		)
	}

	// Checking Constant terms
	if sumAsSLE.C != qe1.C {
		t.Errorf("Expected sum's constant to be %v; received %v.", qe1.C, sumAsSLE.C)
	}
}
