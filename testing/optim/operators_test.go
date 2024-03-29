package optim

import (
	"fmt"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"strings"
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
TestOperators_Comparison1
Description:

	Tests whether or not Comparison works for two valid expressions.
*/
func TestOperators_Comparison1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-comparison1")
	vec1 := m.AddVariable()
	vec2 := m.AddVariable()
	c1 := optim.K(1.2)
	e2, err := c1.Plus(vec2)
	if err != nil {
		t.Errorf("There was an issue adding vec2 to c1: %v", err)
	}

	// Algorithms
	constr0, err := optim.Comparison(vec1, e2, optim.SenseGreaterThanEqual)
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
TestOperators_Comparison2
Description:

	Tests whether or not Comparison works for a valid expression and
	a boolean.
*/
func TestOperators_Comparison2(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-comparison2")
	vec1 := m.AddVariable()
	vec2 := m.AddVariable()
	c1 := optim.K(1.2)
	err0 := fmt.Errorf("Comparison err 2")
	e2, err := c1.Plus(vec1.Plus(vec2))
	if err != nil {
		t.Errorf("There was an issue adding vec2 to c1: %v", err)
	}

	// Algorithms
	_, err = optim.Comparison(err0, e2, optim.SenseGreaterThanEqual)
	if strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Comparison in sense '%v' is not defined for lhs type %T and rhs type %T!",
			optim.SenseGreaterThanEqual, e2, err0,
		),
	) {
		t.Errorf("The Eq() comparison appears to be equal to the two vectors: %v", err)
	}
}

/*
TestOperators_Comparison3
Description:

	Tests whether or not Comparison works for a valid expression and
	a boolean.
*/
func TestOperators_Comparison3(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-comparison3")
	vv1 := m.AddVariableVector(10)
	kv1 := optim.OnesVector(vv1.Len())

	// Algorithms
	constr, err := optim.Comparison(vv1, kv1, optim.SenseLessThanEqual)
	if err != nil {
		t.Errorf("There was an error computing Comparison(): %v", err)
	}

	if constr.(optim.VectorConstraint).Sense != optim.SenseLessThanEqual {
		t.Errorf(
			"Expected sense of constraint to be %v; received %v",
			optim.SenseLessThanEqual, constr.(optim.ScalarConstraint).Sense,
		)
	}
}

/*
TestOperators_Comparison4
Description:

	Tests whether or not Comparison works for two valid expressions.
*/
func TestOperators_Comparison4(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-comparison4")
	v1 := m.AddVariable()
	f1 := 2.3

	// Algorithms
	constr0, err := optim.Comparison(f1, v1, optim.SenseGreaterThanEqual)
	if err != nil {
		t.Errorf("The Eq() comparison appears to be equal to the two vectors: %v", err)
	}

	sc0, _ := constr0.(optim.ScalarConstraint)
	if _, ok1 := sc0.LeftHandSide.(optim.K); !ok1 {
		t.Errorf("The left hand side of the equality is not a variable!")
	}

	if _, ok2 := sc0.RightHandSide.(optim.Variable); !ok2 {
		t.Errorf("The right hand side of the equality is not a variable!")
	}

}

/*
TestOperators_Comparison5
Description:

	Tests whether or not Comparison works for a valid expression and
	a boolean.
*/
func TestOperators_Comparison5(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-comparison3")
	vv1 := m.AddVariableVector(10)
	kv1 := optim.OnesVector(vv1.Len())

	// Algorithms
	constr, err := optim.Comparison(kv1, vv1, optim.SenseLessThanEqual)
	if err != nil {
		t.Errorf("There was an error computing Comparison(): %v", err)
	}

	if constr.(optim.VectorConstraint).Sense != optim.SenseLessThanEqual {
		t.Errorf(
			"Expected sense of constraint to be %v; received %v",
			optim.SenseLessThanEqual, constr.(optim.ScalarConstraint).Sense,
		)
	}
}

