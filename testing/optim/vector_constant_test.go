package optim_test

import (
	"fmt"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"gonum.org/v1/gonum/mat"
	"strings"
	"testing"
)

/*
vector_constant_test.go
Description:
	Tests the new type KVector which represents a constant vector.
*/

/*
TestKVector_At1
Description:

	This test verifies whether or not a 1 is retrieved when we create a KVector
	using OnesVector().
*/
func TestKVector_AtVec1(t *testing.T) {
	// Create a KVector
	desLength := 4
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	targetIndex := 2

	if vec1.AtVec(targetIndex).(optim.K) != 1.0 {
		t.Errorf("vec1[%v] = %v; expected %v.", targetIndex, vec1.AtVec(targetIndex), 1.0)
	}
}

/*
TestKVector_At2
Description:

	This test verifies whether or not an arbitrary number is retrieved when we create a KVector
	using NewVecDense().
*/
func TestKVector_AtVec2(t *testing.T) {
	// Create a KVector
	vec1Elts := []float64{1.0, 3.0, 5.0, 7.0, 9.0}
	var vec1 = optim.KVector(*mat.NewVecDense(5, vec1Elts))
	targetIndex := 3

	if vec1.AtVec(targetIndex).(optim.K) != optim.K(vec1Elts[targetIndex]) {
		t.Errorf("vec1[%v] = %v; expected %v.", targetIndex, vec1.AtVec(targetIndex), vec1Elts[targetIndex])
	}
}

/*
TestKVector_Len1
Description:

	This function tests that the Len() method works.
	(Should be inherited from the base type mat.DenseVec)
*/
func TestKVector_Len1(t *testing.T) {
	// Create a KVector
	desLength := 4
	var vec1 = optim.KVector(optim.OnesVector(desLength))

	if vec1.Len() != desLength {
		t.Errorf("The length of vec1 should be %v, but instead it is %v.", desLength, vec1.Len())
	}
}

/*
TestKVector_Len2
Description:

	This function tests that the Len() method is properly inherited by KVector.
*/
func TestKVector_Len2(t *testing.T) {
	// Create a KVector
	desLength := 10
	var vec1 = optim.KVector(optim.OnesVector(desLength))

	if vec1.Len() != desLength {
		t.Errorf("The length of vec1 should be %v, but instead it is %v.", desLength, vec1.Len())
	}
}

/*
TestKVector_NumVars1
Description:

	Verify that the number of variables associated with the constant vector is zero.
*/
func TestKVector_NumVars1(t *testing.T) {
	// Constant
	kv1 := optim.KVector(optim.OnesVector(10))

	// Test
	if kv1.NumVars() != 0 {
		t.Errorf(
			"Expected for there to be zero variables in KVector; receive %v",
			kv1.NumVars(),
		)
	}

}

/*
TestKVector_IDs1()
Description:

	Verify that the number of variables associated with the constant vector is zero.
*/
func TestKVector_IDs1(t *testing.T) {
	// Constant
	kv1 := optim.KVector(optim.OnesVector(10))

	// Test
	if len(kv1.IDs()) != 0 {
		t.Errorf(
			"Expected for there to be zero variables in KVector; receive %v",
			len(kv1.IDs()),
		)
	}

}

/*
TestKVector_LinearCoeff1
Description:

	Verify that the number of variables associated with the constant vector is zero.
*/
func TestKVector_LinearCoeff1(t *testing.T) {
	// Constant
	kv1 := optim.KVector(optim.OnesVector(10))

	// Test
	coeff := kv1.LinearCoeff()
	nx, ny := coeff.Dims()
	for rowIndex := 0; rowIndex < nx; rowIndex++ {
		for colIndex := 0; colIndex < ny; colIndex++ {
			if coeff.At(rowIndex, colIndex) != 0.0 {
				t.Errorf(
					"Expected coeff[%v,%v] = %v; expected %v",
					rowIndex, colIndex, coeff.At(rowIndex, colIndex),
					0.0,
				)
			}

		}

	}
	if kv1.NumVars() != 0 {
		t.Errorf(
			"Expected for there to be zero variables in KVector; receive %v",
			kv1.NumVars(),
		)
	}

}

