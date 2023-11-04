package optim_test

import (
	"fmt"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"strings"
	"testing"
)

/*
TestVectorExpression_IsVectorExpression1
Description:

	Tests whether or not IsVectorExpression() works on a KVector object.
*/
func TestVectorExpression_IsVectorExpression1(t *testing.T) {
	// Constants
	N := 10
	K1 := optim.KVector(optim.OnesVector(N))

	// Check
	if !optim.IsVectorExpression(K1) {
		t.Errorf("K1 was determined to NOT be a vector expression, but it is!")
	}
}

/*
TestVectorExpression_IsVectorExpression2
Description:

	Tests whether or not IsVectorExpression() works on a KVectorTranspose object.
*/
func TestVectorExpression_IsVectorExpression2(t *testing.T) {
	// Constants
	N := 10
	K1 := optim.KVectorTranspose(optim.OnesVector(N))

	// Check
	if !optim.IsVectorExpression(K1) {
		t.Errorf("K1 was determined to NOT be a vector expression, but it is!")
	}
}

/*
TestVectorExpression_IsVectorExpression3
Description:

	Tests whether or not IsVectorExpression() works on a VarVectorTranspose object.
*/
func TestVectorExpression_IsVectorExpression3(t *testing.T) {
	// Constants
	N := 10
	m := optim.NewModel("IsVectorExpression3")
	vv1 := m.AddVariableVector(N)

	// Check
	if !optim.IsVectorExpression(vv1.Transpose()) {
		t.Errorf(
			"vv1.T was determined to NOT be a vector expression, but it is (type %T)!",
			vv1.Transpose(),
		)
	}
}

/*
TestVectorExpression_IsVectorExpression4
Description:

	Tests whether or not IsVectorExpression() works on a VectorLinearExpression object.
*/
func TestVectorExpression_IsVectorExpression4(t *testing.T) {
	// Constants
	N := 10
	c1 := optim.ZerosVector(N)
	vle1 := optim.NewVectorExpression(c1)

	// Check
	if !optim.IsVectorExpression(vle1) {
		t.Errorf(
			"vle1 was determined to NOT be a vector expression, but it is (type %T)!",
			vle1,
		)
	}
}

/*
TestVectorExpression_IsVectorExpression5
Description:

	Tests whether or not IsVectorExpression() works on a VectorLinearExpressionTranspose object.
*/
func TestVectorExpression_IsVectorExpression5(t *testing.T) {
	// Constants
	N := 10
	c1 := optim.ZerosVector(N)
	vle1 := optim.NewVectorExpression(c1)

	// Check
	if !optim.IsVectorExpression(vle1.Transpose()) {
		t.Errorf(
			"vle1.T was determined to NOT be a vector expression, but it is (type %T)!",
			vle1,
		)
	}
}

/*
TestVectorExpression_NewVectorExpression1
Description:

	Tests whether or not the NewVectorExpression function returns a vector expression.
*/
func TestVectorExpression_NewVectorExpression1(t *testing.T) {
	// Constants
	N := 10
	c := optim.ZerosVector(N)
	ve1 := optim.NewVectorExpression(c)

	// Check that ve1 is a VectorExpression
	if !optim.IsVectorExpression(ve1) {
		t.Errorf(
			"Expected ve1 to be a vector expression; instead it is %T",
			ve1,
		)
	}

}

/*
TestVectorExpression_ToVectorExpression1
Description:

	Tests whether or not the ToVectorExpression properly handles bad inputs.
*/
func TestVectorExpression_ToVectorExpression1(t *testing.T) {
	// Constants
	b1 := false

	// Check
	_, err := optim.ToVectorExpression(b1)
	if err == nil {
		t.Errorf(
			"There was no error thrown, but there should have been!",
		)
	}

	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"the input interface is of type %T, which is not recognized as a VectorExpression.",
			b1,
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}
}

/*
TestVectorExpression_ToVectorExpression2
Description:

	Tests whether or not the ToVectorExpression properly handles
	VarVectorTranspose.
*/
func TestVectorExpression_ToVectorExpression2(t *testing.T) {
	// Constants
	N := 10
	m := optim.NewModel("ToVectorExpression2")
	vv1 := m.AddVariableVector(N)

	// Check
	ve1, err := optim.ToVectorExpression(vv1.Transpose())
	if err != nil {
		t.Errorf("Unexpected error during conversion: %v", err)
	}

	_, ok := ve1.(optim.VarVectorTranspose)
	if !ok {
		t.Errorf(
			"ve1 was not a VarVectorTranspose as expected; received %T",
			ve1,
		)
	}
}

/*
TestVectorExpression_ToVectorExpression3
Description:

	Tests whether or not the ToVectorExpression properly handles
	LinearVectorExpr.
*/
func TestVectorExpression_ToVectorExpression3(t *testing.T) {
	// Constants
	N := 10
	c1 := optim.ZerosVector(N)
	vle1 := optim.NewVectorExpression(c1)

	// Check
	ve1, err := optim.ToVectorExpression(vle1)
	if err != nil {
		t.Errorf("Unexpected error during conversion: %v", err)
	}

	_, ok := ve1.(optim.VectorLinearExpr)
	if !ok {
		t.Errorf(
			"ve1 was not a VectorLinearExpr as expected; received %T",
			ve1,
		)
	}
}

/*
TestVectorExpression_ToVectorExpression4
Description:

	Tests whether or not the ToVectorExpression properly handles
	LinearVectorExpressionTranspose.
*/
func TestVectorExpression_ToVectorExpression4(t *testing.T) {
	// Constants
	N := 10
	c1 := optim.ZerosVector(N)
	vle1 := optim.NewVectorExpression(c1)

	// Check
	ve1, err := optim.ToVectorExpression(vle1.Transpose())
	if err != nil {
		t.Errorf("Unexpected error during conversion: %v", err)
	}

	_, ok := ve1.(optim.VectorLinearExpressionTranspose)
	if !ok {
		t.Errorf(
			"ve1 was not a VectorLinearExpressionTranspose as expected; received %T",
			ve1,
		)
	}
}

/*
TestVectorExpression_ToVectorExpression5
Description:

	Tests whether or not the ToVectorExpression properly handles
	mat.VecDense.
*/
func TestVectorExpression_ToVectorExpression5(t *testing.T) {
	// Constants
	N := 10
	c1 := optim.ZerosVector(N)

	// Check
	ve1, err := optim.ToVectorExpression(c1)
	if err != nil {
		t.Errorf("Unexpected error during conversion: %v", err)
	}

	_, ok := ve1.(optim.KVector)
	if !ok {
		t.Errorf(
			"ve1 was not a KVector as expected; received %T",
			ve1,
		)
	}
}
