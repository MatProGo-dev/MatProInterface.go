package optim_test

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"testing"
)

/*
expression_test.go
Description:
	Tests some of the functions for our Expression interface.
*/

/*
TestExpression_NumVars1()
Description:

	Tests that the expression's NumVars() expression works.
*/
func TestExpression_NumVars1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-expression-numvars1")
	vv1 := m.AddVariableVector(3)
	sle1, err := optim.Sum(vv1, optim.OnesVector(vv1.Len()))
	if err != nil {
		t.Errorf("Unexpected issue in Sum(): %v", err)
	}

	// Tests
	if sle1.NumVars() != vv1.Len() {
		t.Errorf(
			"Expected sle to have %v variables, but found %v",
			vv1.Len(),
			sle1.NumVars(),
		)
	}

	id0 := sle1.IDs()
	if len(id0) != vv1.Len() {
		t.Errorf(
			"%v IDs found when we expected %v.",
			len(id0),
			vv1.Len(),
		)
	}

	for idIndex := 0; idIndex < vv1.Len(); idIndex++ {
		// Compare
		if existentIndex, _ := optim.FindInSlice(vv1.Elements[idIndex].ID, sle1.IDs()); existentIndex == -1 {
			t.Errorf(
				"ID %v was not found in variable vector IDs.",
				vv1.Elements[idIndex].ID,
			)
		}
	}

}

/*
TestExpression_ToExpression1
Description:

	Tests whether or not a variable is correctly detected as an expression.
*/
func TestExpression_ToExpression1(t *testing.T) {
	// Constant
	v := optim.Variable{
		ID: 12, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}

	var e1 interface{} = v
	// Algorithm

	e1AsE, err := optim.ToExpression(e1)
	if err != nil {
		t.Errorf("Unexpected error when evaluating ToExpression(): %v", err)
	}

	_, ok1 := e1AsE.(optim.Variable)
	if !ok1 {
		t.Errorf("Expected for e1 to be of Variable type; received %T", e1AsE)
	}

}

/*
TestExpression_ToExpression2
Description:

	Tests whether or not a bool is correctly detected as an expression.
*/
func TestExpression_ToExpression2(t *testing.T) {
	// Constant
	b1 := false

	var e1 interface{} = b1
	// Algorithm

	_, err := optim.ToExpression(e1)
	if err == nil {
		t.Errorf("Expected error when evaluating ToExpression(); received none.")
	}

}

/*
TestExpression_ToExpression3
Description:

	Tests whether or not a ScalarQuadraticExpression is correctly detected as an expression.
*/
func TestExpression_ToExpression3(t *testing.T) {
	// Constant
	v1 := optim.Variable{
		ID: 12, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}
	v2 := optim.Variable{
		ID: 13, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}
	sqe1 := optim.ScalarQuadraticExpression{
		Q: optim.Identity(2),
		L: optim.OnesVector(2),
		X: optim.VarVector{[]optim.Variable{v1, v2}},
		C: 1.0,
	}

	var e1 interface{} = sqe1
	// Algorithm

	e1AsSQE, err := optim.ToExpression(e1)
	if err != nil {
		t.Errorf("Received unexpected error when evaluating ToExpression(): %v", err)
	}

	_, ok1 := e1AsSQE.(optim.ScalarQuadraticExpression)
	if !ok1 {
		t.Errorf("Expected for e1 to be of Variable type; received %T", e1AsSQE)
	}

}

/*
TestExpression_ToExpression4
Description:

	Tests whether or not a VarVectorTranspose is correctly detected as an expression.
*/
func TestExpression_ToExpression4(t *testing.T) {
	// Constant
	m := optim.NewModel("ToExpression4")
	vv1 := m.AddVariableVector(10)

	var e1 interface{} = vv1.Transpose()
	// Algorithm

	e1AsSQE, err := optim.ToExpression(e1)
	if err != nil {
		t.Errorf("Received unexpected error when evaluating ToExpression(): %v", err)
	}

	_, ok1 := e1AsSQE.(optim.VarVectorTranspose)
	if !ok1 {
		t.Errorf("Expected for e1 to be of Variable type; received %T", e1AsSQE)
	}

}

/*
TestExpression_ToExpression5
Description:

	Tests whether or not a K is correctly detected as an expression.
*/
func TestExpression_ToExpression5(t *testing.T) {
	// Constant
	k1 := optim.K(3.14)

	var e1 interface{} = k1
	// Algorithm

	e1AsSQE, err := optim.ToExpression(e1)
	if err != nil {
		t.Errorf("Received unexpected error when evaluating ToExpression(): %v", err)
	}

	_, ok1 := e1AsSQE.(optim.K)
	if !ok1 {
		t.Errorf("Expected for e1 to be of Variable type; received %T", e1AsSQE)
	}

}

