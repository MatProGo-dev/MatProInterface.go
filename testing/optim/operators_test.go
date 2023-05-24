package optim

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"testing"
)

/*
TestOperators_Eq1
Description:

	Tests whether or not Eq works for two valid expressions.
*/
func TestOperators_Eq1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-eq1")
	vec1 := m.AddVariable()
	vec2 := m.AddVariable()

	// Algorithms
	constr0, err := optim.Eq(vec1, vec2)
	if err != nil {
		t.Errorf("The Eq() comparison appears to be equal to the two vectors: %v", err)
	}

	sc0, _ := constr0.(optim.ScalarConstraint)
	if _, ok1 := sc0.LeftHandSide.(optim.Variable); !ok1 {
		t.Errorf("The left hand side of the equality is not a variable!")
	}

	if _, ok2 := sc0.RightHandSide.(optim.Variable); !ok2 {
		t.Errorf("The right hand side of the equality is not a variable!")
	}

}

/*
TestOperators_LessEq1
Description:

	Tests whether or not a good LessEq comparison successfully is built and contains the right variables.
*/
func TestOperators_LessEq1(t *testing.T) {
	// Constants
	desLength := 5

	m := optim.NewModel("LessEq1")
	var vec1 = m.AddVariableVector(desLength)
	var vec2 = optim.OnesVector(desLength)

	// Algorithm
	constr, err := optim.LessEq(vec1, vec2)
	if err != nil {
		t.Errorf("There was an issue compusing the LessEq comparison: %v", err)
	}

	vecConstr, ok := constr.(optim.VectorConstraint)
	if !ok {
		t.Errorf("expected constraint to be a Vector constraint, but it was really of type %T.", constr)
	}

	lhs := vecConstr.LeftHandSide
	lhsAsVarVector, ok := lhs.(optim.VarVector)
	if !ok {
		t.Errorf("The left hand side was expected to be a VarVector, but instead it was %T.", lhs)
	}

	for varIndex := 0; varIndex < vec1.Len(); varIndex++ {
		if vec1.AtVec(varIndex).(optim.Variable).ID != lhsAsVarVector.AtVec(varIndex).(optim.Variable).ID {
			t.Errorf(
				"vec1's %v-th element (%v) is different from left hand side's %v-th element (%v).",
				varIndex,
				vec1.AtVec(varIndex),
				varIndex,
				lhsAsVarVector.AtVec(varIndex),
			)
		}

	}

}

/*
TestOperators_GreaterEq1
Description:

	Tests whether or not GreaterEq works for two valid expressions.
*/
func TestOperators_GreaterEq1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-greatereq1")
	vec1 := m.AddVariable()
	vec2 := m.AddVariable()
	c1 := optim.K(1.2)
	e2, err := c1.Plus(vec2)
	if err != nil {
		t.Errorf("There was an issue adding vec2 to c1: %v", err)
	}

	// Algorithms
	constr0, err := optim.GreaterEq(vec1, e2)
	if err != nil {
		t.Errorf("The Eq() comparison appears to be equal to the two vectors: %v", err)
	}

	sc0, _ := constr0.(optim.ScalarConstraint)
	if _, ok1 := sc0.LeftHandSide.(optim.Variable); !ok1 {
		t.Errorf("The left hand side of the equality is not a variable!")
	}

	if _, ok2 := sc0.RightHandSide.(optim.ScalarLinearExpr); !ok2 {
		t.Errorf("The right hand side of the equality is not a variable!")
	}

}

/*
TestOperators_Multiply1
Description:
*/
func TestOperators_Multiply1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-multiply1")
	v1 := m.AddVariable()
	f1 := 6.4

	// Algorithm
	prod1, err := optim.Multiply(v1, f1)
	if err != nil {
		t.Errorf("There was an issue using Multiply(): %v", err)
	}

	prod1AsSLE, ok1 := prod1.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("There was an issue converting this product to an SLE! Type %T", prod1)
	}

	if prod1AsSLE.X.Len() != 1 {
		t.Errorf("Expected X to have length 1; received %v", prod1AsSLE.X.Len())
	}

}

/*
TestOperators_Multiply2
Description:

	Tests multiply of scalar float with ScalarLinearExpression
*/
func TestOperators_Multiply2(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-multiply1")
	vv1 := m.AddVariableVector(2)
	sle1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(2),
		X: vv1,
		C: 2.14,
	}
	f1 := 6.4

	// Algorithm
	prod1, err := optim.Multiply(sle1, f1)
	if err != nil {
		t.Errorf("There was an issue using Multiply(): %v", err)
	}

	prod1AsSLE, ok1 := prod1.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("There was an issue converting this product to an SLE! Type %T", prod1)
	}

	if prod1AsSLE.X.Len() != vv1.Len() {
		t.Errorf("Expected X to have length 2; received %v", vv1.Len())
	}

	for vecIndex := 0; vecIndex < 2; vecIndex++ {
		// Check each element of vector coefficient
		if prod1AsSLE.L.AtVec(vecIndex) != 1.0*f1 {
			t.Errorf(
				"Product coefficient L[%v] = %v. Expected %v",
				vecIndex,
				prod1AsSLE.L.AtVec(vecIndex),
				1.0*f1,
			)
		}

	}

}

//func TestDot(t *testing.T) {
//	N := 10
//	m := optim.NewModel("testDot")
//	xs := m.AddBinaryVariableVector(N)
//	coeffs := make([]float64, N)
//
//	for i := 0; i < N; i++ {
//		coeffs[i] = float64(i + 1)
//	}
//
//	expr := optim.Dot(xs.Elements, coeffs)
//
//	for i, coeff := range expr.Coeffs() {
//		if coeffs[i] != coeff {
//			t.Errorf("Coeff mismatch: %v != %v", coeff, coeffs[i])
//		}
//	}
//
//	if expr.Constant() != 0 {
//		t.Errorf("Constant mismatch: %v != 0", expr.Constant())
//	}
//}
//
//func TestDotPanic(t *testing.T) {
//	N := 10
//	m := optim.NewModel("TestDotPanic")
//	xs := m.AddBinaryVariableVector(N)
//	coeffs := make([]float64, N-1)
//
//	for i := 0; i < N-1; i++ {
//		coeffs[i] = float64(i + 1)
//	}
//
//	defer func() {
//		if r := recover(); r == nil {
//			t.Error("Coeff size mismatch: Code did not panic")
//		}
//	}()
//
//	optim.Dot(xs.Elements, coeffs)
//}

/*
TestOperators_Sum1
Description:

	Sums two expressions together.
*/
func TestOperators_Sum1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-sum1")

	f1 := 3.1
	v1 := m.AddVariable()

	// Test
	sum1, err := optim.Sum(v1, f1)
	if err != nil {
		t.Errorf("Sum of variable and float failed! %v", err)
	}

	sum1AsSLE, ok1 := sum1.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf(
			"Did not successuflly cast sum to ScalarLinearExpr; received %T",
			sum1,
		)
	}

	if sum1AsSLE.X.Len() != 1 {
		t.Errorf(
			"Expected sum to be an SLE with 1 variable; received %v variables!",
			sum1AsSLE.X.Len(),
		)
	}
}