/*
TestKVector_Constant1
Description:

	Tests that the constant function correctly retrieves a matrix.
*/
func TestKVector_Constant1(t *testing.T) {
	// Constant
	kv1 := optim.KVector(optim.OnesVector(10))

	// Test
	mat1 := kv1.Constant()

	if mat1.Len() != kv1.Len() {
		t.Errorf(
			"Expected the constant to have length %v; received %v.",
			kv1.Len(), mat1.Len(),
		)
	}

	for i := 0; i < kv1.Len(); i++ {
		if float64(kv1.AtVec(i).(optim.K)) != mat1.AtVec(i) {
			t.Errorf(
				"Expected vector at index %v to be %v; received %v",
				i, kv1.AtVec(i),
				mat1.AtVec(i),
			)
		}
	}
}

/*
TestKVector_Comparison1
Description:

	This function tests that the Comparison() method is properly working for KVector inputs.
*/
func TestKVector_Comparison1(t *testing.T) {
	// Constants
	desLength := 10
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var vec2 = optim.KVector(optim.ZerosVector(desLength))

	// Create Constraint
	constr, err := vec1.Comparison(vec2, optim.SenseEqual)
	if err != nil {
		t.Errorf("There was an issue creating equality constraint between vec1 and vec2: %v", err)
	}

	if constr.LeftHandSide.Len() != vec1.Len() {
		t.Errorf(
			"Expected left hand side (length %v) to have same length as vec1 (length %v).",
			constr.LeftHandSide.Len(),
			vec1.Len(),
		)
	}
}

/*
TestKVector_Comparison2
Description:

	This function tests that the Comparison() method is properly working for KVector inputs.
	Uses SenseLessThanEqual.
	Comparison of:
	- KVector
	- VarVector
*/
func TestKVector_Comparison2(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Comparison2")
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var vec2 = m.AddVariableVector(desLength)

	// Create Constraint
	constr, err := vec1.Comparison(vec2, optim.SenseLessThanEqual)
	if err != nil {
		t.Errorf("There was an issue creating equality constraint between vec1 and vec2: %v", err)
	}

	if constr.LeftHandSide.Len() != vec1.Len() {
		t.Errorf(
			"Expected left hand side (length %v) to have same length as vec1 (length %v).",
			constr.LeftHandSide.Len(),
			vec1.Len(),
		)
	}
}

/*
TestKVector_Comparison3
Description:

	This function tests that the Comparison() method is properly working for KVector inputs.
	Uses SenseGreaterThanEqual.
	Comparison of:
	- KVector
	- VectorLinearExpression
*/
func TestKVector_Comparison3(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Comparison3")
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var vec2 = m.AddVariableVector(desLength)

	L1 := optim.Identity(desLength)
	c1 := optim.OnesVector(desLength)

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vec2, L1, c1,
	}

	// Create Constraint
	constr, err := vec1.Comparison(ve1, optim.SenseGreaterThanEqual)
	if err != nil {
		t.Errorf("There was an issue creating equality constraint between vec1 and vec2: %v", err)
	}

	if constr.LeftHandSide.Len() != vec1.Len() {
		t.Errorf(
			"Expected left hand side (length %v) to have same length as vec1 (length %v).",
			constr.LeftHandSide.Len(),
			vec1.Len(),
		)
	}
}

/*
TestKVector_Comparison4
Description:

	This function tests that the Comparison() method is properly working for KVector inputs.
	Input is bad (dimension of linear vector expression is different from constant vector) and error should be thrown.
	Uses SenseGreaterThanEqual.
	Comparison of:
	- KVector
	- VectorLinearExpression
*/
func TestKVector_Comparison4(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Comparison4")
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var vec2 = m.AddVariableVector(desLength - 1)

	L1 := optim.Identity(desLength - 1)
	c1 := optim.OnesVector(desLength - 1)

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vec2, L1, c1,
	}

	// Create Constraint
	_, err := vec1.Comparison(ve1, optim.SenseGreaterThanEqual)
	if strings.Contains(err.Error(), fmt.Sprintf("The two vector inputs to Eq() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!", vec1.Len(), ve1.Len())) {
		t.Errorf("There was an issue creating equality constraint between vec1 and vec2: %v", err)
	}
}