/*
TestExpression_ToExpression6
Description:

	Tests whether or not a mat.VecDense is correctly detected as an expression.
*/
func TestExpression_ToExpression6(t *testing.T) {
	// Constant
	vd1 := optim.OnesVector(7)

	var e1 interface{} = vd1
	// Algorithm

	e1AsSQE, err := optim.ToExpression(e1)
	if err != nil {
		t.Errorf("Received unexpected error when evaluating ToExpression(): %v", err)
	}

	_, ok1 := e1AsSQE.(optim.KVector)
	if !ok1 {
		t.Errorf("Expected for e1 to be of Variable type; received %T", e1AsSQE)
	}

}

/*
TestExpression_ToExpression7
Description:

	Tests whether or not a KVector is correctly detected as an expression.
*/
func TestExpression_ToExpression7(t *testing.T) {
	// Constant
	kv1 := optim.KVector(optim.OnesVector(7))

	var e1 interface{} = kv1
	// Algorithm

	e1AsSQE, err := optim.ToExpression(e1)
	if err != nil {
		t.Errorf("Received unexpected error when evaluating ToExpression(): %v", err)
	}

	_, ok1 := e1AsSQE.(optim.KVector)
	if !ok1 {
		t.Errorf("Expected for e1 to be of Variable type; received %T", e1AsSQE)
	}

}

/*
TestExpression_ToExpression8
Description:

	Tests whether or not a VectorLinearExpr is correctly detected as an expression.
*/
func TestExpression_ToExpression8(t *testing.T) {
	// Constant
	m := optim.NewModel("ToExpression8")
	vv1 := m.AddVariableVector(10)
	sum1, _ := vv1.Plus(optim.OnesVector(vv1.Len()))

	var e1 interface{} = sum1
	// Algorithm

	e1AsSQE, err := optim.ToExpression(e1)
	if err != nil {
		t.Errorf("Received unexpected error when evaluating ToExpression(): %v", err)
	}

	_, ok1 := e1AsSQE.(optim.VectorLinearExpr)
	if !ok1 {
		t.Errorf("Expected for e1 to be of Variable type; received %T", e1AsSQE)
	}

}

/*
TestExpression_ToExpression9
Description:

	Tests whether or not a KVectorTranspose is correctly detected as an expression.
*/
func TestExpression_ToExpression9(t *testing.T) {
	// Constant
	kv1 := optim.KVector(optim.OnesVector(7))

	var e1 interface{} = kv1.Transpose()
	// Algorithm

	e1AsSQE, err := optim.ToExpression(e1)
	if err != nil {
		t.Errorf("Received unexpected error when evaluating ToExpression(): %v", err)
	}

	_, ok1 := e1AsSQE.(optim.KVectorTranspose)
	if !ok1 {
		t.Errorf("Expected for e1 to be of Variable type; received %T", e1AsSQE)
	}

}

/*
TestExpression_ToExpression10
Description:

	Tests whether or not a VectorLinearExpressionTranspose is correctly detected as an expression.
*/
func TestExpression_ToExpression10(t *testing.T) {
	// Constant
	m := optim.NewModel("ToExpression10")
	vv1 := m.AddVariableVector(10).Transpose()
	sum1, _ := vv1.Plus(optim.KVector(optim.OnesVector(vv1.Len())).Transpose())

	var e1 interface{} = sum1
	// Algorithm

	e1AsSQE, err := optim.ToExpression(e1)
	if err != nil {
		t.Errorf("Received unexpected error when evaluating ToExpression(): %v", err)
	}

	_, ok1 := e1AsSQE.(optim.VectorLinearExpressionTranspose)
	if !ok1 {
		t.Errorf("Expected for e1 to be of Variable type; received %T", e1AsSQE)
	}

}

/*
TestExpression_IsExpression1
Description:

	Tests whether or not a ScalarQuadraticExpression is correctly detected as an expression.
*/
func TestExpression_IsExpression1(t *testing.T) {
	// Constant
	v1 := optim.Variable{
		ID: 12, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}
	v2 := optim.Variable{
		ID: 13, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}
	sqe1 := optim.ScalarQuadraticExpression{
		Q: optim.Identity(2),
		L: optim.OnesVector(2),
		X: optim.VarVector{[]optim.Variable{v1, v2}},
		C: 1.0,
	}

	var e1 interface{} = sqe1
	// Algorithm

	ok1 := optim.IsExpression(e1)
	if !ok1 {
		t.Errorf("Expected for e1 to be an Expression; but IsExpression() disagrees!")
	}

}

/*
TestExpression_IsExpression2
Description:

	Tests whether or not a Bool is correctly detected as an expression.
*/
func TestExpression_IsExpression2(t *testing.T) {
	// Constant
	b1 := true

	var e1 interface{} = b1
	// Algorithm

	ok1 := optim.IsExpression(e1)
	if ok1 {
		t.Errorf("Expected for e1 to NOT be an Expression; but IsExpression() disagrees!")
	}

}
