package optim_test

import (
	"fmt"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"gonum.org/v1/gonum/mat"
	"strings"
	"testing"
)

func TestLinearExpr_CoeffsAndConstant1(t *testing.T) {
	m := optim.NewModel("CoeffsAndConstant")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// 2 * x + 4 * y - 5
	coeffs := []float64{2, 4}
	constant := -5.0
	expr1, err := x.Multiply(coeffs[0])
	if err != nil {
		t.Errorf("There was an error computing the first multiplication: %v", err)
	}
	expr2, err := y.Multiply(coeffs[1])
	if err != nil {
		t.Errorf("There was an error computing the second multiplication: %v", err)
	}

	expr, err := optim.Sum(expr1, expr2, optim.K(constant))
	if err != nil {
		t.Errorf("There was an issue computing the Sum of the expressions: %v", err)
	}

	exprAsSLE, _ := expr.(optim.ScalarLinearExpr)
	for i, coeff := range exprAsSLE.Coeffs() {
		if coeffs[i] != coeff {
			t.Errorf("Coeff mismatch: %v != %v", coeff, coeffs[i])
		}
	}

	if exprAsSLE.Constant() != constant {
		t.Errorf("Constant mismatch: %v != %v", exprAsSLE.Constant(), constant)
	}
}

/*
TestScalarLinearExpr_IDs1
Description:

	Tests how well the IDs() method works for the ScalarLinearExpr
*/
func TestScalarLinearExpr_IDs1(t *testing.T) {
	// Constants
	N := 4
	m := optim.NewModel("Plus1")
	vv1 := m.AddVariableVector(N)

	sle1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(N),
		X: vv1,
		C: 2.14,
	}

	// Check the IDs method
	ids0 := sle1.IDs()
	if len(ids0) != N {
		t.Errorf(
			"Expected %v IDs; received %v",
			len(ids0),
			N,
		)
	}

	var idsToCompare []uint64
	for _, variable := range vv1.Elements {
		idsToCompare = append(idsToCompare, variable.ID)
	}

	for _, elt := range ids0 {
		if foundIndex, _ := optim.FindInSlice(elt, idsToCompare); foundIndex == -1 {
			t.Errorf(
				"could not find id %v in original set of IDs.",
				elt,
			)
		}
	}

}