/*
TestKVector_Comparison5
Description:

	Tests the Eq comparison of KVector with a bool
	results in a proper error.
*/
func TestKVector_Comparison5(t *testing.T) {
	// Constants
	desLength := 10
	var vec1 = optim.KVector(optim.OnesVector(desLength))

	// Algorithm
	_, err := vec1.Comparison(false, optim.SenseEqual)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"The input to KVector's '%v' comparison (%v) has unexpected type: %T",
			optim.SenseEqual, false, false,
		),
	) {
		t.Errorf(
			"The input to KVector's '%v' comparison (%v) has unexpected type: %T",
			optim.SenseEqual, false, false,
		)
		t.Errorf(
			"Unexpected error when comparing kVector with bool! %v",
			err,
		)
	}

}

/*
TestKVector_Plus1
Description:

	Tests the addition of KVector with another KVector
*/
func TestKVector_Plus1(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel()
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var vec2 = optim.KVector(optim.ZerosVector(desLength))

	// Algorithm
	eOut, err := vec1.Plus(vec2)
	if err != nil {
		t.Errorf("There was an issue adding the two expression.")
	}

	vec3, ok := eOut.(optim.KVector)
	if !ok {
		t.Errorf("Expected vec3 to be of type optim.KVector; received %T", eOut)
	}

	for dimIndex := 0; dimIndex < desLength; dimIndex++ {
		if vec3.AtVec(dimIndex) != vec1.AtVec(dimIndex).(optim.K)+vec2.AtVec(dimIndex).(optim.K) {
			t.Errorf(
				"Expected v3.AtVec(%v) = %v; received %v",
				dimIndex,
				vec3.AtVec(dimIndex),
				vec1.AtVec(dimIndex).(optim.K)+vec2.AtVec(dimIndex).(optim.K),
			)
		}
	}
}

/*
TestKVector_Plus2
Description:

	Tests the addition of KVector with a float64
*/
func TestKVector_Plus2(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel()
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var f1 float64 = 3.1

	// Algorithm
	eOut, err := vec1.Plus(f1)
	if err != nil {
		t.Errorf("There was an issue adding the two expression: %v", err)
	}

	vec3, ok := eOut.(optim.KVector)
	if !ok {
		t.Errorf("Expected vec3 to be of type optim.KVector; received %T", eOut)
	}

	for dimIndex := 0; dimIndex < desLength; dimIndex++ {
		if vec3.AtVec(dimIndex) != vec1.AtVec(dimIndex).(optim.K)+optim.K(f1) {
			t.Errorf(
				"Expected v3.AtVec(%v) = %v; received %v",
				dimIndex,
				vec3.AtVec(dimIndex),
				vec1.AtVec(dimIndex).(optim.K)+optim.K(f1),
			)
		}
	}
}

/*
TestKVector_Plus3
Description:

	Tests the addition of KVector with a K
*/
func TestKVector_Plus3(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel()
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var f1 = optim.K(3.1)

	// Algorithm
	eOut, err := vec1.Plus(f1)
	if err != nil {
		t.Errorf("There was an issue adding the two expression: %v", err)
	}

	vec3, ok := eOut.(optim.KVector)
	if !ok {
		t.Errorf("Expected vec3 to be of type optim.KVector; received %T", eOut)
	}

	for dimIndex := 0; dimIndex < desLength; dimIndex++ {
		if vec3.AtVec(dimIndex) != vec1.AtVec(dimIndex).(optim.K)+f1 {
			t.Errorf(
				"Expected v3.AtVec(%v) = %v; received %v",
				dimIndex,
				vec3.AtVec(dimIndex),
				vec1.AtVec(dimIndex).(optim.K)+f1,
			)
		}
	}
}

/*
TestKVector_Plus4
Description:

	Tests the addition of KVector with a mat.VecDense vector
*/
func TestKVector_Plus4(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel()
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var vec2 = optim.ZerosVector(desLength)

	// Algorithm
	eOut, err := vec1.Plus(vec2)
	if err != nil {
		t.Errorf("There was an issue adding the two expression.")
	}

	vec3, ok := eOut.(optim.KVector)
	if !ok {
		t.Errorf("Expected vec3 to be of type optim.KVector; received %T", eOut)
	}

	for dimIndex := 0; dimIndex < desLength; dimIndex++ {
		if float64(vec3.AtVec(dimIndex).(optim.K)) != float64(vec1.AtVec(dimIndex).(optim.K))+vec2.AtVec(dimIndex) {
			t.Errorf(
				"Expected v3.AtVec(%v) = %v; received %v",
				dimIndex,
				vec3.AtVec(dimIndex),
				float64(vec1.AtVec(dimIndex).(optim.K))+vec2.AtVec(dimIndex),
			)
		}
	}
}

