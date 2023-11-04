package optim

import (
	"fmt"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"gonum.org/v1/gonum/mat"
	"strings"
	"testing"
)

/*
constant_test.go
Description:
	Test for the functions of the constant class for MatProInterface.go.
*/

/*
TestK_K1
Description:

	Tests the ability to convert a float to a variable of type K
*/
func TestK_K1(t *testing.T) {
	c1 := 2.1
	c2 := optim.K(3.2)

	if c1 != 2.1 {
		t.Errorf("Expected c1 to be equivalent to 2.1, but c1 = %v", c1)
	}

	if c2 != 3.2 {
		t.Errorf("Expected c2 to be equivalent to 3.2, but c2 = %v", c2)
	}
}

/*
TestK_Variables1
Description:

	Tests the method for extracting variables from the constant K.
*/
func TestK_Variables1(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)

	// Checking
	tempVars := c1.Variables()
	if len(tempVars) != 0 {
		t.Errorf("Expected there to be 0 variables in constant c1; received %v", len(tempVars))
	}
}

/*
TestK_NumVars1
Description:

	Tests the method for extracting variables from the constant K.
*/
func TestK_NumVars1(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)

	// Checking
	nVars := c1.NumVars()
	if nVars != 0 {
		t.Errorf(
			"Expected there to be 0 variables in constant c1; received %v",
			nVars,
		)
	}
}

/*
TestK_IDs1
Description:

	Tests the method for extracting IDs of variables in the constant K.
	There should be no such ids available.
*/
func TestK_IDs1(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)

	// Checking
	tempIDs := c1.IDs()
	if tempIDs != nil {
		t.Errorf(
			"Expected there to be no variable IDs in constant c1; received %v",
			tempIDs,
		)
	}
}

/*
TestK_Coeffs1
Description:

	Tests the method for extracting coefficients of the constant K.
*/
func TestK_Coeffs1(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)

	// Checking
	coeffs1 := c1.Coeffs()
	if coeffs1 != nil {
		t.Errorf(
			"Expected there to be no coefficients associated with the constant c1; received %v",
			coeffs1,
		)
	}
}

/*
TestK_Constant1
Description:

	Tests the method for extracting the constant of the constant K.
*/
func TestK_Constant1(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)

	// Checking
	const1 := c1.Constant()
	if const1 != 3.14 {
		t.Errorf(
			"Expected there to be a single constant associated with the constant c1; but c1 != %v",
			const1,
		)
	}
}

/*
TestK_Plus1
Description:

	Tests the addition operator of a constant with another expression.
*/
func TestK_Plus1(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	c2 := optim.K(6.86)

	// Checking
	expr1, err := c1.Plus(c2)
	if err != nil {
		t.Errorf(
			"There was an issue adding the two constants together! %v",
			err,
		)
	}

	expr1AsK := expr1.(optim.K)
	if expr1AsK != 10.0 {
		t.Errorf(
			"The sum of both constants should be 10.0; received %v",
			expr1AsK,
		)
	}
}

/*
TestK_Plus2
Description:

	Tests the addition operator of a constant with a variable.
*/
func TestK_Plus2(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	v1 := optim.Variable{
		ID: 4, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}

	// Checking
	expr1, err := c1.Plus(v1)
	if err != nil {
		t.Errorf(
			"There was an issue adding a constant and a variable together! %v",
			err,
		)
	}

	expr1AsSLE := expr1.(optim.ScalarLinearExpr)
	if expr1AsSLE.Constant() != float64(c1) {
		t.Errorf(
			"The sum should have constant 3.14; received %v",
			expr1AsSLE.Constant(),
		)
	}

	if expr1AsSLE.NumVars() != 1 {
		t.Errorf(
			"The expression should contain 1 variable; received %v",
			expr1AsSLE.NumVars(),
		)
	}

	if expr1AsSLE.Variables()[0].ID != v1.ID {
		t.Errorf(
			"The expression's one variable should have ID 4; received %v",
			expr1AsSLE.Variables()[0].ID,
		)
	}
}

