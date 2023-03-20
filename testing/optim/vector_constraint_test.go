package optim

/*
vector_constraint_test.go
Description:
	A set of tests for the class VectorConstraint and its associated methods.
*/

import (
	"fmt"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"strings"
	"testing"
)

/*
TestVectorConstraint_Check1
Description:

	Checks to see whether or not a bad vector constraint is good or not.
	Provide a vector constraint that contains a slightly bad expression (length of one expression is different than the other).
*/
func TestVectorConstraint_Check1(t *testing.T) {
	// Constants
	dim := 4
	lhs0 := optim.OnesVector(dim - 1)
	L1 := optim.Identity(dim)

	m := optim.NewModel("vc-check1")
	x := m.AddVariableClassic(0, 3.0, optim.Continuous)
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, x, x, x},
	}

	vle := optim.VectorLinearExpr{
		L: L1,
		X: vv1,
		C: optim.ZerosVector(dim),
	}

	vc1 := optim.VectorConstraint{
		LeftHandSide:  optim.KVector(lhs0),
		RightHandSide: vle,
		Sense:         optim.SenseLessThanEqual,
	}

	// Algorithm
	if vc1.Check() == nil {
		t.Errorf("The vector constraint passed all checks, but it is a bad expression!")
	}
}

/*
TestVectorConstraint_Check2
Description:

	This function checks whether or not a good vector constraint is good or not.
	All dimensions of the vectors in the constraint should match.
*/
func TestVectorConstraint_Check2(t *testing.T) {
	// Constants
	dim := 4
	lhs0 := optim.OnesVector(dim)

	m := optim.NewModel("vc-check1")
	x := m.AddVariableClassic(0, 3.0, optim.Continuous)
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, x, x, x},
	}

	// Create constraint
	vc1 := optim.VectorConstraint{
		LeftHandSide:  optim.KVector(lhs0),
		RightHandSide: vv1,
	}

	// Algorithm
	if vc1.Check() != nil {
		t.Errorf("The vector constraint should be valid, but received error: %v", vc1.Check())
	}

}

/*
TestVectorConstraint_AtVec1
Description:

	Tests whether or not the function AtVec() throws an error properly if given a bad expression to use AtVec on.
*/
func TestVectorConstraint_AtVec1(t *testing.T) {
	// Constants
	dim := 4
	lhs0 := optim.OnesVector(dim - 1)
	L1 := optim.Identity(dim)

	m := optim.NewModel("vc-check1")
	x := m.AddVariableClassic(0, 3.0, optim.Continuous)
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, x, x, x},
	}

	vle := optim.VectorLinearExpr{
		L: L1,
		X: vv1,
		C: optim.ZerosVector(dim),
	}

	// Create Vector Constraint
	vc1 := optim.VectorConstraint{
		LeftHandSide:  optim.KVector(lhs0),
		RightHandSide: vle,
	}

	// Algorithm
	_, err := vc1.AtVec(1)
	if err == nil {
		t.Errorf("The vector constraint should be invalid and fail checks, but received none!")
	}

	if !strings.Contains(
		err.Error(), fmt.Sprintf(
			"Left hand side has dimension %v, but right hand side has dimension %v!",
			vc1.LeftHandSide.Len(),
			vc1.RightHandSide.Len(),
		)) {
		t.Errorf("Unexpected error occured: %v", err)
	}
}

/*
TestVectorConstraint_AtVec2
Description:

	Tests whether or not the function AtVec() throws an error properly if given a bad index to the AtVec function.
*/
func TestVectorConstraint_AtVec2(t *testing.T) {
	// Constants
	dim := 4
	lhs0 := optim.OnesVector(dim)
	L1 := optim.Identity(dim)

	m := optim.NewModel("vc-check1")
	x := m.AddVariableClassic(0, 3.0, optim.Continuous)
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, x, x, x},
	}

	vle := optim.VectorLinearExpr{
		L: L1,
		X: vv1,
		C: optim.ZerosVector(dim),
	}

	// Create Vector Constraint
	vc1 := optim.VectorConstraint{
		LeftHandSide:  optim.KVector(lhs0),
		RightHandSide: vle,
	}

	// Algorithm
	_, err := vc1.AtVec(dim)
	if err == nil {
		t.Errorf("The vector constraint should be invalid and fail checks, but received none!")
	}

	if !strings.Contains(
		err.Error(), fmt.Sprintf(
			"Cannot extract VectorConstraint element at %v; VectorConstraint has length %v.",
			dim, vle.Len(),
		)) {
		t.Errorf("Unexpected error occured: %v", err)
	}
}

/*
TestVectorConstraint_AtVec3
Description:

	Tests whether or not the function AtVec() doesn't throw an error when properly accessing a well-structured vector constraint.
*/
func TestVectorConstraint_AtVec3(t *testing.T) {
	// Constants
	dim := 4
	lhs0 := optim.OnesVector(dim)
	L1 := optim.Identity(dim)

	m := optim.NewModel("vc-check1")
	x := m.AddVariableClassic(0, 3.0, optim.Continuous)
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, x, x, x},
	}

	vle := optim.VectorLinearExpr{
		L: L1,
		X: vv1,
		C: optim.ZerosVector(dim),
	}

	// Create Vector Constraint
	vc1 := optim.VectorConstraint{
		LeftHandSide:  optim.KVector(lhs0),
		RightHandSide: vle,
	}

	// Algorithm
	targetIndex := 1
	sc1, err := vc1.AtVec(targetIndex)
	if err != nil {
		t.Errorf("The vector constraint should be invalid and fail checks, but received none!")
	}

	// Match all vector indices
	sc1LHSAsK, tf := sc1.LeftHandSide.(optim.K)
	if !tf {
		t.Errorf("The left hand side was not of type optim.K; actually, it's %T", sc1.LeftHandSide)
	}

	if float64(sc1LHSAsK) != lhs0.AtVec(targetIndex) {
		t.Errorf("Value of slc[%v] (%v) doesn't match lhs0[%v] (%v)!",
			targetIndex, sc1LHSAsK,
			targetIndex, lhs0.AtVec(targetIndex),
		)
	}

	sc1RHSAsSLE, tf := sc1.RightHandSide.(optim.ScalarLinearExpr)
	if !tf {
		t.Errorf("The right hand side was not of type optim.ScalarLinearExpr; actually, it's %T", sc1.RightHandSide)
	}

	for i, elt := range sc1RHSAsSLE.X.Elements {
		if elt.ID != vv1.Elements[i].ID {
			t.Errorf(
				"The %vth vector in rhs (%v) doesn't match vv1[%v] (%v)!",
				i, elt,
				i, vv1.Elements[i].ID,
			)
		}
	}

}