/*
TestKVector_Plus5
Description:

	Tests the addition of KVector with a mat.VecDense of improper length
*/
func TestKVector_Plus5(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel()
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var vec2 = optim.ZerosVector(desLength - 1)

	// Algorithm
	_, err := vec1.Plus(vec2)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Length of vectors in sum do not match! Vectors have lengths %v and %v!",
			vec1.Len(), vec2.Len(),
		)) {
		t.Errorf("The wrong error was detected! %v", err)
	}
}

/*
TestKVector_Plus6
Description:

	Tests the addition of KVector with another KVector. Length mismatch.
*/
func TestKVector_Plus6(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel()
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var vec2 = optim.KVector(optim.ZerosVector(desLength - 1))

	// Algorithm
	_, err := vec1.Plus(vec2)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Length of vectors in sum do not match! Vectors have lengths %v and %v!",
			vec1.Len(), vec2.Len(),
		)) {
		t.Errorf("The wrong error was detected! %v", err)
	}

}

/*
TestKVector_Plus7
Description:

	Tests the addition of KVector with a VarVector
*/
func TestKVector_Plus7(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("test-kvector-plus7")
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	vec2 := m.AddVariableVector(desLength)

	// Algorithm
	eOut, err := vec1.Plus(vec2)
	if err != nil {
		t.Errorf("There was an issue adding the two expression: %v", err)
	}

	vec3, ok := eOut.(optim.VectorLinearExpr)
	if !ok {
		t.Errorf("Expected vec3 to be of type optim.KVector; received %T", eOut)
	}

	for dimIndex := 0; dimIndex < desLength; dimIndex++ {
		if vec3.C.AtVec(dimIndex) != vec1.AtVec(dimIndex).Constant() {
			t.Errorf(
				"Expected v3.AtVec(%v) = %v; received %v",
				dimIndex,
				vec3.AtVec(dimIndex),
				vec1.AtVec(dimIndex).(optim.K),
			)
		}
	}

	for dimIndex := 0; dimIndex < vec3.Len(); dimIndex++ {
		if vec3.X.AtVec(dimIndex).IDs()[0] != vec2.AtVec(dimIndex).IDs()[0] {
			t.Errorf(
				"Expected variable at %v to have ID %v; received %v.",
				dimIndex, vec2.AtVec(dimIndex).IDs()[0],
				vec3.X.AtVec(dimIndex).IDs()[0],
			)
		}
	}

	tempIdentity := optim.Identity(vec2.Len())
	for rowIndex := 0; rowIndex < vec2.Len(); rowIndex++ {
		for colIndex := 0; colIndex < vec2.Len(); colIndex++ {
			// Make sure L is identity
			if vec3.L.At(rowIndex, colIndex) != (&tempIdentity).At(rowIndex, colIndex) {
				t.Errorf(
					"Expected L to be identity matrix, but entry at (%v,%v) (%v) does not match identity.",
					rowIndex, colIndex, vec3.L.At(rowIndex, colIndex),
				)
			}
		}
	}
}