/*
TestOperators_Comparison6
Description:

	Tests whether or not Comparison works for a valid expression and
	a boolean.
*/
func TestOperators_Comparison6(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-comparison2")
	vec1 := m.AddVariable()
	vec2 := m.AddVariable()
	c1 := optim.K(1.2)
	e2, err := c1.Plus(vec1.Plus(vec2))
	if err != nil {
		t.Errorf("There was an issue adding vec2 to c1: %v", err)
	}

	// Algorithms
	constr0, err := optim.Comparison(e2, c1, optim.SenseGreaterThanEqual)
	if err != nil {
		t.Errorf("Unexpected error computing comparison: %v", err)
	}

	scalarConstr0 := constr0.(optim.ScalarConstraint)
	_, ok1 := scalarConstr0.LeftHandSide.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("Expected LHS to be a ScalarLinearExpr but received %T.", scalarConstr0.LeftHandSide)
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

/*
TestOperators_Multiply3
Description:
*/
func TestOperators_Multiply3(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-multiply3")
	v1 := m.AddVariable()
	f1 := 6.4

	// Algorithm
	prod1, err := optim.Multiply(f1, v1)
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
TestOperators_Multiply4
Description:

	Multiply a vecdense with a VarVector. Idk if this will work at all?
*/
func TestOperators_Multiply4(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-multiply3")
	v1 := m.AddVariableVector(3)
	kv1 := optim.OnesVector(v1.Len())

	// Algorithm
	_, err := optim.Multiply(kv1, v1)
	if !strings.Contains(
		err.Error(),
		optim.DimensionError{
			Operation: "Multiply",
			Arg1:      optim.KVector(kv1),
			Arg2:      v1,
		}.Error(),
	) {
		t.Errorf("Unexpected error value for multiply: %v", err)
	}

}

/*
TestOperators_Multiply5
Description:

	Multiply a KVectorTranspose with a VarVector. Should return a LinearExpression value
*/
func TestOperators_Multiply5(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-multiply3")
	v1 := m.AddVariableVector(3)
	kv1 := optim.KVector(optim.OnesVector(v1.Len()))

	// Algorithm
	prod, err := optim.Multiply(kv1.Transpose(), v1)
	if err != nil {
		t.Errorf("Unexpected error during valid multiplication: %v", err)
	}

	prodAsSLE, ok := prod.(optim.ScalarLinearExpr)
	if !ok {
		t.Errorf("Expected product to be a ScalarLinearExpr; received %T", prod)
	}

	if prodAsSLE.X.Len() != v1.Len() {
		t.Errorf(
			"Expected product to contain %v variables; received %v",
			v1.Len(),
			prodAsSLE.X.Len(),
		)
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

/*
TestOperators_Sum2
Description:

	Sums two expressions together. Tests whether or not error handling works when
	first argument is not an expression.
*/
func TestOperators_Sum2(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-sum1")

	f1 := 3.1
	v1 := m.AddVariable()
	tf1 := true

	// Test
	_, err := optim.Sum(tf1, v1, f1)
	if err == nil {
		t.Errorf("Expected there to be an error when parsing, but there was none!")
	}

	if !strings.Contains(err.Error(), "The first input to Sum must be an expression! Received type") {
		t.Errorf(
			"Wrong error detected in Sum: %v", err,
		)
	}
}

/*
TestOperators_Sum3
Description:

	Sums two expressions together. Tests whether or not Sum of single expression is
	properly returned.
*/
func TestOperators_Sum3(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-sum3")

	f1 := 3.1
	v1 := m.AddVariableVector(10)
	se1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(v1.Len()),
		X: v1,
		C: f1,
	}

	// Test
	sumOut, err := optim.Sum(se1)
	if err != nil {
		t.Errorf(
			"There was an issue computing the sum: %v",
			err,
		)
	}

	sumAsSLE, ok1 := sumOut.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("There was an issue converting sum to SLE!")
	}

	// Check that sum is the same as SLE
	for Lindex := 0; Lindex < v1.Len(); Lindex++ {
		if sumAsSLE.L.AtVec(Lindex) != se1.L.AtVec(Lindex) {
			t.Errorf(
				"Expected L[%v] = %v; received %v",
				Lindex, se1.L.AtVec(Lindex),
				sumAsSLE.L.AtVec(Lindex),
			)
		}
		if sumAsSLE.X.AtVec(Lindex).IDs()[0] != se1.X.AtVec(Lindex).IDs()[0] {
			t.Errorf(
				"One of the variable's ids in the sle %v is not the same as what is in the sum %v!",
				se1.X.AtVec(Lindex).IDs()[0],
				sumAsSLE.X.AtVec(Lindex).IDs()[0],
			)
		}
	}

	if sumAsSLE.C != se1.C {
		t.Errorf(
			"Expected offset to be %v; received %v.",
			se1.C,
			sumAsSLE.C,
		)
	}
}

/*
TestOperators_Sum4
Description:

	Sums two expressions together. Tests whether or not single expression input
	with error input is properly returned (when error is nil).
*/
func TestOperators_Sum4(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-sum4")

	f1 := 3.1
	v1 := m.AddVariableVector(10)
	se1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(v1.Len()),
		X: v1,
		C: f1,
	}

	// Test
	var err error = nil
	sumOut, err := optim.Sum(se1, err)
	if err != nil {
		t.Errorf(
			"There was an issue computing the sum: %v",
			err,
		)
	}

	sumAsSLE, ok1 := sumOut.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("There was an issue converting sum to SLE!")
	}

	// Check that sum is the same as SLE
	for Lindex := 0; Lindex < v1.Len(); Lindex++ {
		if sumAsSLE.L.AtVec(Lindex) != se1.L.AtVec(Lindex) {
			t.Errorf(
				"Expected L[%v] = %v; received %v",
				Lindex, se1.L.AtVec(Lindex),
				sumAsSLE.L.AtVec(Lindex),
			)
		}
		if sumAsSLE.X.AtVec(Lindex).IDs()[0] != se1.X.AtVec(Lindex).IDs()[0] {
			t.Errorf(
				"One of the variable's ids in the sle %v is not the same as what is in the sum %v!",
				se1.X.AtVec(Lindex).IDs()[0],
				sumAsSLE.X.AtVec(Lindex).IDs()[0],
			)
		}
	}

	if sumAsSLE.C != se1.C {
		t.Errorf(
			"Expected offset to be %v; received %v.",
			se1.C,
			sumAsSLE.C,
		)
	}
}