/*
TestScalarLinearExpr_GreaterEq1
Description:

	Makes sure that GreaterEq works for an arbitrary input.
*/
func TestScalarLinearExpr_GreaterEq1(t *testing.T) {
	// Constants
	m := optim.NewModel("SLE-GreaterEq1")
	N := 3
	vv1 := m.AddVariableVector(N)
	L1 := optim.OnesVector(N)

	sle1 := optim.ScalarLinearExpr{
		L: L1,
		X: vv1,
		C: 3.14,
	}

	// Algorithm
	constr2, err := sle1.GreaterEq(1.0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	lhs, tf := constr2.(optim.ScalarConstraint).LeftHandSide.(optim.ScalarLinearExpr)
	if !tf {
		t.Errorf(
			"the left hand side was not identified as a ScalarLinearExpr but instead %T",
			constr2.(optim.ScalarConstraint).LeftHandSide,
		)
	}

	if lhs.C != sle1.C {
		t.Errorf(
			"lhs.C = %v =/= %v = sle1.C",
			lhs.C,
			sle1.C,
		)
	}

}

/*
TestScalarLinearExpr_Eq1
Description:

	Makes sure that Eq works with a
	variable.
*/
func TestScalarLinearExpr_Eq1(t *testing.T) {
	// Constants
	m := optim.NewModel("SLE-Eq1")
	N := 3
	vv1 := m.AddVariableVector(N)
	L1 := optim.OnesVector(N)

	sle1 := optim.ScalarLinearExpr{
		L: L1,
		X: vv1,
		C: 3.14,
	}

	v2 := m.AddVariable()

	// Algorithm
	constr2, err := sle1.GreaterEq(v2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	lhs, tf := constr2.(optim.ScalarConstraint).LeftHandSide.(optim.ScalarLinearExpr)
	if !tf {
		t.Errorf(
			"the left hand side was not identified as a ScalarLinearExpr but instead %T",
			constr2.(optim.ScalarConstraint).LeftHandSide,
		)
	}

	if lhs.C != sle1.C {
		t.Errorf(
			"lhs.C = %v =/= %v = sle1.C",
			lhs.C,
			sle1.C,
		)
	}

	rhs, tf := constr2.(optim.ScalarConstraint).RightHandSide.(optim.Variable)
	if !tf {
		t.Errorf(
			"the right hand side was not identified as a variable; it is instead a %T",
			constr2.(optim.ScalarConstraint).RightHandSide,
		)
	}

	if rhs.ID != v2.ID {
		t.Errorf(
			"rhs.ID = %v =/= %v = v2.ID",
			rhs.ID, v2.ID,
		)
	}

}

/*
TestScalarLinearExpr_Comparison1
Description:

	Makes sure that Comparison throws an error when compared with a
	constant with a malformed sle!
*/
func TestScalarLinearExpr_Comparison1(t *testing.T) {
	// Constants
	m := optim.NewModel("SLE-Eq1")
	N := 3
	vv1 := m.AddVariableVector(N - 1)
	L1 := optim.OnesVector(N)

	sle1 := optim.ScalarLinearExpr{
		L: L1,
		X: vv1,
		C: 3.14,
	}

	v2 := m.AddVariable()

	// Algorithm
	_, err := sle1.GreaterEq(v2)
	if err == nil {
		t.Errorf("no error was thrown, but there should have been!")
	} else {
		if !strings.Contains(
			err.Error(),
			fmt.Sprintf(
				"the length of L (%v) does not match that of X (%v)!",
				sle1.L.Len(),
				sle1.X.Len(),
			),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}

}

/*
TestScalarLinearExpr_Plus1
Description:

	This function should test the Plus method of ScalarLinearExpr for a very nice case.
	Two SLEs with the SAME varvector and simple constants.
*/
func TestScalarLinearExpr_Plus1(t *testing.T) {
	// Constants
	L1 := optim.OnesVector(2)
	c1 := 2.0

	L2 := optim.OnesVector(2)
	L2.ScaleVec(3.0, &L2)
	c2 := 5.0

	m := optim.NewModel("Plus1")
	vv1 := m.AddVariableVector(2)

	// Create sle's
	sle1 := optim.ScalarLinearExpr{
		L: L1, C: c1, X: vv1,
	}

	sle2 := optim.ScalarLinearExpr{
		L: L2, C: c2, X: vv1,
	}

	// Algorithm
	sle3, err := sle1.Plus(sle2)
	if err != nil {
		t.Errorf("There was an issue computing the sum of sle1 and sle2: %v", err)
	}

	sle3AsSLE, ok1 := sle3.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("Expected the addition of ScalarLinearExpr with another ScalarLinearExpr to create another ScalarLinearExpr. Received %T.", sle3)
	}

	for dimIndex := 0; dimIndex < 2; dimIndex++ {
		if sle3AsSLE.L.AtVec(dimIndex) != sle1.L.AtVec(dimIndex)+sle2.L.AtVec(dimIndex) {
			t.Errorf(
				"Sum failed for L at index %v; %v != %v + %v",
				dimIndex,
				sle3AsSLE.L.AtVec(dimIndex),
				sle1.L.AtVec(dimIndex),
				sle2.L.AtVec(dimIndex),
			)
		}
	}
}

/*
TestScalarLinearExpr_Plus2
Description:

	This function should test the Plus method of ScalarLinearExpr for a very nice case.
	Two SLEs with very similar varvector objects simple constants.
*/
func TestScalarLinearExpr_Plus2(t *testing.T) {
	// Constants
	L1 := optim.OnesVector(2)
	c1 := 2.0

	L2 := optim.OnesVector(2)
	L2.ScaleVec(3.0, &L2)
	c2 := 5.0

	m := optim.NewModel("Plus2")
	vv1 := m.AddVariableVector(3)

	vv2 := optim.VarVector{
		vv1.Elements[:2],
	}
	vv3 := optim.VarVector{
		vv1.Elements[1:],
	}

	// Create sle's
	sle1 := optim.ScalarLinearExpr{
		L: L1, C: c1, X: vv2,
	}

	sle2 := optim.ScalarLinearExpr{
		L: L2, C: c2, X: vv3,
	}

	// Algorithm
	sle3, err := sle1.Plus(sle2)
	if err != nil {
		t.Errorf("There was an issue computing the product of sle1 and sle2: %v", err)
	}

	sle3AsSLE, ok1 := sle3.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("Expected the addition of ScalarLinearExpr with another ScalarLinearExpr to create another ScalarLinearExpr. Received %T.", sle3)
	}

	// Check that dimension of new expression has three d X
	if (sle3AsSLE.X.Len() != 3) || (sle1.X.Len() != 2) || (sle2.X.Len() != 2) {
		t.Errorf("The ScalarLinearExpression created by this sum should have dimension three even though the original two had dimension 2.")
	}

	for XIndex, elt := range sle3AsSLE.X.Elements {
		switch elt.ID {
		case vv2.Elements[0].ID:
			if sle3AsSLE.L.AtVec(XIndex) != L1.AtVec(0) {
				t.Errorf(
					"The variable with ID %v is expected to have coefficient %v; received %v",
					elt.ID,
					L1.AtVec(0),
					sle3AsSLE.L.AtVec(XIndex),
				)
			}
		case vv3.Elements[0].ID:
			if sle3AsSLE.L.AtVec(XIndex) != L1.AtVec(0)+L2.AtVec(0) {
				t.Errorf(
					"The variable with ID %v is expected to have coefficient %v; received %v",
					elt.ID,
					L1.AtVec(0)+L2.AtVec(0),
					sle3AsSLE.L.AtVec(XIndex),
				)
			}
		case vv3.Elements[1].ID:
			if sle3AsSLE.L.AtVec(XIndex) != L2.AtVec(0) {
				t.Errorf(
					"The variable with ID %v is expected to have coefficient %v; received %v",
					elt.ID,
					L2.AtVec(0),
					sle3AsSLE.L.AtVec(XIndex),
				)
			}
		default:
			t.Errorf("Unexpected ID received! %v", elt.ID)
		}

	}
}

/*
TestScalarLinearExpr_Plus3
Description:

	This function should test the Plus method of ScalarLinearExpr for the case of (SLE + K).
*/
func TestScalarLinearExpr_Plus3(t *testing.T) {
	// Constants
	L1 := optim.OnesVector(2)
	c1 := 2.0

	K1 := optim.K(5)

	m := optim.NewModel("Plus3")
	vv1 := m.AddVariableVector(2)

	// Create sle's
	sle1 := optim.ScalarLinearExpr{
		L: L1, C: c1, X: vv1,
	}

	// Algorithm
	sle3, err := sle1.Plus(K1)
	if err != nil {
		t.Errorf("There was an issue computing the product of sle1 and sle2: %v", err)
	}

	sle3AsSLE, ok1 := sle3.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("Expected the addition of ScalarLinearExpr with another ScalarLinearExpr to create another ScalarLinearExpr. Received %T.", sle3)
	}

	if sle3AsSLE.C != sle1.C+float64(K1) {
		t.Errorf(
			"Expected for the new SLE's constant to be equal to the sum of Kq and c1. %v != %v + %v",
			sle3AsSLE.C,
			sle1.C,
			K1,
		)
	}

	for LIndex := 0; LIndex < sle3AsSLE.L.Len(); LIndex++ {
		if sle3AsSLE.L.AtVec(LIndex) != sle1.L.AtVec(LIndex) {
			t.Errorf(
				"The linear vector multiplying X was expected to be the same for sle1 and sle3, but sle3[%v] = %v != %v = sle1[%v]",
				LIndex,
				sle3AsSLE.L.AtVec(LIndex),
				LIndex,
				sle1.L.AtVec(LIndex),
			)
		}
	}
}

