package optim_test

import (
	"fmt"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"gonum.org/v1/gonum/mat"
	"strings"
	"testing"
)

/*
vector_constant_transposed_test.go
Description:
	Tests the new type KVectorTranspose which represents a constant vector.
*/

/*
TestKVectorTranspose_At1
Description:

	This test verifies whether or not a 1 is retrieved when we create a KVectorTranspose
	using OnesVector().
*/
func TestKVectorTranspose_AtVec1(t *testing.T) {
	// Create a KVectorTranspose
	desLength := 4
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	targetIndex := 2

	if vec1.AtVec(targetIndex).(optim.K) != 1.0 {
		t.Errorf("vec1[%v] = %v; expected %v.", targetIndex, vec1.AtVec(targetIndex), 1.0)
	}
}

/*
TestKVectorTranspose_At2
Description:

	This test verifies whether or not an arbitrary number is retrieved when we create a KVectorTranspose
	using NewVecDense().
*/
func TestKVectorTranspose_AtVec2(t *testing.T) {
	// Create a KVectorTranspose
	vec1Elts := []float64{1.0, 3.0, 5.0, 7.0, 9.0}
	var vec1 = optim.KVectorTranspose(*mat.NewVecDense(5, vec1Elts))
	targetIndex := 3

	if vec1.AtVec(targetIndex).(optim.K) != optim.K(vec1Elts[targetIndex]) {
		t.Errorf("vec1[%v] = %v; expected %v.", targetIndex, vec1.AtVec(targetIndex), vec1Elts[targetIndex])
	}
}

/*
TestKVectorTranspose_Len1
Description:

	This function tests that the Len() method works.
	(Should be inherited from the base type mat.DenseVec)
*/
func TestKVectorTranspose_Len1(t *testing.T) {
	// Create a KVectorTranspose
	desLength := 4
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))

	if vec1.Len() != desLength {
		t.Errorf("The length of vec1 should be %v, but instead it is %v.", desLength, vec1.Len())
	}
}

/*
TestKVectorTranspose_Len2
Description:

	This function tests that the Len() method is properly inherited by KVectorTranspose.
*/
func TestKVectorTranspose_Len2(t *testing.T) {
	// Create a KVectorTranspose
	desLength := 10
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))

	if vec1.Len() != desLength {
		t.Errorf("The length of vec1 should be %v, but instead it is %v.", desLength, vec1.Len())
	}
}

/*
TestKVectorTranspose_NumVars1
Description:

	Verify that the number of variables associated with the constant vector is zero.
*/
func TestKVectorTranspose_NumVars1(t *testing.T) {
	// Constant
	kv1 := optim.KVectorTranspose(optim.OnesVector(10))

	// Test
	if kv1.NumVars() != 0 {
		t.Errorf(
			"Expected for there to be zero variables in KVectorTranspose; receive %v",
			kv1.NumVars(),
		)
	}

}

/*
TestKVectorTranspose_IDs1()
Description:

	Verify that the number of variables associated with the constant vector is zero.
*/
func TestKVectorTranspose_IDs1(t *testing.T) {
	// Constant
	kv1 := optim.KVectorTranspose(optim.OnesVector(10))

	// Test
	if len(kv1.IDs()) != 0 {
		t.Errorf(
			"Expected for there to be zero variables in KVectorTranspose; receive %v",
			len(kv1.IDs()),
		)
	}

}

/*
TestKVectorTranspose_LinearCoeff1
Description:

	Verify that the number of variables associated with the constant vector is zero.
*/
func TestKVectorTranspose_LinearCoeff1(t *testing.T) {
	// Constant
	kv1 := optim.KVectorTranspose(optim.OnesVector(10))

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
			"Expected for there to be zero variables in KVectorTranspose; receive %v",
			kv1.NumVars(),
		)
	}

}

