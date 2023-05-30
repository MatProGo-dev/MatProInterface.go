package optim_test

import (
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

//func TestLinearExprCoeffsAndConstant(t *testing.T) {
//	m := optim.NewModel()
//	x := m.AddBinaryVar()
//	y := m.AddBinaryVar()
//
//	// 2 * x + 4 * y - 5
//	coeffs := []float64{2, 4}
//	constant := -5.0
//	expr := optim.Sum(x.Mult(coeffs[0]), y.Mult(coeffs[1]), optim.K(constant))
//
//	for i, coeff := range expr.Coeffs() {
//		if coeffs[i] != coeff {
//			t.Errorf("Coeff mismatch: %v != %v", coeff, coeffs[i])
//		}
//	}
//
//	if expr.Constant() != constant {
//		t.Errorf("Constant mismatch: %v != %v", expr.Constant(), constant)
//	}
//}

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
		[]float64{1.0, 2.0},
		[]float64{3.0, 4.0},
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