/*
TestScalarLinearExpression_Plus3
Description:

	Tests whether or not the Plus() function works for a linear expression and a quadratic one containing
	slightly different variables.
*/
func TestScalarLinearExpression_Plus3(t *testing.T) {
	// Constants
	m := optim.NewModel("Plus3")

	v1 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v2 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v3 := m.AddVariableClassic(-10, 10, optim.Continuous)

	Q1_aoa := [][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
	}

	L1_a := []float64{1.0, 7.0}

	C1 := 3.14

	// Preparing constants for NewQuadraticExpr
	Q1_vals := append(Q1_aoa[0], Q1_aoa[1]...)
	Q1 := *mat.NewDense(2, 2, Q1_vals)

	L1 := *mat.NewVecDense(2, L1_a)

	vv1 := optim.VarVector{
		[]optim.Variable{v1, v2},
	}

	// Quantities for Second Expression
	L2 := *mat.NewVecDense(2, []float64{2.0, 11.0})
	C2 := 1.25

	vv2 := optim.VarVector{
		[]optim.Variable{v2, v3},
	}

	// Algorithm
	qe1, err := optim.NewQuadraticExpr(Q1, L1, C1, vv1)
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	le2 := optim.ScalarLinearExpr{
		L: L2,
		C: C2,
		X: vv2,
	}

	e3, err := le2.Plus(qe1)
	if err != nil {
		t.Errorf("There was an issue adding qe1 and le2: %v", err)
	}

	qv3, ok := e3.(optim.ScalarQuadraticExpression)
	if !ok {
		t.Errorf("Unable to convert expression to Quadratic Expression.")
	}

	// Number of variables for this quadratic expression should be 2
	if qv3.NumVars() != 3 {
		t.Errorf("Expected for 3 variable to be found in quadratic expression; function says %v variables exist.", qv3.NumVars())
	}

	if qv3.L.AtVec(0) != qe1.L.AtVec(0) {
		t.Errorf("Expected for L's 0-th element to be 1.0; received %v", qv3.L.AtVec(0))
	}

	if qv3.L.AtVec(1) != qe1.L.AtVec(1)+le2.L.AtVec(0) {
		t.Errorf("Expected for L's 1-th element to be 5.0; received %v", qv3.L.AtVec(1))
	}

	if qv3.L.AtVec(2) != le2.L.AtVec(1) {
		t.Errorf("Expected for L's 2-th element to be 11.0; received %v", qv3.L.AtVec(2))
	}

	if qv3.C != qe1.C+le2.C {
		t.Errorf("Expected for constant of final quadratic expression to be %v; received %v", qe1.C+le2.C, qv3.C)
	}

}

