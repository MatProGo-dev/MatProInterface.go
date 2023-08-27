package optim_test

import (
	"fmt"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"strings"
	"testing"
)

func TestVarVectorTranspose_Length1(t *testing.T) {
	m := optim.NewModel("Length1")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVectorTranspose{
		Elements: []optim.Variable{x, y},
	}

	if vv1.Length() != 2 {
		t.Errorf("The length of vv1 was %v; expected %v", vv1.Length(), 2)
	}

}

/*
TestVarVectorTranspose_Length2
Description:

	Tests that a larger vector variable (contains 5 elements) properly returns the right length.
*/
func TestVarVectorTranspose_Length2(t *testing.T) {
	m := optim.NewModel("Length2")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVectorTranspose{
		Elements: []optim.Variable{x, y, x, y, x},
	}

	if vv1.Length() != 5 {
		t.Errorf("The length of vv1 was %v; expected %v", vv1.Length(), 5)
	}

}

/*
TestVarVectorTranspose_NumVars1
Description:

	Tests that the NumVars() method is working properly for a length 10 VarVectorTranspose.
*/
func TestVarVectorTranspose_NumVars1(t *testing.T) {
	m := optim.NewModel("NumVars1")
	vv1 := m.AddVariableVector(10)

	// Create Vector Variable
	vv1T := vv1.Transpose()

	if vv1T.NumVars() != 10 {
		t.Errorf("The length of vv1 was %v; expected %v", vv1.Length(), 5)
	}
}

/*
TestVarVectorTranspose_LinearCoeff1
Description:

	Tests that the LinearCoeff() method is working properly for a length 10 VarVectorTranspose.
	For a transposed vector, the linear coefficient L is the coefficient on the right of the variable.
	(i.e., x^T L^T + c^T)
*/
func TestVarVectorTranspose_LinearCoeff1(t *testing.T) {
	m := optim.NewModel("LinearCoeff1")
	vv1 := m.AddVariableVector(10)

	// Create Vector Variable
	vv1T := vv1.Transpose()

	L1 := vv1T.LinearCoeff()

	nL, mL := L1.Dims()
	if nL != 10 {
		t.Errorf("Linear coefficient has %v rows; expected 10!", nL)
	}
	if mL != 10 {
		t.Errorf("Linear coefficient has %v cols; expected 10!", mL)
	}

	for rowIndex := 0; rowIndex < nL; rowIndex++ {
		for colIndex := 0; colIndex < mL; colIndex++ {
			// Get elt and compare with 0 or 1.
			if (rowIndex == colIndex) && (L1.At(rowIndex, colIndex) != 1.0) {
				t.Errorf(
					"The diagonal element at (%v,%v) should be 1.0; received %v",
					rowIndex, colIndex,
					L1.At(rowIndex, colIndex),
				)
			}
			if (rowIndex != colIndex) && (L1.At(rowIndex, colIndex) != 0.0) {
				t.Errorf(
					"The diagonal element at (%v,%v) should be 0.0; received %v",
					rowIndex, colIndex,
					L1.At(rowIndex, colIndex),
				)
			}
		}

	}
}

/*
TestVarVectorTranspose_At1
Description:

	Tests whether or not we can properly retrieve an element from a given vector.
*/
func TestVarVectorTranspose_At1(t *testing.T) {
	m := optim.NewModel("At1")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVectorTranspose{
		Elements: []optim.Variable{x, y},
	}

	extractedV := vv1.AtVec(1)
	if extractedV.(optim.Variable) != y {
		t.Errorf("Expected for extracted variable, %v, to be the same as %v. They were different!", extractedV, y)
	}
}

/*
TestVarVectorTranspose_At2
Description:

	Tests whether or not we can properly retrieve an element from a given vector.
	Makes sure that if we change the extracted vector, it does not effect the element saved in the slice.
*/
func TestVarVectorTranspose_At2(t *testing.T) {
	m := optim.NewModel("At2")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVectorTranspose{
		Elements: []optim.Variable{x, y},
	}

	extractedV := vv1.AtVec(1).(optim.Variable)
	extractedV.ID = 100

	if extractedV == y {
		t.Errorf("Expected for extracted variable, %v, to be DIFFERENT from %v. They were the same!", extractedV, y)
	}
}