/*
TestKVectorTranspose_Constant1
Description:

	Tests that the constant function correctly retrieves a matrix.
*/
func TestKVectorTranspose_Constant1(t *testing.T) {
	// Constant
	kv1 := optim.KVectorTranspose(optim.OnesVector(10))

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
TestKVectorTranspose_Comparison1
Description:

	This function tests that the Comparison() method is properly working for KVectorTranspose inputs.
*/
func TestKVectorTranspose_Comparison1(t *testing.T) {
	// Constants
	desLength := 10
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	var vec2 = optim.KVectorTranspose(optim.ZerosVector(desLength))

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
TestKVectorTranspose_Comparison2
Description:

	This function tests that the Comparison() method is properly working for KVectorTranspose inputs.
	Uses SenseLessThanEqual.
	Comparison of:
	- KVectorTranspose
	- VarVector
*/
func TestKVectorTranspose_Comparison2(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Comparison2")
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	var vec2 = m.AddVariableVector(desLength).Transpose()

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
TestKVectorTranspose_Comparison3
Description:

	This function tests that the Comparison() method is properly working for KVectorTranspose inputs.
	Uses SenseGreaterThanEqual.
	Comparison of:
	- KVectorTranspose
	- VectorLinearExpression
*/
func TestKVectorTranspose_Comparison3(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Comparison3")
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	var vec2 = m.AddVariableVector(desLength)

	L1 := optim.Identity(desLength)
	c1 := optim.OnesVector(desLength)

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
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
TestKVectorTranspose_Comparison4
Description:

	This function tests that the Comparison() method is properly working for KVectorTranspose inputs.
	Input is bad (dimension of linear vector expression is different from constant vector) and error should be thrown.
	Uses SenseGreaterThanEqual.
	Comparison of:
	- KVectorTranspose
	- VectorLinearExpression
*/
func TestKVectorTranspose_Comparison4(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Comparison4")
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
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
TestKVectorTranspose_Comparison5
Description:

	This function tests that the Comparison() method is properly working
	for KVectorTranspose.
	Uses SenseLessThanEqual.
	Comparison of:
	- KVectorTranspose
	- VarVector
	Should throw error
*/
func TestKVectorTranspose_Comparison5(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Comparison5")
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	var vec2 = m.AddVariableVector(desLength)

	// Create Constraint
	_, err := vec1.Comparison(vec2, optim.SenseLessThanEqual)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot compare KVectorTranspose with a normal vector of type %T; Try transposing one or the other!",
			vec2,
		),
	) {
		t.Errorf("There was an unexpected error when creating equality constraint between vec1 and vec2: %v", err)
	}
}

/*
TestKVectorTranspose_Plus1
Description:

	Tests the addition of KVectorTranspose with another KVectorTranspose
*/
func TestKVectorTranspose_Plus1(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel()
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	var vec2 = optim.KVectorTranspose(optim.ZerosVector(desLength))

	// Algorithm
	eOut, err := vec1.Plus(vec2)
	if err != nil {
		t.Errorf("There was an issue adding the two expression.")
	}

	vec3, ok := eOut.(optim.KVectorTranspose)
	if !ok {
		t.Errorf("Expected vec3 to be of type optim.KVectorTranspose; received %T", eOut)
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
TestKVectorTranspose_Plus2
Description:

	Tests the addition of KVectorTranspose with a float64
*/
func TestKVectorTranspose_Plus2(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel()
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	var f1 float64 = 3.1

	// Algorithm
	eOut, err := vec1.Plus(f1)
	if err != nil {
		t.Errorf("There was an issue adding the two expression: %v", err)
	}

	vec3, ok := eOut.(optim.KVectorTranspose)
	if !ok {
		t.Errorf("Expected vec3 to be of type optim.KVectorTranspose; received %T", eOut)
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
TestKVectorTranspose_Plus3
Description:

	Tests the addition of KVectorTranspose with a K
*/
func TestKVectorTranspose_Plus3(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel()
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	var f1 = optim.K(3.1)

	// Algorithm
	eOut, err := vec1.Plus(f1)
	if err != nil {
		t.Errorf("There was an issue adding the two expression: %v", err)
	}

	vec3, ok := eOut.(optim.KVectorTranspose)
	if !ok {
		t.Errorf("Expected vec3 to be of type optim.KVectorTranspose; received %T", eOut)
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
TestKVectorTranspose_Plus4
Description:

	Tests the addition of KVectorTranspose with a mat.VecDense vector
*/
func TestKVectorTranspose_Plus4(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel()
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	var vec2 = optim.KVector(optim.ZerosVector(desLength)).Transpose()

	// Algorithm
	eOut, err := vec1.Plus(vec2)
	if err != nil {
		t.Errorf("There was an issue adding the two expression.")
	}

	vec3, ok := eOut.(optim.KVectorTranspose)
	if !ok {
		t.Errorf("Expected vec3 to be of type optim.KVectorTranspose; received %T", eOut)
	}

	for dimIndex := 0; dimIndex < desLength; dimIndex++ {
		if float64(vec3.AtVec(dimIndex).(optim.K)) != float64(vec1.AtVec(dimIndex).(optim.K))+float64(vec2.AtVec(dimIndex).(optim.K)) {
			t.Errorf(
				"Expected v3.AtVec(%v) = %v; received %v",
				dimIndex,
				vec3.AtVec(dimIndex),
				float64(vec1.AtVec(dimIndex).(optim.K))+float64(vec2.AtVec(dimIndex).(optim.K)),
			)
		}
	}
}

/*
TestKVectorTranspose_Plus5
Description:

	Tests the addition of KVectorTranspose with a mat.VecDense of improper length
*/
func TestKVectorTranspose_Plus5(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel()
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	var vec2 = optim.KVectorTranspose(optim.ZerosVector(desLength - 1))

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
TestKVectorTranspose_Plus6
Description:

	Tests the addition of KVectorTranspose with another KVectorTranspose. Length mismatch.
*/
func TestKVectorTranspose_Plus6(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel()
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	var vec2 = optim.KVectorTranspose(optim.ZerosVector(desLength - 1))

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
TestKVectorTranspose_Plus7
Description:

	Tests the addition of KVectorTranspose with a VarVector
*/
func TestKVectorTranspose_Plus7(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("test-KVectorTranspose-plus7")
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	vec2 := m.AddVariableVector(desLength).Transpose()

	// Algorithm
	eOut, err := vec1.Plus(vec2)
	if err != nil {
		t.Errorf("There was an issue adding the two expression: %v", err)
	}

	vec3, ok := eOut.(optim.VectorLinearExpressionTranspose)
	if !ok {
		t.Errorf("Expected vec3 to be of type optim.KVectorTranspose; received %T", eOut)
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
TestKVectorTranspose_Plus8
Description:

	Tests the addition of KVectorTranspose with a VectorLinearExpr
*/
func TestKVectorTranspose_Plus8(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("test-KVectorTranspose-plus7")
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	vec2 := m.AddVariableVector(desLength).Transpose()

	// Algorithm
	eOut, err := vec2.Plus(vec1.Plus(vec2))
	if err != nil {
		t.Errorf("There was an issue adding the two expression: %v", err)
	}

	vec3, ok := eOut.(optim.VectorLinearExpressionTranspose)
	if !ok {
		t.Errorf("Expected vec3 to be of type optim.KVectorTranspose; received %T", eOut)
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
TestKVectorTranspose_Plus9
Description:

	Tests the addition of KVectorTranspose with a bool
*/
func TestKVectorTranspose_Plus9(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel("test-KVectorTranspose-plus7")
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	b1 := false

	// Algorithm
	_, err := vec1.Plus(b1)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf("Unrecognized expression type %T for addition of KVectorTranspose kvt.Plus(%v)!", b1, b1),
	) {
		t.Errorf("Unexpected error when adding KVectorTranspose with bool! %v", err)
	}
}

/*
TestKVectorTranspose_Plus10
Description:

	Tests the addition of KVectorTranspose with a bool
*/
func TestKVectorTranspose_Plus10(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel("test-KVectorTranspose-plus7")
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	v2 := optim.OnesVector(desLength)

	// Algorithm
	_, err := vec1.Plus(v2)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Can not add KVectorTranspose to normal vector %v (type %T); transpose one or the other!",
			v2, v2,
		),
	) {
		t.Errorf("Unexpected error when adding KVectorTranspose with bool! %v", err)
	}
}

/*
TestKVectorTranspose_Plus11
Description:

	Tests the addition of KVectorTranspose with a bool
*/
func TestKVectorTranspose_Plus11(t *testing.T) {
	// Constants
	desLength := 10
	//m := optim.NewModel("test-KVectorTranspose-plus7")
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	v2 := optim.KVector(optim.OnesVector(desLength))

	// Algorithm
	_, err := vec1.Plus(v2)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Can not add KVectorTranspose to normal vector %v (type %T); transpose one or the other!",
			v2, v2,
		),
	) {
		t.Errorf("Unexpected error when adding KVectorTranspose with bool! %v", err)
	}
}

/*
TestKVectorTranspose_Plus12
Description:

	Tests the addition of KVectorTranspose with a varvector
*/
func TestKVectorTranspose_Plus12(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("test-KVectorTranspose-plus7")
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))
	v2 := m.AddVariableVector(desLength)

	// Algorithm
	_, err := vec1.Plus(v2)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Can not add KVectorTranspose to normal vector %v (type %T); transpose one or the other!",
			v2, v2,
		),
	) {
		t.Errorf("Unexpected error when adding KVectorTranspose with bool! %v", err)
	}
}

/*
TestKVectorTranspose_Plus13
Description:

	Tests the addition of KVectorTranspose with a vle
*/
func TestKVectorTranspose_Plus13(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("test-KVectorTranspose-plus13")
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))

	vle1 := optim.VectorLinearExpr{
		L: optim.Identity(desLength),
		X: m.AddVariableVector(desLength),
		C: optim.ZerosVector(desLength),
	}

	// Algorithm
	_, err := vec1.Plus(vle1)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Can not add KVectorTranspose to normal vector %v (type %T); transpose one or the other!",
			vle1, vle1,
		),
	) {
		t.Errorf("Unexpected error when adding KVectorTranspose with bool! %v", err)
	}
}

/*
TestKVectorTranspose_Plus14
Description:

	Tests the addition of KVectorTranspose with a vletranspose
*/
func TestKVectorTranspose_Plus14(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("test-KVectorTranspose-plus14")
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))

	vle1 := optim.VectorLinearExpr{
		L: optim.Identity(desLength),
		X: m.AddVariableVector(desLength),
		C: optim.ZerosVector(desLength),
	}

	// Algorithm
	sum1, err := vec1.Plus(vle1.Transpose())
	if err != nil {
		t.Errorf("Expected no error; received %v", err)
	}

	if _, ok := sum1.(optim.VectorLinearExpressionTranspose); !ok {
		t.Errorf(
			"Sum of KVectorTranspose and VectorLinearExpressionTranspose is %T; expected optim.VectorLinearExpressionTranspose",
			sum1,
		)
	}
}

/*
TestKVectorTranspose_Mult1
Description:

	Tests that the scalar multiplication function works as expected.
*/
func TestKVectorTranspose_Mult1(t *testing.T) {
	// Constants
	desLength := 10
	var vec1 = optim.KVectorTranspose(optim.OnesVector(desLength))

	// Algorithm
	prod_out, err := vec1.Mult(30.0)
	if err != nil {
		t.Errorf("Unexpected error in multiplication: %v", err)
	}

	// Check elements of product
	prod, ok := prod_out.(optim.KVectorTranspose)
	if !ok {
		t.Errorf("Unable to cast prod to optim.KVectorTranspose")
	}

	for elt_index := 0; elt_index < prod.Len(); elt_index++ {
		if prod.AtVec(elt_index).(optim.K) != 30.0 {
			t.Errorf(
				"Expected all elements of L to ve 30.0; received %v",
				prod.AtVec(elt_index),
			)
		}
	}
}