/*
TestK_Plus3
Description:

	Tests the addition operator of a constant with a scalar linear expression.
*/
func TestK_Plus3(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	v1 := optim.Variable{
		ID: 4, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}
	sle1 := optim.ScalarLinearExpr{
		L: *mat.NewVecDense(1, []float64{0.7}),
		X: optim.VarVector{
			Elements: []optim.Variable{v1},
		},
		C: 12.1,
	}

	// Checking
	expr1, err := c1.Plus(sle1)
	if err != nil {
		t.Errorf(
			"There was an issue adding a constant and a variable together! %v",
			err,
		)
	}

	expr1AsSLE := expr1.(optim.ScalarLinearExpr)
	if expr1AsSLE.Constant() != float64(c1)+sle1.C {
		t.Errorf(
			"The sum should have constant %v; received %v",
			float64(c1)+sle1.C,
			expr1AsSLE.Constant(),
		)
	}

	if expr1AsSLE.NumVars() != 1 {
		t.Errorf(
			"The expression should contain 1 variable; received %v",
			expr1AsSLE.NumVars(),
		)
	}

	if expr1AsSLE.Variables()[0].ID != v1.ID {
		t.Errorf(
			"The expression's one variable should have ID 4; received %v",
			expr1AsSLE.Variables()[0].ID,
		)
	}
}

/*
TestK_Plus4
Description:

	Tests the addition operator of a constant with a scalar quadratic expression.
*/
func TestK_Plus4(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	v1 := optim.Variable{
		ID: 4, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}
	v2 := optim.Variable{
		ID: 7, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}
	sqe1 := optim.ScalarQuadraticExpression{
		Q: optim.Identity(2),
		L: *mat.NewVecDense(2, []float64{0.7, 1.7}),
		X: optim.VarVector{
			Elements: []optim.Variable{v1, v2},
		},
		C: 12.3,
	}

	// Checking
	expr1, err := c1.Plus(sqe1)
	if err != nil {
		t.Errorf(
			"There was an issue adding a constant and a variable together! %v",
			err,
		)
	}

	expr1AsSQE := expr1.(optim.ScalarQuadraticExpression)
	if expr1AsSQE.Constant() != float64(c1)+sqe1.C {
		t.Errorf(
			"The sum should have constant %v; received %v",
			float64(c1)+sqe1.C,
			expr1AsSQE.Constant(),
		)
	}

	if expr1AsSQE.NumVars() != 2 {
		t.Errorf(
			"The expression should contain 1 variable; received %v",
			expr1AsSQE.NumVars(),
		)
	}

	if expr1AsSQE.Variables()[0].ID != v1.ID {
		t.Errorf(
			"The expression's one variable should have ID 4; received %v",
			expr1AsSQE.Variables()[0].ID,
		)
	}
	if expr1AsSQE.Variables()[1].ID != v2.ID {
		t.Errorf(
			"The expression's one variable should have ID %v; received %v",
			v2.ID,
			expr1AsSQE.Variables()[1].ID,
		)
	}
}

/*
TestK_Plus5
Description:

	Tests the addition operator of a constant with an error.
*/
func TestK_Plus5(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	e1 := fmt.Errorf(
		"Berleezyyy!",
	)

	// Checking
	_, err := c1.Plus(e1)
	if err == nil {
		t.Errorf(
			"There should be an issue adding a constant and an error together! %v",
			err,
		)
	}

}

/*
TestK_Plus6
Description:

	Tests the addition operator of a constant with a constant but with an optional error included.
*/
func TestK_Plus6(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	c2 := optim.K(6.86)
	e1 := fmt.Errorf(
		"Berleezyyy!",
	)

	// Checking
	_, err := c1.Plus(c2, e1)
	if err != e1 {
		t.Errorf(
			"There should be an issue adding a constant and an error together! %v",
			err,
		)
	}

}

/*
TestK_LessEq1
Description:

	Tests the ability to create constraints using a constant and a variable.
*/
func TestK_LessEq1(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	v1 := optim.Variable{
		ID: 4, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}

	// Test Function
	constr1, err := c1.LessEq(v1)
	if err != nil {
		t.Errorf("There was an issue comparing c1 and v1: %v", err)
	}

	_, ok1 := constr1.LeftHandSide.(optim.K)
	if !ok1 {
		t.Errorf("LHS is expected to be of type optim.K; received %T", constr1.LeftHandSide)
	}

	_, ok2 := constr1.RightHandSide.(optim.Variable)
	if !ok2 {
		t.Errorf("RHS is expected to be of type optim.Variable; received %T", constr1.RightHandSide)
	}
}