/*
TestScalarLinearExpression_Plus4
Description:

	Tests whether or not the Plus() function works for a linear expression and a single variable that is not known
	slightly different variables.
*/
func TestScalarLinearExpression_Plus4(t *testing.T) {
	// Constants
	m := optim.NewModel("Plus4")

	v1 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v2 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v3 := m.AddVariableClassic(-10, 10, optim.Continuous)

	// Quantities for Second Expression
	L2 := *mat.NewVecDense(2, []float64{2.0, 11.0})
	C2 := 1.25

	vv2 := optim.VarVector{
		[]optim.Variable{v2, v3},
	}

	// Algorithm
	le2 := &optim.ScalarLinearExpr{
		L: L2,
		C: C2,
		X: vv2,
	}

	e3, err := le2.Plus(v1)
	if err != nil {
		t.Errorf("There was an issue adding qe1 and le2: %v", err)
	}

	sle3, ok := e3.(optim.ScalarLinearExpr)
	if !ok {
		t.Errorf("Unable to convert expression to Quadratic Expression.")
	}

	// Number of variables for this quadratic expression should be 2
	if sle3.NumVars() != 3 {
		t.Errorf("Expected for 3 variable to be found in quadratic expression; function says %v variables exist.", sle3.NumVars())
	}

	if sle3.L.AtVec(0) != 1.0 {
		t.Errorf("Expected for L's 0-th element to be 1.0; received %v", sle3.L.AtVec(0))
	}

	if sle3.L.AtVec(1) != le2.L.AtVec(0) {
		t.Errorf("Expected for L's 1-th element to be 5.0; received %v", sle3.L.AtVec(1))
	}

	if sle3.L.AtVec(2) != le2.L.AtVec(1) {
		t.Errorf("Expected for L's 2-th element to be 11.0; received %v", sle3.L.AtVec(2))
	}

	if sle3.C != le2.C {
		t.Errorf("Expected for constant of final quadratic expression to be %v; received %v", le2.C+le2.C, sle3.C)
	}

}

/*
TestScalarLinearExpression_Plus5
Description:

	Verifies that Plus() throws an error when the scalar linear expression
	is not well-formed.
*/
func TestScalarLinearExpression_Plus5(t *testing.T) {
	// Constants
	m := optim.NewModel("TestSLE Plus5")
	N := 4

	sle1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(N + 1),
		X: m.AddVariableVector(N),
		C: -1.0,
	}

	// Algorithm
	_, err := sle1.Plus(21.0)
	if err == nil {
		t.Errorf(
			"No error was thrown with a malformed plus, but there should have been one!",
		)
	} else {
		if !strings.Contains(
			err.Error(),
			fmt.Sprintf(
				"the length of L (%v) does not match that of X (%v)!",
				sle1.L.Len(),
				sle1.X.Len(),
			),
		) {
			t.Errorf("Unexpected error: %v", err)
		}
	}

}

/*
TestScalarLinearExpression_Plus6
Description:

	Verifies that Plus() throws an error when a non-nil
	error is provided.
*/
func TestScalarLinearExpression_Plus6(t *testing.T) {
	// Constants
	m := optim.NewModel("TestSLE Plus6")
	N := 4

	sle1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(N),
		X: m.AddVariableVector(N),
		C: -1.0,
	}

	// Algorithm
	_, err := sle1.Plus(21.0, fmt.Errorf("test"))
	if err == nil {
		t.Errorf(
			"No error was thrown with a malformed plus, but there should have been one!",
		)
	} else {
		if !strings.Contains(
			err.Error(),
			"test",
		) {
			t.Errorf("Unexpected error: %v", err)
		}
	}

}