/*
TestVarVectorTranspose_VariableIDs1
Description:

	This test will check to see if 2 unique ids in a VariableVector object will be returned correctly when
	the VariableIDs method is called.
*/
func TestVarVectorTranspose_VariableIDs1(t *testing.T) {
	m := optim.NewModel("VariableIDs1")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVectorTranspose{
		Elements: []optim.Variable{x, y},
	}

	extractedIDs := vv1.IDs()
	// Check to see that x and y have ids in extractedIDs
	if foundIndex, _ := optim.FindInSlice(x.ID, extractedIDs); foundIndex == -1 {
		t.Errorf("x was not found in vv1, but it should have been!")
	}

	if foundIndex, _ := optim.FindInSlice(y.ID, extractedIDs); foundIndex == -1 {
		t.Errorf("x was not found in vv1, but it should have been!")
	}
}

/*
TestVarVectorTranspose_VariableIDs2
Description:

	This test will check to see if a single unique id in a large VariableVector object will be returned correctly when
	the VariableIDs method is called.
*/
func TestVarVectorTranspose_VariableIDs2(t *testing.T) {
	m := optim.NewModel("VariableIDs2")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVectorTranspose{
		Elements: []optim.Variable{x, x, x},
	}

	extractedIDs := vv1.IDs()
	// Check to see that only x has ids in extractedIDs
	if len(extractedIDs) != 1 {
		t.Errorf("There is only one unique variable ID and yet %v IDs were returned by the IDs() method.", len(extractedIDs))
	}

	if foundIndex, _ := optim.FindInSlice(x.ID, extractedIDs); foundIndex == -1 {
		t.Errorf("x was not found in vv1, but it should have been!")
	}

	if foundIndex, _ := optim.FindInSlice(y.ID, extractedIDs); foundIndex != -1 {
		t.Errorf("x was not found in vv1, but it should have been!")
	}
}

/*
TestVarVectorTranspose_Constant1
Description:

	This test verifies that the constant method returns an all zero vector for any VarVectorTranspose object.
*/
func TestVarVectorTranspose_Constant1(t *testing.T) {
	m := optim.NewModel("Constant1")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVectorTranspose{
		Elements: []optim.Variable{x, y},
	}

	extractedConstant := vv1.Constant()

	// Check to see that all elts in the vector are zero.
	for eltIndex := 0; eltIndex < vv1.Len(); eltIndex++ {
		constElt := extractedConstant.AtVec(eltIndex)
		if constElt != 0.0 {
			t.Errorf("Constant vector at index %v is %v; not 0.", eltIndex, constElt)
		}
	}
}

/*
TestVarVectorTranspose_Constant2
Description:

	This test verifies that the constant method returns an all zero vector for any VarVectorTranspose object.
	This one will be extremely long.
*/
func TestVarVectorTranspose_Constant2(t *testing.T) {
	m := optim.NewModel("Constant2")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVectorTranspose{
		Elements: []optim.Variable{x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y},
	}

	extractedConstant := vv1.Constant()

	// Check to see that all elts in the vector are zero.
	for eltIndex := 0; eltIndex < vv1.Len(); eltIndex++ {
		constElt := extractedConstant.AtVec(eltIndex)
		if constElt != 0.0 {
			t.Errorf("Constant vector at index %v is %v; not 0.", eltIndex, constElt)
		}
	}
}

/*
TestVarVectorTranspose_Eq1
Description:

	This test verifies that the Eq method works between a VarVectorTranspose and another object.
*/
func TestVarVectorTranspose_Eq1(t *testing.T) {
	m := optim.NewModel("Eq1")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVectorTranspose{
		Elements: []optim.Variable{x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y},
	}

	zerosAsVecDense := optim.ZerosVector(vv1.Len())
	zerosAsKVector := optim.KVector(zerosAsVecDense).Transpose()

	// Verify that constraint can be created with no issues.
	_, err := vv1.Eq(zerosAsKVector)
	if err != nil {
		t.Errorf("There was an issue creating an equality constraint between vv1 and the zero vector.")
	}
}