/*
TestOperators_Sum5
Description:

	Sums two expressions together. Tests whether or not single expression input
	with error input is properly errored (when error is NOT nil).
*/
func TestOperators_Sum5(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-sum4")

	f1 := 3.1
	v1 := m.AddVariableVector(10)
	se1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(v1.Len()),
		X: v1,
		C: f1,
	}

	// Test
	err0 := fmt.Errorf("My custom error!")
	_, err := optim.Sum(se1, err0)
	if !strings.Contains(err.Error(), err0.Error()) {
		t.Errorf(
			"Expected error \"%v\"; received %v",
			err0, err,
		)
	}

}

/*
TestOperators_Sum6
Description:

	Sums two expressions together. Tests whether or not two expression input
	with error input is properly returned (when error between is nil).
*/
func TestOperators_Sum6(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-sum4")

	f1 := 3.1
	v1 := m.AddVariableVector(10)
	se1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(v1.Len()),
		X: v1,
		C: f1,
	}

	// Test
	var err error = nil
	sumOut, err := optim.Sum(se1, err, f1)
	if err != nil {
		t.Errorf(
			"There was an issue computing the sum: %v",
			err,
		)
	}

	sumAsSLE, ok1 := sumOut.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("There was an issue converting sum to SLE!")
	}

	// Check that sum is the same as SLE
	for Lindex := 0; Lindex < v1.Len(); Lindex++ {
		if sumAsSLE.L.AtVec(Lindex) != se1.L.AtVec(Lindex) {
			t.Errorf(
				"Expected L[%v] = %v; received %v",
				Lindex, se1.L.AtVec(Lindex),
				sumAsSLE.L.AtVec(Lindex),
			)
		}
		if sumAsSLE.X.AtVec(Lindex).IDs()[0] != se1.X.AtVec(Lindex).IDs()[0] {
			t.Errorf(
				"One of the variable's ids in the sle %v is not the same as what is in the sum %v!",
				se1.X.AtVec(Lindex).IDs()[0],
				sumAsSLE.X.AtVec(Lindex).IDs()[0],
			)
		}
	}

	if sumAsSLE.C != se1.C+f1 {
		t.Errorf(
			"Expected offset to be %v; received %v.",
			se1.C,
			sumAsSLE.C,
		)
	}
}