/*
TestKVector_Plus8
Description:

	Tests the addition of KVector with a VectorLinearExpr
*/
func TestKVector_Plus8(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("test-kvector-plus7")
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	vec2 := m.AddVariableVector(desLength)

	// Algorithm
	eOut, err := vec2.Plus(vec1.Plus(vec2))
	if err != nil {
		t.Errorf("There was an issue adding the two expression: %v", err)
	}

	vec3, ok := eOut.(optim.VectorLinearExpr)
	if !ok {
		t.Errorf("Expected vec3 to be of type optim.KVector; received %T", eOut)
	}

	for dimIndex := 0; dimIndex < desLength; dimIndex++ {
		if vec3.C.AtVec(dimIndex) != vec1.AtVec(dimIndex).Constant() {
			t.Errorf(
				"Expected v3.AtVec(%v) = %v; received %v",
				dimIndex,
				vec3.AtVec(dimIndex),
				vec1.AtVec(dimIndex).(optim.K),
			)
		}
	}

	for dimIndex := 0; dimIndex < vec3.Len(); dimIndex++ {
		if vec3.X.AtVec(dimIndex).IDs()[0] != vec2.AtVec(dimIndex).IDs()[0] {
			t.Errorf(
				"Expected variable at %v to have ID %v; received %v.",
				dimIndex, vec2.AtVec(dimIndex).IDs()[0],
				vec3.X.AtVec(dimIndex).IDs()[0],
			)
		}
	}

	tempIdentity := optim.Identity(vec2.Len())

	for rowIndex := 0; rowIndex < vec2.Len(); rowIndex++ {
		for colIndex := 0; colIndex < vec2.Len(); colIndex++ {
			// Make sure L is identity
			if vec3.L.At(rowIndex, colIndex) != 2.0*(&tempIdentity).At(rowIndex, colIndex) {
				t.Errorf(
					"Expected L to be identity matrix, but entry at (%v,%v) (%v) does not match identity.",
					rowIndex, colIndex, vec3.L.At(rowIndex, colIndex),
				)
			}
		}
	}
}

/*
TestKVector_Plus9
Description:

	Tests the addition of KVector with a bool
*/
func TestKVector_Plus9(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel("test-kvector-plus7")
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	b1 := false

	// Algorithm
	_, err := vec1.Plus(b1)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf("Unrecognized expression type %T for addition of KVector kv.Plus(%v)!", b1, b1),
	) {
		t.Errorf("Unexpected error when adding kvector with bool! %v", err)
	}
}

/*
TestKVector_Plus10
Description:

	Tests the addition of KVector with a KVectorTranspose
*/
func TestKVector_Plus10(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel("test-kvector-plus7")
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	vec2 := optim.KVectorTranspose(optim.OnesVector(desLength))

	// Algorithm
	_, err := vec1.Plus(vec2)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot add KVector with a transposed vector %T; Try transposing one or the other!",
			vec2,
		),
	) {
		t.Errorf("Unexpected error when adding kvector with bool! %v", err)
	}
}

/*
TestKVector_Plus11
Description:

	Tests the addition of KVector with a VarVectorTranspose
*/
func TestKVector_Plus11(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("test-kvector-plus11")
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	X1 := m.AddVariableVector(desLength)

	// Algorithm
	_, err := vec1.Plus(X1.Transpose())
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot add KVector with a transposed vector %T; Try transposing one or the other!",
			X1.Transpose(),
		),
	) {
		t.Errorf("Unexpected error when adding kvector with bool! %v", err)
	}
}

/*
TestKVector_Plus12
Description:

	Tests the addition of KVector with a VectorLinearExpression
*/
func TestKVector_Plus12(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("test-kvector-plus12")
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	X1 := m.AddVariableVector(desLength)
	A1 := optim.Identity(X1.Len())
	A1.Set(0, 0, 2.0)
	C1 := optim.ZerosVector(desLength)
	C1.SetVec(2, 13.0)

	vle1 := optim.VectorLinearExpr{
		X1, A1, C1,
	}

	// Algorithm
	sum, err := vec1.Plus(vle1)
	if err != nil {
		t.Errorf(
			"There was an issue computing the sum: %v",
			err,
		)
	}

	sumAsVLE, successful := sum.(optim.VectorLinearExpr)
	if !successful {
		t.Errorf(
			"The received sum is not a VectorLinearExpression; it's %T",
			sum,
		)
	}

	if C1.AtVec(2)+1 != sumAsVLE.C.AtVec(2) {
		t.Errorf(
			"sum.C[%v] = %v =/= %v",
			2, sumAsVLE.C.AtVec(2), C1.AtVec(2)+1.0,
		)
	}

	lhs0 := vec1.AtVec(0).(optim.K)
	if float64(lhs0) != sumAsVLE.C.AtVec(0) {
		t.Errorf(
			"sum.C[%v] = %v =/= %v",
			0, sumAsVLE.C.AtVec(0), vec1.AtVec(0),
		)
	}

}