/*
TestVarVectorTranspose_Eq2
Description:

	This test verifies that the Eq method works between a VarVectorTranspose and another object.
	Comparison should be between var vector and an unsupported type.
*/
func TestVarVectorTranspose_Eq2(t *testing.T) {
	m := optim.NewModel("Eq2")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVectorTranspose{
		Elements: []optim.Variable{x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y},
	}

	badRHS := false

	// Verify that constraint can be created with no issues.
	_, err := vv1.Eq(badRHS)
	expectedError := fmt.Sprintf("The Eq() method for VarVectorTranspose is not implemented yet for type %T!", badRHS)
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Expected error \"%v\"; received \"%v\"", expectedError, err)
	}
}

/*
TestVarVectorTranspose_Eq3
Description:

	This test verifies that the Eq method works between a VarVectorTranspose and another var vector.
*/
func TestVarVectorTranspose_Eq3(t *testing.T) {
	m := optim.NewModel("Eq3")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVectorTranspose{
		Elements: []optim.Variable{x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y},
	}

	vv2 := optim.VarVectorTranspose{
		Elements: []optim.Variable{y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x},
	}

	// Verify that constraint can be created with no issues.
	_, err := vv1.Eq(vv2)
	if err != nil {
		t.Errorf("There was an error creating equality constraint between the two VarVectorTransposes: %v", err)
	}
}

/*
TestVarVectorTranspose_Eq4
Description:

	This test verifies that the Eq method does not work between a
	VarVectorTranspose object and a normal KVector.
*/
func TestVarVectorTranspose_Eq4(t *testing.T) {
	// Constants
	m := optim.NewModel("Eq1")
	vv1 := m.AddVariableVector(10).Transpose()

	zerosAsVecDense := optim.ZerosVector(vv1.Len())
	zerosAsKVector := optim.KVector(zerosAsVecDense)

	// Verify that constraint can be created with no issues.
	_, err := vv1.Eq(zerosAsKVector)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot commpare VarVectorTranspose with a normal vector %v (%T); Try transposing one or the other!",
			zerosAsKVector, zerosAsKVector,
		),
	) {
		t.Errorf("There was an unexpected error when attempting to run Eq: %v", err)
	}
}

/*
TestVarVectorTranspose_Comparison1
Description:

	Tests how well the comparison function works with a VectorLinearExpression comparison.
*/
func TestVarVectorTranspose_Comparison1(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Comparison1")
	var vec1 = m.AddVariableVector(desLength)
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
TestVarVectorTranspose_Comparison2
Description:

	Tests how well the comparison function works with a VectorLinearExpression comparison.
	Valid comparison of
*/
func TestVarVectorTranspose_Comparison2(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Comparison2")
	var vec1 = m.AddVariableVector(desLength)
	var vec2 = m.AddVariableVector(desLength)

	L1 := optim.Identity(desLength)
	c1 := optim.OnesVector(desLength)

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vec2, L1, c1,
	}

	// Create Constraint
	_, err := vec1.Comparison(ve1, optim.SenseGreaterThanEqual)
	if err != nil {
		t.Errorf("There was an error computing a comparison for operator >=: %v", err)
	}
}