/*
TestK_GreaterEq1
Description:

	Tests the ability to create constraints using a constant and a scalar linear expression.
*/
func TestK_GreaterEq1(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	v1 := optim.Variable{
		ID: 4, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}
	sle1 := optim.ScalarLinearExpr{
		L: *mat.NewVecDense(1, []float64{0.7}),
		X: optim.VarVector{
			Elements: []optim.Variable{v1},
		},
		C: 12.1,
	}

	// Test Function
	constr1, err := c1.GreaterEq(sle1)
	if err != nil {
		t.Errorf("There was an issue comparing c1 and v1: %v", err)
	}

	_, ok1 := constr1.LeftHandSide.(optim.K)
	if !ok1 {
		t.Errorf("LHS is expected to be of type optim.K; received %T", constr1.LeftHandSide)
	}

	_, ok2 := constr1.RightHandSide.(optim.ScalarLinearExpr)
	if !ok2 {
		t.Errorf(
			"RHS is expected to be of type optim.ScalarLinearExpr; received %T",
			constr1.RightHandSide,
		)
	}

	if constr1.Sense != optim.SenseGreaterThanEqual {
		t.Errorf(
			"Comparison sense is expected to be optim.SenseGreaterThanEqual; received %v",
			constr1.Sense,
		)
	}
}

/*
TestK_Eq1
Description:

	Tests the ability to create constraints using a constant and a scalar linear expression.
*/
func TestK_Eq1(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	v1 := optim.Variable{
		ID: 4, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}
	sle1 := optim.ScalarLinearExpr{
		L: *mat.NewVecDense(1, []float64{0.7}),
		X: optim.VarVector{
			Elements: []optim.Variable{v1},
		},
		C: 12.1,
	}

	// Test Function
	constr1, err := c1.Eq(sle1)
	if err != nil {
		t.Errorf("There was an issue comparing c1 and v1: %v", err)
	}

	_, ok1 := constr1.LeftHandSide.(optim.K)
	if !ok1 {
		t.Errorf("LHS is expected to be of type optim.K; received %T", constr1.LeftHandSide)
	}

	_, ok2 := constr1.RightHandSide.(optim.ScalarLinearExpr)
	if !ok2 {
		t.Errorf(
			"RHS is expected to be of type optim.ScalarLinearExpr; received %T",
			constr1.RightHandSide,
		)
	}

	if constr1.Sense != optim.SenseEqual {
		t.Errorf(
			"Comparison sense is expected to be optim.SenseEqual; received %v",
			constr1.Sense,
		)
	}
}

/*
TestK_Comparison1
Description:

	Tests the comparison method's error handling properties.
*/
func TestK_Comparison1(t *testing.T) {
	// Constants
	k1 := optim.K(2.3)
	f2 := 2.18
	err0 := fmt.Errorf("Test")

	// Comparison
	_, err := k1.Comparison(f2, optim.SenseGreaterThanEqual, err0)
	if err == nil {
		t.Errorf("No error was thrown, when it should have been")
	} else {
		if !strings.Contains(
			err.Error(),
			err0.Error(),
		) {
			t.Errorf("Unexpected error: %v", err)
		}
	}

}