/*
TestOperators_Sum7
Description:

	Sums two expressions together. Tests whether or not two expression input
	with error input is properly returned (when no error between).
*/
func TestOperators_Sum7(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-sum7")

	f1 := 3.1
	v1 := m.AddVariableVector(10)
	se1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(v1.Len()),
		X: v1,
		C: f1,
	}

	// Test
	var err error = nil
	sumOut, err := optim.Sum(se1, f1)
	if err != nil {
		t.Errorf(
			"There was an issue computing the sum: %v",
			err,
		)
	}

	sumAsSLE, ok1 := sumOut.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("There was an issue converting sum to SLE!")
	}

	// Check that sum is the same as SLE
	for Lindex := 0; Lindex < v1.Len(); Lindex++ {
		if sumAsSLE.L.AtVec(Lindex) != se1.L.AtVec(Lindex) {
			t.Errorf(
				"Expected L[%v] = %v; received %v",
				Lindex, se1.L.AtVec(Lindex),
				sumAsSLE.L.AtVec(Lindex),
			)
		}
		if sumAsSLE.X.AtVec(Lindex).IDs()[0] != se1.X.AtVec(Lindex).IDs()[0] {
			t.Errorf(
				"One of the variable's ids in the sle %v is not the same as what is in the sum %v!",
				se1.X.AtVec(Lindex).IDs()[0],
				sumAsSLE.X.AtVec(Lindex).IDs()[0],
			)
		}
	}

	if sumAsSLE.C != se1.C+f1 {
		t.Errorf(
			"Expected offset to be %v; received %v.",
			se1.C,
			sumAsSLE.C,
		)
	}
}

/*
TestOperators_Sum8
Description:

	Sums two expressions together. Tests whether or not two vector expression input
	with error input is properly returned (when no error between).
*/
func TestOperators_Sum8(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-sum8")

	//f1 := 3.1
	v1 := m.AddVariableVector(10)
	vc1 := optim.OnesVector(v1.Len())

	// Test
	var err error = nil
	sumOut, err := optim.Sum(v1, vc1)
	if err != nil {
		t.Errorf(
			"There was an issue computing the sum: %v",
			err,
		)
	}

	sumAsVLE, ok1 := sumOut.(optim.VectorLinearExpr)
	if !ok1 {
		t.Errorf("There was an issue converting sum to SLE!")
	}

	// Check that sum is the same as SLE
	for Lindex := 0; Lindex < v1.Len(); Lindex++ {
		//if sumAsVLE.L.AtVec(Lindex) != se1.L.AtVec(Lindex) {
		//	t.Errorf(
		//		"Expected L[%v] = %v; received %v",
		//		Lindex, se1.L.AtVec(Lindex),
		//		sumAsVLE.L.AtVec(Lindex),
		//	)
		//}
		if sumAsVLE.X.AtVec(Lindex).IDs()[0] != v1.AtVec(Lindex).IDs()[0] {
			t.Errorf(
				"One of the variable's ids in the sle %v is not the same as what is in the sum %v!",
				v1.AtVec(Lindex).IDs()[0],
				sumAsVLE.X.AtVec(Lindex).IDs()[0],
			)
		}
		if sumAsVLE.C.AtVec(Lindex) != vc1.AtVec(Lindex) {
			t.Errorf(
				"One of the expression's constants (%v) did not match the constant vector at %v (%v).",
				sumAsVLE.C.AtVec(Lindex),
				Lindex, vc1.AtVec(Lindex),
			)
		}
	}
}

/*
TestOperators_Sum9
Description:

	Sums two expressions together. Tests whether or not sum of expression
	and bool throws error (as expected).
*/
func TestOperators_Sum9(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-sum9")

	b1 := false
	v1 := m.AddVariableVector(10)
	se1 := optim.ScalarLinearExpr{
		L: optim.OnesVector(v1.Len()),
		X: v1,
	}

	// Test
	var err error = nil
	_, err = optim.Sum(se1, b1)
	if !strings.Contains(err.Error(), "Unexpected input to Sum ") {
		t.Errorf("Did not receive expected error!")
	}
}

/*
TestOperators_Sum10
Description:

	Sums two expressions together. Tests whether or not sum of three expressions
	is correct.
*/
func TestOperators_Sum10(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-sum10")

	v1 := m.AddVariable()
	f1 := 1.2
	f2 := 3.8

	// Test
	var err error = nil
	sumOut, err := optim.Sum(v1, f1, f2)
	if err != nil {
		t.Errorf(
			"Unexpected error in sum: %v", err,
		)
	}

	sumOutAsSLE, tf1 := sumOut.(optim.ScalarLinearExpr)
	if !tf1 {
		t.Errorf("Unexpected type of sum output! %T", sumOutAsSLE)
	}

	if sumOutAsSLE.C != f1+f2 {
		t.Errorf("Expected offset of sum to be %v; recieved %v.", f1+f2, sumOutAsSLE.C)
	}
}