/*
TestVarVectorTranspose_Plus1
Description:

	Testing the Plus operator between a VarVectorTranspose and a KVectorTranspose. Proper sizes were given.
*/
func TestVarVectorTranspose_Plus1(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus1")
	vec1 := m.AddVariableVector(desLength)
	k2 := optim.KVector(optim.OnesVector(desLength))

	// Algorithm
	sum3, err := vec1.Transpose().Plus(k2.Transpose())
	if err != nil {
		t.Errorf("There was an error computing addition: %v", err)
	}

	sum3AsVLE, ok := sum3.(optim.VectorLinearExpressionTranspose)
	if !ok {
		t.Errorf(
			"There was an issue converting sum3 (type %T) to type optim.VectorLinearExpr.",
			sum3,
		)
	}

	// Check values of the variable vector
	for vecIndex := 0; vecIndex < vec1.Len(); vecIndex++ {
		// Check that values of sum3AsVLE and vec1 match
		if sum3AsVLE.X.AtVec(vecIndex).(optim.Variable).ID != vec1.AtVec(vecIndex).(optim.Variable).ID {
			t.Errorf(
				"Expected the value at index in sum3.X[%v] (%v) to be the same as vec1[%v] (%v).",
				vecIndex,
				sum3AsVLE.X.AtVec(vecIndex),
				vecIndex,
				vec1.AtVec(vecIndex),
			)
		}
	}

	// Check the values of the constant vector
	for vecIndex := 0; vecIndex < sum3AsVLE.Len(); vecIndex++ {
		// Check that values of sum3AsVLE and vec1 match
		if sum3AsVLE.C.AtVec(vecIndex) != float64(k2.AtVec(vecIndex).(optim.K)) {
			t.Errorf(
				"Expected the value at index in sum3.X[%v] (%v) to be the same as k2[%v] (%v).",
				vecIndex,
				sum3AsVLE.X.AtVec(vecIndex),
				vecIndex,
				k2.AtVec(vecIndex),
			)
		}
	}
}

/*
TestVarVectorTranspose_Plus2
Description:

	Testing the Plus operator between a VarVectorTranspose and a KVector. Incorrect sizes were given.
*/
func TestVarVectorTranspose_Plus2(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus2")
	vec1 := m.AddVariableVector(desLength)
	k2 := optim.KVector(optim.OnesVector(desLength - 1))

	// Algorithm
	_, err := vec1.Plus(k2)
	if err == nil {
		t.Errorf("No error detected in bad vector addition!")
	}

	if !strings.Contains(err.Error(), "The lengths of two vectors in Plus must match!") {
		t.Errorf("There was an unexpected error computing addition: %v", err)
	}

}

/*
TestVarVectorTranspose_Plus3
Description:

	Testing the Plus operator between a VarVectorTranspose and a KVector. Proper sizes were given.
*/
func TestVarVectorTranspose_Plus3(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus3")
	vec1 := m.AddVariableVector(desLength).Transpose()
	k2 := optim.KVectorTranspose(optim.OnesVector(desLength))

	// Algorithm
	sum3, err := vec1.Plus(k2)
	if err != nil {
		t.Errorf("There was an error computing addition: %v", err)
	}

	sum3AsVLE, ok := sum3.(optim.VectorLinearExpressionTranspose)
	if !ok {
		t.Errorf(
			"There was an issue converting sum3 (type %T) to type optim.VectorLinearExpr.",
			sum3,
		)
	}

	// Check values of the vector
	for vecIndex := 0; vecIndex < vec1.Len(); vecIndex++ {
		// Check that values of sum3AsVLE and vec1 match
		if sum3AsVLE.X.AtVec(vecIndex).(optim.Variable).ID != vec1.AtVec(vecIndex).(optim.Variable).ID {
			t.Errorf(
				"Expected the value at index in sum3.X[%v] (%v) to be the same as vec1[%v] (%v).",
				vecIndex,
				sum3AsVLE.X.AtVec(vecIndex),
				vecIndex,
				vec1.AtVec(vecIndex),
			)
		}
	}

	// Check the values of the constant vector
	for vecIndex := 0; vecIndex < sum3AsVLE.Len(); vecIndex++ {
		// Check that values of sum3AsVLE and vec1 match
		if sum3AsVLE.C.AtVec(vecIndex) != float64(k2.AtVec(vecIndex).(optim.K)) {
			t.Errorf(
				"Expected the value at index in sum3.X[%v] (%v) to be the same as k2[%v] (%v).",
				vecIndex,
				sum3AsVLE.X.AtVec(vecIndex),
				vecIndex,
				k2.AtVec(vecIndex),
			)
		}
	}
}