/*
TestScalarLinearExpression_Plus7
Description:

	Verifies that Plus() throws an error when a nil
	error is provided.
*/
func TestScalarLinearExpression_Plus7(t *testing.T) {
	// Constants
	m := optim.NewModel("TestSLE Plus7")
	N := 4

	var err0 error

	sle1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(N),
		X: m.AddVariableVector(N),
		C: -1.0,
	}

	// Algorithm
	sum, err := sle1.Plus(21.0, err0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	sumAsSLE, tf := sum.(optim.ScalarLinearExpr)
	if !tf {
		t.Errorf(
			"Expected sum to be of type ScalarLinearExpr; received %T",
			sum,
		)
	}

	if sumAsSLE.C != 20.0 {
		t.Errorf(
			"sumAsSLE.C = %v =/= %v as expected",
			sumAsSLE.C,
			20.0,
		)
	}

}

/*
TestScalarLinearExpression_Plus8
Description:

	Verifies that Plus() throws an error when a bad
	input is provided.
*/
func TestScalarLinearExpression_Plus8(t *testing.T) {
	// Constants
	m := optim.NewModel("TestSLE Plus8")
	N := 4

	var err0 error

	sle1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(N),
		X: m.AddVariableVector(N),
		C: -1.0,
	}

	// Algorithm
	_, err := sle1.Plus(err0)
	if err == nil {
		t.Errorf(
			"No error was thrown, but it should have been!",
		)
	} else {
		if !strings.Contains(
			err.Error(),
			optim.UnexpectedInputError{
				InputInQuestion: err0,
				Operation:       "Plus",
			}.Error(),
		) {
			t.Errorf("Unexpected error: %v", err)
		}
	}

}

/*
TestScalarLinearExpr_Multiply1
Description:

	Tests the Multiply() function for an input of a float.
*/
func TestScalarLinearExpr_Multiply1(t *testing.T) {
	// Constants
	m := optim.NewModel("Test-sle-Multiply1")
	v1 := m.AddVariable()
	f1 := 3.14

	sle, _ := v1.Multiply(f1)
	sle1 := sle.(optim.ScalarLinearExpr)

	// Test Multiply
	prod, err := sle1.Multiply(f1)
	if err != nil {
		t.Errorf("There was an issue using Multiply: %v", err)

	}

	prodAsSLE, ok1 := prod.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf(
			"Expected product to be a ScalarLinearExpr; received %T",
			prod,
		)
	}

	if prodAsSLE.X.Len() != 1 {
		t.Errorf(
			"Expected product to contain only one variable; received %v",
			prodAsSLE.X.Len(),
		)
	}

	if prodAsSLE.L.AtVec(0) != f1*f1 {
		t.Errorf(
			"Expected the single coefficient to be %v; received %v",
			f1*f1,
			prodAsSLE.L.AtVec(0),
		)
	}

	if prodAsSLE.C != 0.0 {
		t.Errorf(
			"Expected constant offset to be 0.0; received %v",
			prodAsSLE.C,
		)
	}
}

/*
TestScalarLinearExpr_Multiply2
Description:

	Tests the Multiply() function for an input of a constant K.
*/
func TestScalarLinearExpr_Multiply2(t *testing.T) {
	// Constants
	m := optim.NewModel("Test-sle-Multiply2")
	v1 := m.AddVariable()
	k1 := optim.K(3.14)

	sle, _ := v1.Multiply(k1)
	sle1 := sle.(optim.ScalarLinearExpr)

	// Test Multiply
	prod, err := sle1.Multiply(k1)
	if err != nil {
		t.Errorf("There was an issue using Multiply: %v", err)

	}

	prodAsSLE, ok1 := prod.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf(
			"Expected product to be a ScalarLinearExpr; received %T",
			prod,
		)
	}

	if prodAsSLE.X.Len() != 1 {
		t.Errorf(
			"Expected product to contain only one variable; received %v",
			prodAsSLE.X.Len(),
		)
	}

	if prodAsSLE.L.AtVec(0) != float64(k1)*float64(k1) {
		t.Errorf(
			"Expected the single coefficient to be %v; received %v",
			float64(k1)*float64(k1),
			prodAsSLE.L.AtVec(0),
		)
	}

	if prodAsSLE.C != 0.0 {
		t.Errorf(
			"Expected constant offset to be 0.0; received %v",
			prodAsSLE.C,
		)
	}
}