/*
TestK_Comparison2
Description:

	Tests the comparison method's error handling properties
	with nil error.
*/
func TestK_Comparison2(t *testing.T) {
	// Constants
	k1 := optim.K(2.3)
	f2 := 2.18
	var err0 error = nil

	// Comparison
	_, err := k1.Comparison(f2, optim.SenseGreaterThanEqual, err0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

}

/*
TestK_Comparison3
Description:

	Tests the comparison method's error handling properties.
*/
func TestK_Comparison3(t *testing.T) {
	// Constants
	k1 := optim.K(2.3)
	kv1 := optim.KVector(optim.OnesVector(10))

	// Comparison
	_, err := k1.Comparison(kv1, optim.SenseGreaterThanEqual)
	if err == nil {
		t.Errorf("There were no errors, but there should have been!")
	} else {
		if !strings.Contains(
			err.Error(),
			fmt.Sprintf(
				"the input interface is of type %T, which is not recognized as a ScalarExpression.",
				kv1,
			),
		) {
			t.Errorf("Unexpected error: %v", err)
		}
	}

}

/*
TestK_Multiply1
Description:

	Tests the ability to multiply a constant with another constant.
*/
func TestK_Multiply1(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	c2 := optim.K(6.86)

	// Algorithm
	expr1, err := c1.Multiply(c2)
	if err != nil {
		t.Errorf("There was an issue multiplying two constants: %v", err)
	}

	expr1AsK, ok := expr1.(optim.K)
	if !ok {
		t.Errorf("There was an issue converting the product to a constant!")
	}

	if expr1AsK != 3.14*6.86 {
		t.Errorf(
			"Expected product to have value %v; received %v",
			3.14*6.86,
			expr1AsK,
		)
	}
}

/*
TestK_Multiply2
Description:

	Tests the ability to multiply a constant with a variable.
*/
func TestK_Multiply2(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	v1 := optim.Variable{
		ID: 4, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}

	// Algorithm
	expr1, err := c1.Multiply(v1)
	if err != nil {
		t.Errorf("There was an issue multiplying two constants: %v", err)
	}

	expr1AsSLE, ok := expr1.(optim.ScalarLinearExpr)
	if !ok {
		t.Errorf("There was an issue converting the product to a constant!")
	}

	if expr1AsSLE.C != 0.0 {
		t.Errorf(
			"Expected product's constant to have value %v; received %v",
			0.0,
			expr1AsSLE.C,
		)
	}

	if expr1AsSLE.L.AtVec(0) != float64(c1) {
		t.Errorf(
			"Expected product to have coefficient %v; received %v",
			c1,
			expr1AsSLE.L.AtVec(0),
		)
	}
}

/*
TestK_Multiply3
Description:

	Tests the ability to multiply a constant with a scalar linear expression.
*/
func TestK_Multiply3(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	v1 := optim.Variable{
		ID: 4, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}
	sle1 := optim.ScalarLinearExpr{
		L: *mat.NewVecDense(1, []float64{0.7}),
		X: optim.VarVector{
			Elements: []optim.Variable{v1},
		},
		C: 12.1,
	}

	// Algorithm
	expr1, err := c1.Multiply(sle1)
	if err != nil {
		t.Errorf("There was an issue multiplying two constants: %v", err)
	}

	expr1AsSLE, ok := expr1.(optim.ScalarLinearExpr)
	if !ok {
		t.Errorf("There was an issue converting the product to a constant!")
	}

	if expr1AsSLE.C != 3.14*12.1 {
		t.Errorf(
			"Expected product's constant to have value %v; received %v",
			float64(c1)*sle1.C,
			expr1AsSLE.C,
		)
	}

	if expr1AsSLE.L.AtVec(0) != float64(c1)*sle1.L.AtVec(0) {
		t.Errorf(
			"Expected product to have coefficient %v; received %v",
			float64(c1)*sle1.L.AtVec(0),
			expr1AsSLE.L.AtVec(0),
		)
	}
}

/*
TestK_Multiply4
Description:

	Tests the ability to multiply a constant with a scalar linear expression.
*/
func TestK_Multiply4(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	v1 := optim.Variable{
		ID: 4, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}
	v2 := optim.Variable{
		ID: 5, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}
	sqe1 := optim.ScalarQuadraticExpression{
		Q: optim.Identity(2),
		L: *mat.NewVecDense(2, []float64{0.7, 1.7}),
		X: optim.VarVector{
			Elements: []optim.Variable{v1, v2},
		},
		C: 12.3,
	}

	// Algorithm
	expr1, err := c1.Multiply(sqe1)
	if err != nil {
		t.Errorf("There was an issue multiplying two constants: %v", err)
	}

	expr1AsSQE, ok := expr1.(optim.ScalarQuadraticExpression)
	if !ok {
		t.Errorf("There was an issue converting the product to a constant!")
	}

	if expr1AsSQE.C != float64(c1)*sqe1.C {
		t.Errorf(
			"Expected product's constant to have value %v; received %v",
			float64(c1)*sqe1.C,
			expr1AsSQE.C,
		)
	}

	if expr1AsSQE.L.AtVec(0) != float64(c1)*sqe1.L.AtVec(0) {
		t.Errorf(
			"Expected product to have coefficient %v; received %v",
			float64(c1)*sqe1.L.AtVec(0),
			expr1AsSQE.L.AtVec(0),
		)
	}

	for rowIndex := 0; rowIndex < 2; rowIndex++ {
		for colIndex := 0; colIndex < 2; colIndex++ {
			if expr1AsSQE.Q.At(rowIndex, colIndex) != float64(c1)*sqe1.Q.At(rowIndex, colIndex) {
				t.Errorf(
					"Expected product term to have matrix with Q.At(%v,%v) = %v; received %v",
					rowIndex, colIndex,
					float64(c1)*sqe1.Q.At(rowIndex, colIndex),
					expr1AsSQE.Q.At(rowIndex, colIndex),
				)
			}
		}
	}
}

/*
TestK_Multiply5
Description:

	Tests the ability to multiply a constant with another constant,
	but with a bad error.
*/
func TestK_Multiply5(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	c2 := optim.K(6.86)
	err0 := fmt.Errorf("Test")

	// Algorithm
	_, err := c1.Multiply(c2, err0)
	if err == nil {
		t.Errorf("There was not an error, but there should have been!")
	} else {
		if !strings.Contains(
			err.Error(),
			err0.Error(),
		) {
			t.Errorf("Unexpected error: %v", err)
		}
	}

}

/*
TestK_Multiply6
Description:

	Tests the ability to multiply a constant with a constant vector.
*/
func TestK_Multiply6(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	kv2 := optim.KVector(optim.OnesVector(21))

	// Algorithm
	_, err := c1.Multiply(kv2)
	if err == nil {
		t.Errorf("There was not an error, but there should have been!")
	} else {
		if !strings.Contains(
			err.Error(),
			optim.DimensionError{Operation: "Multiply", Arg1: c1, Arg2: kv2}.Error(),
		) {
			t.Errorf("Unexpected error: %v", err)
		}
	}

}

/*
TestK_Multiply7
Description:

Tests the ability to multiply a constant with a float.
*/
func TestK_Multiply7(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	f2 := 6.86

	// Algorithm
	expr1, err := c1.Multiply(f2)
	if err != nil {
		t.Errorf("There was an issue multiplying two constants: %v", err)
	}

	expr1AsK, ok := expr1.(optim.K)
	if !ok {
		t.Errorf("There was an issue converting the product to a constant!")
	}

	if float64(expr1AsK) != 3.14*f2 {
		t.Errorf(
			"Expected product to have value %v; received %v",
			3.14*6.86,
			expr1AsK,
		)
	}
}

/*
TestK_Multiply8
Description:

Tests the ability to multiply a constant with a float.
*/
func TestK_Multiply8(t *testing.T) {
	// Constants
	c1 := optim.K(3.14)
	kv2 := optim.KVector(optim.OnesVector(1))

	// Algorithm
	expr1, err := c1.Multiply(kv2)
	if err != nil {
		t.Errorf("There was an issue multiplying two constants: %v", err)
	}

	expr1AsKV, ok := expr1.(optim.KVector)
	if !ok {
		t.Errorf("There was an issue converting the product to a constant!")
	}

	if float64(expr1AsKV.AtVec(0).(optim.K)) != 3.14*1.0 {
		t.Errorf(
			"Expected product to have value %v; received %v",
			3.14*1.0,
			expr1AsKV.AtVec(0),
		)
	}
}

/*
TestK_Multiply9
Description:

Tests the ability to multiply a constant with a KVectorTranspose.
*/
func TestK_Multiply9(t *testing.T) {
	// Constants
	N := 3
	c1 := optim.K(3.14)
	vd2 := optim.OnesVector(N)
	vd2.SetVec(1, 2.0)

	kv2 := optim.KVectorTranspose(vd2)

	// Algorithm
	expr1, err := c1.Multiply(kv2)
	if err != nil {
		t.Errorf("There was an issue multiplying two constants: %v", err)
	}

	expr1AsKV, ok := expr1.(optim.KVectorTranspose)
	if !ok {
		t.Errorf("There was an issue converting the product to a constant!")
	}

	expr1AsVD := mat.VecDense(expr1AsKV)
	for eltIndex := 0; eltIndex < expr1AsVD.Len(); eltIndex++ {
		if expr1AsVD.AtVec(eltIndex) != vd2.AtVec(eltIndex)*3.14 {
			t.Errorf(
				"Expected prod[%v] = %v; received %v",
				eltIndex,
				expr1AsVD.AtVec(eltIndex),
				vd2.AtVec(eltIndex)*3.14,
			)
		}
	}
}

/*
TestK_Multiply10
Description:

	Tests the ability to multiply a constant with a VarVector
	of non-unit length.
*/
func TestK_Multiply10(t *testing.T) {
	// Constants
	m := optim.NewModel("TestK_Multiply10")
	N := 3
	c1 := optim.K(3.14)

	kv2 := m.AddVariableVector(N)

	// Algorithm
	_, err := c1.Multiply(kv2)
	if err == nil {
		t.Errorf("no error was thrown, but there should have been!")
	} else {
		if !strings.Contains(
			err.Error(),
			optim.DimensionError{
				Operation: "Multiply",
				Arg1:      c1,
				Arg2:      kv2,
			}.Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestK_Multiply11
Description:

	Tests the ability to multiply a constant with a VarVector
	of unit length.
*/
func TestK_Multiply11(t *testing.T) {
	// Constants
	m := optim.NewModel("TestK_Multiply10")
	N := 1
	c1 := optim.K(3.14)

	kv2 := m.AddVariableVector(N)

	// Algorithm
	prod, err := c1.Multiply(kv2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	prodAsSLE, tf := prod.(optim.ScalarLinearExpr)
	if !tf {
		t.Errorf("expected product to be of type ScalarLinearExpr; received %T", prod)
	}

	if prodAsSLE.C != 0.0 {
		t.Errorf("prod.C = %v =/= 0.0", prodAsSLE.C)
	}
}

/*
TestK_Multiply12
Description:

	Tests the multiplication of a constant with a
	non-unit VarVectorTranspose.
*/
func TestK_Multiply12(t *testing.T) {
	// Constants
	m := optim.NewModel("TestK_Multiply11")
	k1 := optim.K(3.14)
	vv2 := m.AddVariableVector(21)
	vvt2 := vv2.Transpose()

	// Check Multiplication result
	prod, err := k1.Multiply(vvt2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	prodAsVLET, tf := prod.(optim.VectorLinearExpressionTranspose)
	if !tf {
		t.Errorf(
			"Expected product to be of type %v received type %T",
			"VectorLinearExpressionTranspose",
			prod,
		)
	}

	// Check all elements of the matrix
	for rowIndex := 0; rowIndex < vvt2.Len(); rowIndex++ {
		for colIndex := 0; colIndex < vvt2.Len(); colIndex++ {
			if (prodAsVLET.L.At(rowIndex, colIndex) != 1.0*float64(k1)) && (rowIndex == colIndex) {
				t.Errorf(
					" L[%v,%v] = %v =/= %v",
					rowIndex, colIndex,
					prodAsVLET.L.At(rowIndex, colIndex),
					1.0*float64(k1),
				)
			}

			if (prodAsVLET.L.At(rowIndex, colIndex) != 0.0) && (rowIndex != colIndex) {
				t.Errorf(
					"L[%v,%v] = %v =/= 0.0",
					rowIndex, colIndex,
					prodAsVLET.L.At(rowIndex, colIndex),
				)
			}
		}

	}
}

/*
TestK_Multiply13
Description:

	Tests the multiplication of a constant with a
	non-unit VarVectorTranspose.
*/
func TestK_Multiply13(t *testing.T) {
	// Constants
	m := optim.NewModel("TestK_Multiply13")
	k1 := optim.K(3.14)
	vv2 := m.AddVariableVector(1)
	vvt2 := vv2.Transpose()

	// Check Multiplication result
	prod, err := k1.Multiply(vvt2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	prodAsSLE, tf := prod.(optim.ScalarLinearExpr)
	if !tf {
		t.Errorf(
			"Expected product to be of type %v received type %T",
			"ScalarLinearExpression",
			prod,
		)
	}

	// Check all elements of the matrix
	if prodAsSLE.L.AtVec(0) != 1.0*float64(k1) {
		t.Errorf(
			"L[0] = %v =/= %v",
			prodAsSLE.L.AtVec(0),
			1.0*float64(k1),
		)
	}

	if prodAsSLE.C != 0.0 {
		t.Errorf(
			"C = %v =/= 0.0",
			prodAsSLE.C,
		)
	}
}

/*
TestK_Check1
Description:

	Tests that the Check() method returns nil as expected.
*/
func TestK_Check1(t *testing.T) {
	// Constants
	k1 := optim.K(3.14)

	// Algorithm
	if k1.Check() != nil {
		t.Errorf("unexpected error: %v", k1.Check())
	}
}