/*
TestVarVectorTranspose_Plus4
Description:

	Testing the Plus operator between a VarVectorTranspose and a VarVectorTranspose. All vectors are of same size. Some overlap in the variables but not all.
*/
func TestVarVectorTranspose_Plus4(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus4")
	vec1 := m.AddVariableVector(desLength).Transpose()
	vec2 := m.AddVariableVector(desLength - 2).Transpose().(optim.VarVectorTranspose)
	vec3 := optim.VarVectorTranspose{
		append(vec2.Elements, vec1.AtVec(0).(optim.Variable), vec1.AtVec(1).(optim.Variable)),
	}

	// Algorithm
	sum3, err := vec1.Plus(vec3)
	if err != nil {
		t.Errorf("There was an error computing addition: %v", err)
	}

	sum3AsVLE, ok := sum3.(optim.VectorLinearExpressionTranspose)
	if !ok {
		t.Errorf(
			"There was an issue converting sum3 (type %T) to type optim.VectorLinearExpr.",
			sum3,
		)
	}

	// Check values of the vector of variables
	for vecIndex := 0; vecIndex < vec3.Len(); vecIndex++ {
		// Check that values of sum3AsVLE and vec1 match
		if sum3AsVLE.X.AtVec(vecIndex).(optim.Variable).ID != vec3.AtVec(vecIndex).(optim.Variable).ID {
			t.Errorf(
				"Expected the value at index in sum3.X[%v] (%v) to be the same as vec3[%v] (%v).",
				vecIndex,
				sum3AsVLE.X.AtVec(vecIndex),
				vecIndex,
				vec1.AtVec(vecIndex),
			)
		}
	}

	// Check values of the matrix multiplier.
	for rowIndex := 0; rowIndex < desLength; rowIndex++ {
		// Get elements as needed.
		vec1Atr := vec1.AtVec(rowIndex)
		vec3Atr := vec3.AtVec(rowIndex)

		vec1AtRIndex, _ := optim.FindInSlice(vec1Atr, sum3AsVLE.X.Elements)
		vec3AtRIndex, _ := optim.FindInSlice(vec3Atr, sum3AsVLE.X.Elements)

		// Iterate through all columns (all variables)
		for colIndex := 0; colIndex < sum3AsVLE.X.Len(); colIndex++ {

			switch {
			case (colIndex == vec1AtRIndex) && (vec1AtRIndex == vec3AtRIndex):
				if sum3AsVLE.L.At(rowIndex, colIndex) != 2.0 {
					t.Errorf(
						"Expected L(%v,%v) to be 2.0; received %v",
						rowIndex, colIndex,
						sum3AsVLE.L.At(rowIndex, colIndex),
					)
				}
			case (colIndex == vec1AtRIndex) && (vec1AtRIndex != vec3AtRIndex):
				if sum3AsVLE.L.At(rowIndex, colIndex) != 1.0 {
					t.Errorf(
						"Expected L(%v,%v) to be 1.0; received %v",
						rowIndex, colIndex,
						sum3AsVLE.L.At(rowIndex, colIndex),
					)
				}
			case (colIndex == vec3AtRIndex) && (vec1AtRIndex != vec3AtRIndex):
				if sum3AsVLE.L.At(rowIndex, colIndex) != 1.0 {
					t.Errorf(
						"Expected L(%v,%v) to be 1.0; received %v",
						rowIndex, colIndex,
						sum3AsVLE.L.At(rowIndex, colIndex),
					)
				}
			default:
				// All other elements should be 0.0
				if sum3AsVLE.L.At(rowIndex, colIndex) != 0.0 {
					t.Errorf(
						"Expected L(%v,%v) to be 0.0; received %v",
						rowIndex, colIndex,
						sum3AsVLE.L.At(rowIndex, colIndex),
					)
				}
			}

		}
	}

	// Check offset vector (should be zeros)
	for vecIndex := vec1.Len(); vecIndex < vec1.Len()+vec3.Len()-2; vecIndex++ {
		// Check that values of sum3AsVLE.X matches vec2 at the appropriate indices.
		if sum3AsVLE.X.AtVec(vecIndex) != vec1.AtVec(vecIndex-vec3.Len()+2) {
			t.Errorf(
				"Expected the value at index in sum3.X[%v] (%v) to be the same as vec1[%v] (%v).",
				vecIndex,
				sum3AsVLE.X.AtVec(vecIndex).(optim.Variable).ID,
				vecIndex-vec1.Len(),
				vec1.AtVec(vecIndex-vec3.Len()),
			)
		}
	}

	// Check the values of the constant vector
	for vecIndex := 0; vecIndex < sum3AsVLE.Len(); vecIndex++ {
		// Check that values of sum3AsVLE and vec1 match
		if sum3AsVLE.C.AtVec(vecIndex) != 0.0 {
			t.Errorf(
				"Expected the value of constant to be zero vector, but sum3.C[%v] = %v!",
				vecIndex,
				sum3AsVLE.C.AtVec(vecIndex),
			)
		}
	}
}