/*
TestScalarLinearExpr_Multiply3
Description:

	Tests the Multiply() function for an input of a variable v.
*/
func TestScalarLinearExpr_Multiply3(t *testing.T) {
	// Constants
	m := optim.NewModel("Test-sle-Multiply3")
	v1 := m.AddVariable()

	sle, _ := v1.Multiply(3.14)
	sle1 := sle.(optim.ScalarLinearExpr)

	// Test Multiply
	prod, err := sle1.Multiply(v1)
	if err != nil {
		t.Errorf("There was an issue using Multiply: %v", err)

	}

	prodAsSQE, ok1 := prod.(optim.ScalarQuadraticExpression)
	if !ok1 {
		t.Errorf(
			"Expected product to be a ScalarLinearExpr; received %T",
			prod,
		)
	}

	if prodAsSQE.X.Len() != 1 {
		t.Errorf(
			"Expected product to contain only one variable; received %v",
			prodAsSQE.X.Len(),
		)
	}

	// Check constants
	if prodAsSQE.Q.At(0, 0) != 3.14 {
		t.Errorf(
			"Expected Q.At(0,0) to be %v; received %v",
			3.14,
			prodAsSQE.Q.At(0, 0),
		)
	}
	if prodAsSQE.L.AtVec(0) != 0.0 {
		t.Errorf(
			"Expected the single coefficient to be %v; received %v",
			3.14,
			prodAsSQE.L.AtVec(0),
		)
	}

	if prodAsSQE.C != 0.0 {
		t.Errorf(
			"Expected constant offset to be 0.0; received %v",
			prodAsSQE.C,
		)
	}
}

/*
TestScalarLinearExpr_Multiply4
Description:

	Tests the Multiply() function for an input of a variable v.
	Input sle has offset.
*/
func TestScalarLinearExpr_Multiply4(t *testing.T) {
	// Constants
	m := optim.NewModel("Test-sle-Multiply4")
	v1 := m.AddVariable()

	sle, _ := v1.Multiply(3.14)
	sle1 := sle.(optim.ScalarLinearExpr)
	sle2, err := sle1.Plus(2.71)
	if err != nil {
		t.Errorf(
			"Failed to add sle1 with 2.71: %v", err,
		)
	}
	sle3 := sle2.(optim.ScalarLinearExpr)

	// Test Multiply
	prod, err := sle3.Multiply(v1)
	if err != nil {
		t.Errorf("There was an issue using Multiply: %v", err)

	}

	prodAsSQE, ok1 := prod.(optim.ScalarQuadraticExpression)
	if !ok1 {
		t.Errorf(
			"Expected product to be a ScalarLinearExpr; received %T",
			prod,
		)
	}

	if prodAsSQE.X.Len() != 1 {
		t.Errorf(
			"Expected product to contain only one variable; received %v",
			prodAsSQE.X.Len(),
		)
	}

	// Check constants
	if prodAsSQE.Q.At(0, 0) != 3.14 {
		t.Errorf(
			"Expected Q.At(0,0) to be %v; received %v",
			3.14,
			prodAsSQE.Q.At(0, 0),
		)
	}
	if prodAsSQE.L.AtVec(0) != sle3.C {
		t.Errorf(
			"Expected the single coefficient to be %v; received %v",
			3.14,
			prodAsSQE.L.AtVec(0),
		)
	}

	if prodAsSQE.C != 0.0 {
		t.Errorf(
			"Expected constant offset to be %v; received %v",
			0.0,
			prodAsSQE.C,
		)
	}
}

