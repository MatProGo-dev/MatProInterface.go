package optim_test

import (
	"fmt"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"strings"
	"testing"
)

func TestVarVector_Length1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	if vv1.Length() != 2 {
		t.Errorf("The length of vv1 was %v; expected %v", vv1.Length(), 2)
	}

}

/*
TestVarVector_Length2
Description:

	Tests that a larger vector variable (contains 5 elements) properly returns the right length.
*/
func TestVarVector_Length2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y, x, y, x},
	}

	if vv1.Length() != 5 {
		t.Errorf("The length of vv1 was %v; expected %v", vv1.Length(), 5)
	}

}

/*
TestVarVector_At1
Description:

	Tests whether or not we can properly retrieve an element from a given vector.
*/
func TestVarVector_At1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	extractedV := vv1.AtVec(1)
	if extractedV.(optim.Variable) != y {
		t.Errorf("Expected for extracted variable, %v, to be the same as %v. They were different!", extractedV, y)
	}
}

/*
TestVarVector_At2
Description:

	Tests whether or not we can properly retrieve an element from a given vector.
	Makes sure that if we change the extracted vector, it does not effect the element saved in the slice.
*/
func TestVarVector_At2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	extractedV := vv1.AtVec(1).(optim.Variable)
	extractedV.ID = 100

	if extractedV == y {
		t.Errorf("Expected for extracted variable, %v, to be DIFFERENT from %v. They were the same!", extractedV, y)
	}
}

/*
TestVarVector_VariableIDs1
Description:

	This test will check to see if 2 unique ids in a VariableVector object will be returned correctly when
	the VariableIDs method is called.
*/
func TestVarVector_VariableIDs1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVector{
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
TestVarVector_VariableIDs2
Description:

	This test will check to see if a single unique id in a large VariableVector object will be returned correctly when
	the VariableIDs method is called.
*/
func TestVarVector_VariableIDs2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVector{
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
TestVarVector_Constant1
Description:

	This test verifies that the constant method returns an all zero vector for any varvector object.
*/
func TestVarVector_Constant1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVector{
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
TestVarVector_Constant2
Description:

	This test verifies that the constant method returns an all zero vector for any varvector object.
	This one will be extremely long.
*/
func TestVarVector_Constant2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVector{
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
TestVarVector_Eq1
Description:

	This test verifies that the Eq method works between a varvector and another object.
*/
func TestVarVector_Eq1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y},
	}

	zerosAsVecDense := optim.ZerosVector(vv1.Len())
	zerosAsKVector := optim.KVector(zerosAsVecDense)

	// Verify that constraint can be created with no issues.
	_, err := vv1.Eq(zerosAsKVector)
	if err != nil {
		t.Errorf("There was an issue creating an equality constraint between vv1 and the zero vector.")
	}
}

/*
TestVarVector_Eq2
Description:

	This test verifies that the Eq method works between a varvector and another object.
	Comparison should be between var vector and an unsupported type.
*/
func TestVarVector_Eq2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y},
	}

	badRHS := false

	// Verify that constraint can be created with no issues.
	_, err := vv1.Eq(badRHS)
	expectedError := fmt.Sprintf("The Eq() method for VarVector is not implemented yet for type %T!", badRHS)
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Expected error \"%v\"; received \"%v\"", expectedError, err)
	}
}

/*
TestVarVector_Eq2
Description:

	This test verifies that the Eq method works between a varvector and another var vector.
*/
func TestVarVector_Eq3(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y},
	}

	vv2 := optim.VarVector{
		Elements: []optim.Variable{y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x},
	}

	// Verify that constraint can be created with no issues.
	_, err := vv1.Eq(vv2)
	if err != nil {
		t.Errorf("There was an error creating equality constraint between the two varvectors: %v", err)
	}
}

/*
TestVarVector_Comparison1
Description:

	Tests how well the comparison function works with a VectorLinearExpression comparison.
*/
func TestVarVector_Comparison1(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel()
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
TestVarVector_Comparison2
Description:

	Tests how well the comparison function works with a VectorLinearExpression comparison.
	Valid comparison of
*/
func TestVarVector_Comparison2(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel()
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
TestVarVector_Plus1
Description:

	Testing the Plus operator between a VarVector and a KVector. Proper sizes were given.
*/
func TestVarVector_Plus1(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel()
	vec1 := m.AddVariableVector(desLength)
	k2 := optim.KVector(optim.OnesVector(desLength))

	// Algorithm
	sum3, err := vec1.Plus(k2)
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
TestVarVector_Plus2
Description:

	Testing the Plus operator between a VarVector and a KVector. Incorrect sizes were given.
*/
func TestVarVector_Plus2(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel()
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
TestVarVector_Plus3
Description:

	Testing the Plus operator between a VarVector and a KVector. Proper sizes were given.
*/
func TestVarVector_Plus3(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel()
	vec1 := m.AddVariableVector(desLength)
	k2 := optim.OnesVector(desLength)

	// Algorithm
	sum3, err := vec1.Plus(k2)
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
		if sum3AsVLE.C.AtVec(vecIndex) != k2.AtVec(vecIndex) {
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
TestVarVector_Plus4
Description:

	Testing the Plus operator between a VarVector and a VarVector. All vectors are of same size. Some overlap in the variables but not all.
*/
func TestVarVector_Plus4(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel()
	vec1 := m.AddVariableVector(desLength)
	vec2 := m.AddVariableVector(desLength - 2)
	vec3 := optim.VarVector{
		append(vec2.Elements, vec1.AtVec(0).(optim.Variable), vec1.AtVec(1).(optim.Variable)),
	}

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
TestVarVector_Plus5
Description:

	Testing the Plus operator between a VarVector and a VarVector. All vectors are of the same size.
	No overlap between elements.
*/
func TestVarVector_Plus5(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel()
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
TestVarVector_AtVec1
Description:

	Testing the At operator on a VarVector object.
*/
func TestVarVector_AtVec1(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel()
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