/*
TestVarVectorTranspose_Plus5
Description:

	Testing the Plus operator between a VarVectorTranspose and a VarVectorTranspose. All vectors are of the same size.
	No overlap between elements.
*/
func TestVarVectorTranspose_Plus5(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus5")
	vec1 := m.AddVariableVector(desLength)
	vec3 := m.AddVariableVector(desLength)

	// Algorithm
	sum3, err := vec1.Plus(vec3)
	if err != nil {
		t.Errorf("There was an error computing addition: %v", err)
	}

	sum3AsVLE, ok := sum3.(optim.VectorLinearExpr)
	if !ok {
		t.Errorf(
			"There was an issue converting sum3 (type %T) to type optim.VectorLinearExpr.",
			sum3,
		)
	}

	// Check values of the vector of variables
	for vecIndex := vec3.Len(); vecIndex < vec3.Len()+vec1.Len(); vecIndex++ {
		// Check that values of sum3AsVLE and vec1 match
		if sum3AsVLE.X.AtVec(vecIndex).(optim.Variable).ID != vec1.AtVec(vecIndex-vec3.Len()).(optim.Variable).ID {
			t.Errorf(
				"Expected the value at index in sum3.X[%v] (%v) to be the same as vec1[%v] (%v).",
				vecIndex,
				sum3AsVLE.X.AtVec(vecIndex),
				vecIndex,
				vec1.AtVec(vecIndex-vec3.Len()),
			)
		}
	}

	// Check values of the matrix multiplier.
	for rowIndex := 0; rowIndex < desLength; rowIndex++ {
		// Get elements as needed.
		vec1Atr := vec1.AtVec(rowIndex)
		vec3Atr := vec3.AtVec(rowIndex)

		vec1AtRIndex, _ := optim.FindInSlice(vec1Atr, sum3AsVLE.X.Elements)
		vec3AtRIndex, _ := optim.FindInSlice(vec3Atr, sum3AsVLE.X.Elements)

		// Iterate through all columns (all variables)
		for colIndex := 0; colIndex < sum3AsVLE.X.Len(); colIndex++ {

			switch {
			case (colIndex == vec1AtRIndex) && (vec1AtRIndex == vec3AtRIndex):
				if sum3AsVLE.L.At(rowIndex, colIndex) != 2.0 {
					t.Errorf(
						"Expected L(%v,%v) to be 2.0; received %v",
						rowIndex, colIndex,
						sum3AsVLE.L.At(rowIndex, colIndex),
					)
				}
			case (colIndex == vec1AtRIndex) && (vec1AtRIndex != vec3AtRIndex):
				if sum3AsVLE.L.At(rowIndex, colIndex) != 1.0 {
					t.Errorf(
						"Expected L(%v,%v) to be 1.0; received %v",
						rowIndex, colIndex,
						sum3AsVLE.L.At(rowIndex, colIndex),
					)
				}
			case (colIndex == vec3AtRIndex) && (vec1AtRIndex != vec3AtRIndex):
				if sum3AsVLE.L.At(rowIndex, colIndex) != 1.0 {
					t.Errorf(
						"Expected L(%v,%v) to be 1.0; received %v",
						rowIndex, colIndex,
						sum3AsVLE.L.At(rowIndex, colIndex),
					)
				}
			default:
				// All other elements should be 0.0
				if sum3AsVLE.L.At(rowIndex, colIndex) != 0.0 {
					t.Errorf(
						"Expected L(%v,%v) to be 0.0; received %v",
						rowIndex, colIndex,
						sum3AsVLE.L.At(rowIndex, colIndex),
					)
				}
			}

		}
	}

	// Check offset vector (should be zeros)
	for vecIndex := 0; vecIndex < vec3.Len(); vecIndex++ {
		// Check that values of sum3AsVLE.X matches vec2 at the appropriate indices.
		if sum3AsVLE.X.AtVec(vecIndex) != vec3.AtVec(vecIndex) {
			t.Errorf(
				"Expected the value at index in sum3.X[%v] (%v) to be the same as vec2[%v] (%v).",
				vecIndex,
				sum3AsVLE.X.AtVec(vecIndex).(optim.Variable).ID,
				vecIndex-vec1.Len(),
				vec3.AtVec(vecIndex),
			)
		}
	}

	// Check the values of the constant vector
	for vecIndex := 0; vecIndex < sum3AsVLE.Len(); vecIndex++ {
		// Check that values of sum3AsVLE and vec1 match
		if sum3AsVLE.C.AtVec(vecIndex) != 0.0 {
			t.Errorf(
				"Expected the value of constant to be zero vector, but sum3.C[%v] = %v!",
				vecIndex,
				sum3AsVLE.C.AtVec(vecIndex),
			)
		}
	}
}

