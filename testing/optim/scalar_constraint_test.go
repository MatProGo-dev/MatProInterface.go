package optim

/*
scalar_constraint_test.go
Description:
	Creates the scalar constraint object.
*/

import (
	"testing"

	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"gonum.org/v1/gonum/mat"
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
TestScalarConstraint_IsLinear2
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

/*
TestScalarConstraint_IsLinear3
Description:

	Detects whether a simple inequality between
	a variable and a constant is a quadratic expression.
	The function should identify that this is NOT a linear expression.
*/
func TestScalarConstraint_IsLinear3(t *testing.T) {
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
		LeftHandSide:  k1,
		RightHandSide: sqe2,
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

/*
TestScalarConstraint_Simplify1
Description:

	Attempts to simplify the constraint between
	a scalar linear epression and a scalar linear expression.
*/
func TestScalarConstraint_Simplify1(t *testing.T) {
	// Constants
	m := optim.NewModel("scalar-constraint-test1")
	vv1 := m.AddVariableVector(3)
	sle2 := optim.ScalarLinearExpr{
		L: optim.OnesVector(vv1.Len()),
		X: vv1,
		C: 2.0,
	}
	sle3 := optim.ScalarLinearExpr{
		L: *mat.NewVecDense(vv1.Len(), []float64{1.0, 2.0, 3.0}),
		X: vv1,
		C: 1.0,
	}

	// Create sles
	sc1 := optim.ScalarConstraint{
		LeftHandSide:  sle2,
		RightHandSide: sle3,
		Sense:         optim.SenseEqual,
	}

	// Attempt to simplify
	sc2, err := sc1.Simplify()
	if err != nil {
		t.Errorf("unexpected error during simplify(): %v", err)
	}

	if float64(sc2.RightHandSide.(optim.K)) != 1.0 {
		t.Errorf("Remainder on LHS was not contained properly")
	}
}

/*
TestScalarConstraint_Simplify2
Description:

	Attempts to simplify the constraint between
	a scalar linear epression and a variable.
	The resulting constraint should have a zero remainder on the right hand side
	and the left hand side should contain a variable vector with 1 more element than it started with.
*/
func TestScalarConstraint_Simplify2(t *testing.T) {
	// Constants
	m := optim.NewModel("scalar-constraint-test2")
	vv1 := m.AddVariableVector(3)
	sle2 := optim.ScalarLinearExpr{
		L: optim.OnesVector(vv1.Len()),
		X: vv1,
		C: 2.0,
	}
	v3 := m.AddVariable()

	// Create sles
	sc1 := optim.ScalarConstraint{
		LeftHandSide:  sle2,
		RightHandSide: v3,
		Sense:         optim.SenseEqual,
	}

	// Attempt to simplify
	sc2, err := sc1.Simplify()
	if err != nil {
		t.Errorf("unexpected error during simplify(): %v", err)
	}

	if float64(sc2.RightHandSide.(optim.K)) != 0.0 {
		t.Errorf("Remainder on LHS was not contained properly")
	}

	if sc2.LeftHandSide.(optim.ScalarLinearExpr).X.Len() != 4 {
		t.Errorf("Variable vector did not increase in size")
	}
}

/*
TestScalarConstraint_Simplify3
Description:

	Attempts to simplify the constraint between
	a scalar linear epression and a scalar quadratic expression.
*/
func TestScalarConstraint_Simplify3(t *testing.T) {
	// Constants
	m := optim.NewModel("scalar-constraint-test3")
	vv1 := m.AddVariableVector(3)
	sle2 := optim.ScalarLinearExpr{
		L: optim.OnesVector(vv1.Len()),
		X: vv1,
		C: 2.0,
	}
	sqe3 := optim.ScalarQuadraticExpression{
		L: optim.OnesVector(3),
		Q: *mat.NewDense(
			3, 3,
			[]float64{
				3.14, 0.0, 0.0,
				0.0, 3.14, 0.0,
				0.0, 0.0, 3.14,
			},
		),
		X: vv1,
		C: 1.5,
	}

	// Create sles
	sc1 := optim.ScalarConstraint{
		LeftHandSide:  sle2,
		RightHandSide: sqe3,
		Sense:         optim.SenseLessThanEqual,
	}

	// Attempt to simplify
	sc2, err := sc1.Simplify()
	if err != nil {
		t.Errorf("unexpected error during simplify(): %v", err)
	}

	if float64(sc2.RightHandSide.(optim.K)) != sqe3.C {
		t.Errorf(
			"Remainder on RHS was not contained properly; expected %v, received %v",
			sqe3.C,
			sc2.RightHandSide.(optim.K),
		)
	}
}