/*
TestScalarLinearExpr_Multiply5
Description:

	Tests the Multiply() function for an input of a variable (different from original).
*/
func TestScalarLinearExpr_Multiply5(t *testing.T) {
	// Constants
	m := optim.NewModel("Test-sle-Multiply5")
	v1 := m.AddVariable()
	v2 := m.AddVariable()

	sle, _ := v1.Multiply(3.14)
	sle1 := sle.(optim.ScalarLinearExpr)
	sle2, err := sle1.Plus(2.71)
	if err != nil {
		t.Errorf(
			"Failed to add sle1 with 2.71: %v", err,
		)
	}
	sle3 := sle2.(optim.ScalarLinearExpr)

	// Test Multiply
	prod, err := sle3.Multiply(v2)
	if err != nil {
		t.Errorf("There was an issue using Multiply: %v", err)

	}

	prodAsSQE, ok1 := prod.(optim.ScalarQuadraticExpression)
	if !ok1 {
		t.Errorf(
			"Expected product to be a ScalarLinearExpr; received %T",
			prod,
		)
	}

	if prodAsSQE.X.Len() != 2 {
		t.Errorf(
			"Expected product to contain only one variable; received %v",
			prodAsSQE.X.Len(),
		)
	}

	// Check constants
	if prodAsSQE.Q.At(0, 1) != 3.14*0.5 {
		t.Errorf(
			"Expected Q.At(0,1) to be %v; received %v",
			3.14*0.5,
			prodAsSQE.Q.At(0, 1),
		)
	}
	if prodAsSQE.Q.At(1, 0) != 3.14*0.5 {
		t.Errorf(
			"Expected Q.At(1,0) to be %v; received %v",
			3.14*0.5,
			prodAsSQE.Q.At(1, 0),
		)
	}

	if prodAsSQE.Q.At(0, 0) != 0.0 {
		t.Errorf(
			"Expected Q.At(0,0) to be %v; received %v",
			0.0,
			prodAsSQE.Q.At(0, 0),
		)
	}
	if prodAsSQE.Q.At(1, 1) != 0.0 {
		t.Errorf(
			"Expected Q.At(1,1) to be %v; received %v",
			0.0,
			prodAsSQE.Q.At(1, 1),
		)
	}

	if prodAsSQE.L.AtVec(1) != sle3.C {
		t.Errorf(
			"Expected the single coefficient to be %v; received %v",
			3.14,
			prodAsSQE.L.AtVec(0),
		)
	}

	if prodAsSQE.C != 0.0 {
		t.Errorf(
			"Expected constant offset to be %v; received %v",
			0.0,
			prodAsSQE.C,
		)
	}
}

/*
TestScalarLinearExpr_Multiply6
Description:

	Tests the Multiply() function for an input of a scalar linear expression.
*/
func TestScalarLinearExpr_Multiply6(t *testing.T) {
	// Constants
	m := optim.NewModel("Test-sle-Multiply5")
	v1 := m.AddVariable()

	sle, _ := v1.Multiply(3.14)
	sle1 := sle.(optim.ScalarLinearExpr)
	sle2, err := sle1.Plus(2.71)
	if err != nil {
		t.Errorf(
			"Failed to add sle1 with 2.71: %v", err,
		)
	}
	sle3 := sle2.(optim.ScalarLinearExpr)

	// Test Multiply
	prod, err := sle3.Multiply(sle3)
	if err != nil {
		t.Errorf("There was an issue using Multiply: %v", err)

	}

	prodAsSQE, ok1 := prod.(optim.ScalarQuadraticExpression)
	if !ok1 {
		t.Errorf(
			"Expected product to be a ScalarLinearExpr; received %T",
			prod,
		)
	}

	if prodAsSQE.X.Len() != 1 {
		t.Errorf(
			"Expected product to contain only one variable; received %v",
			prodAsSQE.X.Len(),
		)
	}

	// Check constants
	if prodAsSQE.Q.At(0, 0) != 3.14*3.14 {
		t.Errorf(
			"Expected Q.At(0,0) to be %v; received %v",
			0.0,
			prodAsSQE.Q.At(0, 0),
		)
	}

	if prodAsSQE.L.AtVec(0) != sle3.C*3.14*2.0 {
		t.Errorf(
			"Expected the single coefficient to be %v; received %v",
			3.14,
			prodAsSQE.L.AtVec(0),
		)
	}

	if prodAsSQE.C != sle3.C*sle3.C {
		t.Errorf(
			"Expected constant offset to be %v; received %v",
			sle3.C*sle3.C,
			prodAsSQE.C,
		)
	}
}

/*
TestScalarLinearExpr_Multiply7
Description:

	Tests the Multiply() function for an input of a scalar linear expression.
*/
func TestScalarLinearExpr_Multiply7(t *testing.T) {
	// Constants
	m := optim.NewModel("Test-sle-Multiply7")
	v1 := m.AddVariable()
	v2 := m.AddVariable()

	sle, _ := v1.Multiply(3.14)
	sle1 := sle.(optim.ScalarLinearExpr)
	sle2, err := sle1.Multiply(v2)
	if err != nil {
		t.Errorf(
			"Failed to add sle1 with 2.71: %v", err,
		)
	}
	sle3 := sle2.(optim.ScalarQuadraticExpression)

	// Test Multiply
	_, err = sle1.Multiply(sle3)
	if err == nil {
		t.Errorf("There should be an an issue using Multiply; received none")

	}

	if strings.Compare(
		err.Error(),
		"Can not multiply ScalarLinearExpr with ScalarQuadraticExpression. MatProInterface can not represent polynomials higher than degree 2.",
	) != 0 {
		t.Errorf(
			"Expected for specific error to occur, but received %v", err)
	}

}