/*
TestVarVectorTranspose_Plus6
Description:

	Testing the Plus operator between a VarVectorTranspose and a KVector. Proper sizes were given.
*/
func TestVarVectorTranspose_Plus6(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus6")
	vec1 := m.AddVariableVector(desLength).Transpose()
	k2 := optim.OnesVector(desLength)

	// Algorithm
	_, err := vec1.Plus(k2)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot add VarVectorTranspose to a normal vector %v (%T); Try transposing one or the other!",
			k2, k2,
		),
	) {
		t.Errorf("There was an unexpected error computing addition: %v", err)
	}
}

/*
TestVarVectorTranspose_Plus7
Description:

	Testing the Plus operator between a VarVectorTranspose and a KVector. Proper sizes were given.
*/
func TestVarVectorTranspose_Plus7(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus7")
	vec1 := m.AddVariableVector(desLength)
	k2 := optim.KVector(optim.OnesVector(desLength))

	// Algorithm
	_, err := vec1.Transpose().Plus(k2)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot add VarVectorTranspose to a normal vector %v (%T); Try transposing one or the other!",
			k2, k2,
		),
	) {
		t.Errorf("There was an error computing addition: %v", err)
	}
}

/*
TestVarVectorTranspose_Plus8
Description:

	Testing the Plus operator between a VarVectorTranspose and a KVectorTranspose.
	Improper lengths are used.
*/
func TestVarVectorTranspose_Plus8(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus8")
	vec1 := m.AddVariableVector(desLength)
	k2 := optim.KVector(optim.OnesVector(desLength - 1))

	// Algorithm
	_, err := vec1.Transpose().Plus(k2.Transpose())
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"The lengths of two vectors in Plus must match! VarVectorTranspose has dimension %v, KVector has dimension %v",
			vec1.Transpose().Len(),
			k2.Transpose().Len(),
		),
	) {
		t.Errorf("There was an error computing addition: %v", err)
	}
}

/*
TestVarVectorTranspose_Plus9
Description:

	Testing the Plus operator between a VarVectorTranspose and a VarVector.
*/
func TestVarVectorTranspose_Plus9(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus9")
	vec1 := m.AddVariableVector(desLength)
	vec2 := m.AddVariableVector(desLength)

	// Algorithm
	_, err := vec1.Transpose().Plus(vec2)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot add VarVectorTranspose to a normal vector %v (%T); Try transposing one or the other!",
			vec2, vec2,
		),
	) {
		t.Errorf("There was an error computing addition: %v", err)
	}
}