/*
TestKVector_Plus13
Description:

	Tests the addition of KVector with a VectorLinearExpressionTranspose
*/
func TestKVector_Plus13(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("test-kvector-plus13")
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	X1 := m.AddVariableVector(desLength)
	A1 := optim.Identity(X1.Len())
	A1.Set(0, 0, 2.0)
	C1 := optim.ZerosVector(desLength)
	C1.SetVec(2, 13.0)

	vlet1 := optim.VectorLinearExpressionTranspose{
		X1, A1, C1,
	}

	// Algorithm
	_, err := vec1.Plus(vlet1)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot add KVector with a transposed vector %T; Try transposing one or the other!",
			vlet1,
		),
	) {
		t.Errorf(
			"Unexpected error when adding kvector with VectorLinearExpressionTranspose! %v",
			err,
		)
	}

}

/*
TestKVector_Multiply1
Description:

	Tests the multiplication of the KVector with a float.
*/
func TestKVector_Multiply1(t *testing.T) {
	// Constants
	desLength := 10
	var vec1 = optim.KVector(optim.OnesVector(desLength))

	// Algorithm
	_, err := vec1.Multiply(3.14)
	if !strings.Contains(
		err.Error(),
		"The Multiply() method for KVector has not been implemented yet!",
	) {
		t.Errorf(
			"There was an unexpected error performing multiplication: %v",
			err,
		)
	}
}

/*
TestKVector_LessEq1
Description:

	Tests the LessEq comparison of KVector with a VectorLinearExpressionTranspose.
	(should result in an error)
*/
func TestKVector_LessEq1(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("test-kvector-less-eq1")
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	X1 := m.AddVariableVector(desLength)
	A1 := optim.Identity(X1.Len())
	A1.Set(0, 0, 2.0)
	C1 := optim.ZerosVector(desLength)
	C1.SetVec(2, 13.0)

	vlet1 := optim.VectorLinearExpressionTranspose{
		X1, A1, C1,
	}

	// Algorithm
	_, err := vec1.LessEq(vlet1)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot compare KVector with a transposed vector %T; Try transposing one or the other!",
			vlet1,
		),
	) {
		t.Errorf(
			"Unexpected error when comparing kvector with VectorLinearExpressionTranspose! %v",
			err,
		)
	}

}

/*
TestKVector_GreaterEq1
Description:

	Tests the LessEq comparison of KVector with a VectorLinearExpression
	with different variable lengths.
	(should result in an error)
*/
func TestKVector_GreaterEq1(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("test-kvector-greater-eq1")
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	X1 := m.AddVariableVector(desLength - 1)
	A1 := optim.Identity(X1.Len())
	A1.Set(0, 0, 2.0)
	C1 := optim.ZerosVector(X1.Len())
	C1.SetVec(2, 13.0)

	vle1 := optim.VectorLinearExpr{
		X1, A1, C1,
	}

	// Algorithm
	_, err := vec1.GreaterEq(vle1)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"The two vector inputs to Comparison() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
			vle1.Len(),
			vec1.Len(),
		),
	) {
		t.Errorf(
			"Unexpected error when comparing kvector with VectorLinearExpr! %v",
			err,
		)
	}

}

/*
TestKVector_Eq1
Description:

	Tests the Eq comparison of KVector with a KVectorTranspose
	results in a proper error.
*/
func TestKVector_Eq1(t *testing.T) {
	// Constants
	desLength := 10
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	C1 := optim.ZerosVector(desLength)
	C1.SetVec(2, 13.0)

	// Algorithm
	_, err := vec1.GreaterEq(optim.KVector(C1).Transpose())
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot compare KVector with a transposed vector %T; Try transposing one or the other!",
			optim.KVector(C1).Transpose(),
		),
	) {
		t.Errorf(
			"Unexpected error when comparing kvector with KVectorTranspose! %v",
			err,
		)
	}

}

/*
TestKVector_Transpose1
Description:

	Tests the transposition of a given vector returns the new type KVectorTranspose.
*/
func TestKVector_Transpose1(t *testing.T) {
	// Constants
	desLength := 10
	var vec1 = optim.KVector(optim.OnesVector(desLength))

	// Algorithm
	vecTransposed := vec1.Transpose()

	if _, ok1 := vecTransposed.(optim.KVectorTranspose); !ok1 {
		t.Errorf(
			"Expected transposed KVector to be of type KVectorTranspose; received %Te instead.",
			vecTransposed,
		)
	}
}