/*
TestScalarLinearExpr_Multiply8
Description:

	Verifies that a malformed scalar linear expression throws an
	error when used in a multiply().
*/
func TestScalarLinearExpr_Multiply8(t *testing.T) {
	// Constants
	N := 4
	m := optim.NewModel("Multiply8")
	vv1 := m.AddVariableVector(N)

	sle1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(N - 1),
		X: vv1,
		C: 2.14,
	}

	// Multiply!
	_, err := sle1.Multiply(3.14)
	if err == nil {
		t.Errorf("there were no errors, but we expected some!")
	} else {
		if !strings.Contains(
			err.Error(),
			fmt.Sprintf(
				"the length of L (%v) does not match that of X (%v)!",
				sle1.L.Len(),
				sle1.X.Len(),
			),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestScalarLinearExpr_Multiply9
Description:

	Verifies that a multiplication with a vector of non-unit length
	throws an error when used in a multiply().
*/
func TestScalarLinearExpr_Multiply9(t *testing.T) {
	// Constants
	N := 4
	m := optim.NewModel("TestSLE-Multiply9")
	vv1 := m.AddVariableVector(N)

	sle1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(N),
		X: vv1,
		C: 2.14,
	}

	kv2 := optim.KVector(optim.OnesVector(N))

	// Multiply!
	_, err := sle1.Multiply(kv2)
	if err == nil {
		t.Errorf("there were no errors, but we expected some!")
	} else {
		if !strings.Contains(
			err.Error(),
			optim.DimensionError{
				Operation: "Multiply",
				Arg1:      sle1,
				Arg2:      kv2,
			}.Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestScalarLinearExpr_Multiply10
Description:

	Verifies that a multiplication with a KVector of unit length
	doesn't throw an error.
*/
func TestScalarLinearExpr_Multiply10(t *testing.T) {
	// Constants
	N := 2
	m := optim.NewModel("TestSLE-Multiply10")
	vv1 := m.AddVariableVector(N)

	sle1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(N),
		X: vv1,
		C: 2.14,
	}

	kv2 := optim.KVector(optim.OnesVector(1))

	// Multiply!
	prod, err := sle1.Multiply(kv2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	prodAsSLE, tf := prod.(optim.ScalarLinearExpr)
	if !tf {
		t.Errorf(
			"Expected product to be a ScalarLinearExpr; received %T",
			prod,
		)
	}

	if prodAsSLE.C != sle1.C {
		t.Errorf(
			"prodAsSLE.C = %v =/= %v = sle1.C",
			prodAsSLE.C,
			sle1.C,
		)
	}
}

/*
TestScalarLinearExpr_NewLinearExpr1
Description:

	Tests that the scalar linear expression is properly created by NewLinearExpr.
*/
func TestScalarLinearExpr_NewLinearExpr1(t *testing.T) {
	// Constants
	se := optim.NewLinearExpr(2.1)

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
TestScalarLinearExpr_Variables1
Description:

	Verifies that this function works well for SLE's of one variable.
*/
func TestScalarLinearExpr_Variables1(t *testing.T) {
	// Constants
	m := optim.NewModel("testSLE-variables1")
	vv1 := m.AddVariableVector(1)
	sle1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(1),
		X: vv1,
		C: 0.1,
	}

	// Algorithm
	varsOut := sle1.Variables()
	if len(varsOut) != 1 {
		t.Errorf(
			"Expected for a variablevector of length 1; received length %v",
			len(varsOut),
		)
	}

	if varsOut[0].ID != vv1.Elements[0].ID {
		t.Errorf(
			"Expected first element of varsOut (%v) to be the same as the first element of vv1 (%v). They were not!",
			varsOut[0],
			vv1.Elements[0],
		)
	}
}

/*
TestScalarLinearExpr_Variables2
Description:

	Verifies that this function works well for SLE's of ten variable.
*/
func TestScalarLinearExpr_Variables2(t *testing.T) {
	// Constants
	m := optim.NewModel("testSLE-variables1")
	n := 11
	vv1 := m.AddVariableVector(n)
	sle1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(n),
		X: vv1,
		C: 0.1,
	}

	// Algorithm
	varsOut := sle1.Variables()
	if len(varsOut) != n {
		t.Errorf(
			"Expected for a variablevector of length %v; received length %v",
			n, len(varsOut),
		)
	}

	for vv1Index := 0; vv1Index < vv1.Len(); vv1Index++ {
		// check each element
		if vv1.Elements[vv1Index] != varsOut[vv1Index] {
			t.Errorf(
				"Expected %v-th element of varsOut (%v) to be the same as vv1[%v]. IT wasn't!",
				vv1Index,
				varsOut[vv1Index],
				vv1.Elements[vv1Index],
			)
		}
	}

}