/*
TestVarVectorTranspose_Plus10
Description:

	Testing the Plus operator between a VarVectorTranspose and a ScalarLinearExpr.
*/
func TestVarVectorTranspose_Plus10(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus10")
	vec1 := m.AddVariableVector(desLength)
	vec2 := m.AddVariableVector(desLength)
	k3 := optim.KVector(optim.OnesVector(desLength))

	// Algorithm
	expr1, err := vec2.Plus(k3)
	if err != nil {
		t.Errorf("Unexpected error when adding together VarVectorTranspose and KVectorTranspose: %v", err)
	}

	_, err = vec1.Transpose().Plus(expr1)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot add VarVectorTranspose with a normal vector %v (%T); Try transposing one or the other!",
			expr1, expr1,
		),
	) {
		t.Errorf("There was an error computing addition: %v", err)
	}
}

/*
TestVarVectorTranspose_Plus11
Description:

	Testing the Plus operator between a VarVectorTranspose and a bool.
*/
func TestVarVectorTranspose_Plus11(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus11")
	vec1 := m.AddVariableVector(desLength)
	b2 := false

	// Algorithm
	_, err := vec1.Transpose().Plus(b2)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Unrecognized expression type %T for addition of VarVectorTranspose vvt.Plus(%v)!",
			b2, b2,
		),
	) {
		t.Errorf("There was an error computing addition: %v", err)
	}
}

/*
TestVarVectorTranspose_AtVec1
Description:

	Testing the At operator on a VarVectorTranspose object.
*/
func TestVarVectorTranspose_AtVec1(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("AtVec1")
	vec1 := m.AddVariableVector(desLength)
	idx1 := 2

	// Algorithm

	vec1AtIdx1Casted, tf := vec1.AtVec(idx1).(optim.Variable)
	if !tf {
		t.Errorf(
			"vec1.AtVec(%v) is not of type optim.Variable; instead it is %T",
			idx1,
			vec1.AtVec(idx1),
		)
	}

	if vec1AtIdx1Casted.ID != vec1.Elements[idx1].ID {
		t.Errorf(
			"vec1.AtVec(%v) = %v != %v = vec1.Elements[%v]",
			idx1,
			vec1AtIdx1Casted,
			vec1.Elements[idx1],
			idx1,
		)
	}
}

/*
TestVarVectorTranspose_Mult1
Description:

	Verifies that the Mult() operator throws an error in its current form.
*/
func TestVarVectorTranspose_Mult1(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("AtVec1")
	vec1 := m.AddVariableVector(desLength).Transpose()

	// Check
	_, err := vec1.Mult(2.9)
	if !strings.Contains(
		err.Error(),
		"The Mult() method for VarVectorTranspose is not implemented yet!",
	) {
		t.Errorf("Unexpected error: %v", err)
	}
}

/*
TestVarVectorTranspose_LessEq1
Description:

	Verifies that the LessEq method throws an error when the KVectorTranspose is
	of the wrong length.
*/
func TestVarVectorTranspose_LessEq1(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("AtVec1")
	vec1 := m.AddVariableVector(desLength).Transpose()
	kv2 := optim.KVectorTranspose(optim.OnesVector(desLength - 1))

	// Algorithm
	_, err := vec1.LessEq(kv2)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"The two inputs to comparison '%v' must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
			optim.SenseLessThanEqual,
			vec1.Len(),
			kv2.Len(),
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}
}

/*
TestVarVectorTranspose_GreaterEq1
Description:

	Verifies that the GreaterEq method throws an error when the KVectorTranspose is
	of the wrong length.
*/
func TestVarVectorTranspose_GreaterEq1(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("AtVec1")
	vec1 := m.AddVariableVector(desLength).Transpose()
	kv2 := optim.OnesVector(desLength)

	// Algorithm
	constraint0, err := vec1.GreaterEq(kv2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if constraint0.Sense != optim.SenseGreaterThanEqual {
		t.Errorf("Constraint should be a GreaterThanEqual constraint, but found the sense: %v", constraint0.Sense)
	}

	lhs0, ok1 := constraint0.LeftHandSide.(optim.VarVectorTranspose)
	if !ok1 {
		t.Errorf("Unexpected left hand side type %T for %v", lhs0, lhs0)
	}

}
