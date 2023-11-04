package optim_test

import (
	"fmt"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"gonum.org/v1/gonum/mat"
	"strings"
	"testing"
)

func TestVarVector_Length1(t *testing.T) {
	m := optim.NewModel("Length1")
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
	m := optim.NewModel("Length2")
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
	m := optim.NewModel("At1")
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
	m := optim.NewModel("At2")
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
	m := optim.NewModel("VariableIDs1")
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
	m := optim.NewModel("VariableIDs2")
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
TestVarVector_NumVars1
Description:

	This test will check to see vector of length 10
	contains 2 n Vars.
*/
func TestVarVector_NumVars1(t *testing.T) {
	m := optim.NewModel("NumVars1")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y, x},
	}

	nv := vv1.NumVars()
	// Check to see that only x has ids in extractedIDs
	if nv != 2 {
		t.Errorf(
			"There are two variable ID in the VarVector vv1, and yet %v was returned by the NumVars() method.",
			nv,
		)
	}
}

/*
TestVarVector_Constant1
Description:

	This test verifies that the constant method returns an all zero vector for any varvector object.
*/
func TestVarVector_Constant1(t *testing.T) {
	m := optim.NewModel("Constant1")
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
	m := optim.NewModel("Constant2")
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
TestVarVector_LinearCoeff1
Description:

	This test will check to see vector of length 10
	contains 10 n Vars.
*/
func TestVarVector_LinearCoeff1(t *testing.T) {
	m := optim.NewModel("LinearCoeff1")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y, x},
	}

	identity0 := vv1.LinearCoeff()
	// Check to see that only x has ids in extractedIDs
	nR, nC := identity0.Dims()
	if (nR != 3) || (nC != 3) {
		t.Errorf(
			"Exoected identity0 to be of shape (3,3); received (%v,%v)",
			nR,
			nC,
		)
	}

	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			if rowIndex == colIndex {
				if identity0.At(rowIndex, colIndex) != 1.0 {
					t.Errorf(
						"Expected diagonal elements to be 1.0; received %v",
						identity0.At(rowIndex, colIndex),
					)
				}
			}

			if rowIndex != colIndex {
				if identity0.At(rowIndex, colIndex) != 0.0 {
					t.Errorf(
						"Expected diagonal elements to be 0.0; received %v",
						identity0.At(rowIndex, colIndex),
					)
				}
			}

		}

	}
}

/*
TestVarVector_Eq1
Description:

	This test verifies that the Eq method works between a varvector and another object.
*/
func TestVarVector_Eq1(t *testing.T) {
	m := optim.NewModel("Eq1")
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
	m := optim.NewModel("Eq2")
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
	m := optim.NewModel("Eq3")
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
TestVarVector_Comparison2
Description:

	Tests how well the comparison function works with a VectorLinearExpression comparison.
	Valid comparison of
*/
func TestVarVector_Comparison2(t *testing.T) {
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
TestVarVector_Comparison3
Description:

	Tests that the Comparison() Method works well for future users.
*/
func TestVarVector_Comparison3(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("AtVec1")
	vec1 := m.AddVariableVector(desLength)
	kv1 := optim.KVector(optim.OnesVector(desLength)).Transpose()

	// Compare
	_, err := vec1.GreaterEq(kv1)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot compare VarVector with a transposed vector %v (%T); Try transposing one or the other!",
			kv1, kv1,
		),
	) {
		t.Errorf("Unexpected error when comparing two vectors: %v", err)
	}

}

/*
TestVarVector_Comparison4
Description:

	Tests that the Comparison() Method works for two VarVectors of different lengths.
*/
func TestVarVector_Comparison4(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("AtVec1")
	vec1 := m.AddVariableVector(desLength)
	vec2 := m.AddVariableVector(desLength - 1)

	// Compare
	_, err := vec1.Comparison(vec2, optim.SenseEqual)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"The two inputs to comparison '%v' must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
			optim.SenseEqual,
			vec1.Len(),
			vec2.Len(),
		),
	) {
		t.Errorf("Unexpected error when comparing VarVectors of different lengths: %v", err)
	}

}

/*
TestVarVector_Comparison5
Description:

	Tests that the Comparison() Method works well for a VarVector and
	a VarVectorTranspose.
*/
func TestVarVector_Comparison5(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("AtVec1")
	vec1 := m.AddVariableVector(desLength)
	vec2 := m.AddVariableVector(desLength).Transpose()

	// Compare
	_, err := vec1.Comparison(vec2, optim.SenseEqual)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot compare VarVector with a transposed vector %v (%T); Try transposing one or the other!",
			vec2, vec2,
		),
	) {
		t.Errorf("Unexpected error when comparing VarVectors of different lengths: %v", err)
	}

}

/*
TestVarVector_Comparison6
Description:

	Tests that the Comparison() Method works well for a VarVector and
	a VectorLinearExpressionTranspose.
*/
func TestVarVector_Comparison6(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("AtVec1")
	vec1 := m.AddVariableVector(desLength)
	vec2 := m.AddVariableVector(desLength)

	sum1, _ := vec2.Plus(optim.OnesVector(desLength))

	// Compare
	_, err := vec1.Comparison(sum1.Transpose(), optim.SenseEqual)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot compare VarVector with a transposed vector %v (%T); Try transposing one or the other!",
			sum1.Transpose(), sum1.Transpose(),
		),
	) {
		t.Errorf("Unexpected error when comparing VarVectors of different lengths: %v", err)
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
	m := optim.NewModel("Plus1")
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
TestVarVector_Plus3
Description:

	Testing the Plus operator between a VarVector and a KVector. Proper sizes were given.
*/
func TestVarVector_Plus3(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus3")
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
	m := optim.NewModel("Plus4")
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
TestVarVector_Plus6
Description:

	Testing the Plus operator between a VarVector and a KVectorTranspose.
	All vectors are of the same size.
	No overlap between elements.
*/
func TestVarVector_Plus6(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus6")
	vec1 := m.AddVariableVector(desLength)
	k2 := optim.KVectorTranspose(optim.OnesVector(desLength))

	// Algorithm
	_, err := vec1.Plus(k2)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot add VarVector with a transposed vector %v (%T); Try transposing one or the other!",
			k2, k2,
		),
	) {
		t.Errorf("There was an unexpected error computing addition: %v", err)
	}

}

/*
TestVarVector_Plus7
Description:

	Testing the Plus operator between a VarVector and a VarVectorTranspose.
*/
func TestVarVector_Plus7(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus5")
	vec1 := m.AddVariableVector(desLength)
	vec3 := m.AddVariableVector(desLength).Transpose()

	// Algorithm
	_, err := vec1.Plus(vec3)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot add VarVector with a transposed vector %v (%T); Try transposing one or the other!",
			vec3, vec3,
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}
}

/*
TestVarVector_Plus8
Description:

	Testing the Plus operator between a VarVector and a VectorLinearExpressionTranspose.
*/
func TestVarVector_Plus8(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus8")
	vec1 := m.AddVariableVector(desLength)

	vec3 := m.AddVariableVector(desLength)
	sum, _ := vec3.Plus(optim.OnesVector(desLength))

	// Algorithm
	_, err := vec1.Plus(sum.Transpose())
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot add VarVector with a transposed vector %v (%T); Try transposing one or the other!",
			sum.Transpose(), sum.Transpose(),
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}
}

/*
TestVarVector_Plus9
Description:

	Testing the Plus operator between a VarVector and a VectorLinearExpressionTranspose.
*/
func TestVarVector_Plus9(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("Plus9")
	vec1 := m.AddVariableVector(desLength)

	b1 := false

	// Algorithm
	_, err := vec1.Plus(b1)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Unrecognized expression type %T for addition of VarVector vv.Plus(%v)!",
			b1, b1,
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}
}

/*
TestVarVector_Mult1
Descripiion:

	Tests that the Mult() method currently returns errors.
*/
func TestVarVector_Mult1(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("AtVec1")
	vec1 := m.AddVariableVector(desLength)

	// Compute
	_, err := vec1.Mult(21.0)
	if !strings.Contains(
		err.Error(),
		"The Mult() method for VarVector is not implemented yet!",
	) {
		t.Errorf("Unexpected error: %v", err)
	}
}

/*
TestVarVector_GreaterEq1
Description:

	Tests that the GreaterEq() Method works well for future users.
*/
func TestVarVector_GreaterEq1(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel("AtVec1")
	vec1 := m.AddVariableVector(desLength)
	kv1 := optim.KVector(optim.OnesVector(desLength - 1))

	// Compare
	_, err := vec1.GreaterEq(kv1)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"The two inputs to comparison '%v' must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
			optim.SenseGreaterThanEqual,
			vec1.Len(),
			kv1.Len(),
		),
	) {
		t.Errorf("Unexpected error when comparing two vectors: %v", err)
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
TestVarVector_1
Description:

	Tests that the error catching behavior works when
	VarVector is malformed.
*/
func TestVarVector_Check1(t *testing.T) {
	// Constants
	m := optim.NewModel("TestVarVector_Check1")
	v1 := m.AddVariable()
	v2 := m.AddVariable()
	v2.Lower = 1
	v2.Upper = 0

	vv3 := optim.VarVector{
		Elements: []optim.Variable{v1, v2},
	}

	// Run Check
	err := vv3.Check()
	if err == nil {
		t.Errorf("No error was thrown, but we expected one!")
	} else {
		if !strings.Contains(
			err.Error(),
			fmt.Sprintf(
				"element %v has an issue: %v",
				1, fmt.Sprintf(
					"lower bound (%v) of variable is above upper bound (%v).",
					v2.Lower, v2.Upper,
				),
			),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestVarVector_Multiply1
Description:

	Tests that the error catching behavior works when
	VarVector is malformed.
*/
func TestVarVector_Multiply1(t *testing.T) {
	// Constants
	m := optim.NewModel("TestVarVector_Multiply1")
	v1 := m.AddVariable()
	v2 := m.AddVariable()
	v2.Lower = 1
	v2.Upper = 0

	vv3 := optim.VarVector{
		Elements: []optim.Variable{v1, v2},
	}

	// Run Check
	_, err := vv3.Multiply(3.0)
	if err == nil {
		t.Errorf("No error was thrown, but we expected one!")
	} else {
		if !strings.Contains(
			err.Error(),
			fmt.Sprintf(
				"element %v has an issue: %v",
				1, fmt.Sprintf(
					"lower bound (%v) of variable is above upper bound (%v).",
					v2.Lower, v2.Upper,
				),
			),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestVarVector_Multiply2
Description:

	Tests that the error catching behavior works when
	a bad error input is given.
*/
func TestVarVector_Multiply2(t *testing.T) {
	// Constants
	m := optim.NewModel("TestVarVector_Multiply2")
	vv1 := m.AddVariableVector(2)
	err2 := fmt.Errorf("test")

	// Run Check
	_, err := vv1.Multiply(3.0, err2)
	if err == nil {
		t.Errorf("No error was thrown, but we expected one!")
	} else {
		if !strings.Contains(
			err.Error(),
			err2.Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestVarVector_Multiply3
Description:

	Tests that the error catching behavior works when
	an input with bad shape is given.
*/
func TestVarVector_Multiply3(t *testing.T) {
	// Constants
	m := optim.NewModel("TestVarVector_Multiply3")
	vv1 := m.AddVariableVector(2)
	vv2 := m.AddVariableVector(4)

	// Run Check
	_, err := vv1.Multiply(vv2)
	if err == nil {
		t.Errorf("No error was thrown, but we expected one!")
	} else {
		if !strings.Contains(
			err.Error(),
			optim.DimensionError{
				Operation: "Multiply",
				Arg1:      vv1,
				Arg2:      vv2,
			}.Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestVarVector_Multiply4
Description:

	Tests that the error catching behavior works when
	a scalar float is given. Should result in VectorLinearExpression
*/
func TestVarVector_Multiply4(t *testing.T) {
	// Constants
	m := optim.NewModel("TestVarVector_Multiply4")
	vv1 := m.AddVariableVector(2)
	f2 := 3.1415

	// Run Check
	prod, err := vv1.Multiply(f2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	prodAsVLE, tf := prod.(optim.VectorLinearExpr)
	if !tf {
		t.Errorf("expected product to be of type VectorLinearExpr; received %T instead", prod)
	}

	// Investigate components of VLE
	for rowIndex := 0; rowIndex < vv1.Len(); rowIndex++ {
		for colIndex := 0; colIndex < vv1.Len(); colIndex++ {
			if (rowIndex == colIndex) && (prodAsVLE.L.At(rowIndex, colIndex) != 1.0*f2) {
				t.Errorf(
					"prod.L[%v,%v] = %v =/= 1.0 as expected",
					rowIndex, colIndex,
					prodAsVLE.L.At(rowIndex, colIndex),
				)
			}

			if (rowIndex != colIndex) && (prodAsVLE.L.At(rowIndex, colIndex) != 0.0) {
				t.Errorf(
					"prod.L[%v,%v] = %v =/= 0.0 as expected",
					rowIndex, colIndex,
					prodAsVLE.L.At(rowIndex, colIndex),
				)
			}

		}

	}

}

/*
TestVarVector_Multiply5
Description:

	Tests that the error catching behavior works when
	a scalar K is given. Should result in VectorLinearExpression
*/
func TestVarVector_Multiply5(t *testing.T) {
	// Constants
	m := optim.NewModel("TestVarVector_Multiply5")
	vv1 := m.AddVariableVector(2)
	f2 := 3.1415

	// Run Check
	prod, err := vv1.Multiply(optim.K(f2))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	prodAsVLE, tf := prod.(optim.VectorLinearExpr)
	if !tf {
		t.Errorf("expected product to be of type VectorLinearExpr; received %T instead", prod)
	}

	// Investigate components of VLE
	for rowIndex := 0; rowIndex < vv1.Len(); rowIndex++ {
		for colIndex := 0; colIndex < vv1.Len(); colIndex++ {
			if (rowIndex == colIndex) && (prodAsVLE.L.At(rowIndex, colIndex) != 1.0*f2) {
				t.Errorf(
					"prod.L[%v,%v] = %v =/= 1.0 as expected",
					rowIndex, colIndex,
					prodAsVLE.L.At(rowIndex, colIndex),
				)
			}

			if (rowIndex != colIndex) && (prodAsVLE.L.At(rowIndex, colIndex) != 0.0) {
				t.Errorf(
					"prod.L[%v,%v] = %v =/= 0.0 as expected",
					rowIndex, colIndex,
					prodAsVLE.L.At(rowIndex, colIndex),
				)
			}

		}

	}

}

/*
TestVarVector_Multiply6
Description:

	Tests that the error catching behavior works when
	a one element KVector is given. Should result in VectorLinearExpression
*/
func TestVarVector_Multiply6(t *testing.T) {
	// Constants
	m := optim.NewModel("TestVarVector_Multiply3")
	vv1 := m.AddVariableVector(2)
	f2 := 3.1415

	vd3 := mat.NewVecDense(1, []float64{f2})

	// Run Check
	prod, err := vv1.Multiply(optim.KVector(*vd3))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	prodAsVLE, tf := prod.(optim.VectorLinearExpr)
	if !tf {
		t.Errorf("expected product to be of type VectorLinearExpr; received %T instead", prod)
	}

	// Investigate components of VLE
	for rowIndex := 0; rowIndex < vv1.Len(); rowIndex++ {
		for colIndex := 0; colIndex < vv1.Len(); colIndex++ {
			if (rowIndex == colIndex) && (prodAsVLE.L.At(rowIndex, colIndex) != 1.0*f2) {
				t.Errorf(
					"prod.L[%v,%v] = %v =/= 1.0 as expected",
					rowIndex, colIndex,
					prodAsVLE.L.At(rowIndex, colIndex),
				)
			}

			if (rowIndex != colIndex) && (prodAsVLE.L.At(rowIndex, colIndex) != 0.0) {
				t.Errorf(
					"prod.L[%v,%v] = %v =/= 0.0 as expected",
					rowIndex, colIndex,
					prodAsVLE.L.At(rowIndex, colIndex),
				)
			}

		}

	}
}

/*
TestVarVector_Multiply7
Description:

	Tests that the error catching behavior works when
	a one element KVectorTranspose is given. Should result in VectorLinearExpression
*/
func TestVarVector_Multiply7(t *testing.T) {
	// Constants
	m := optim.NewModel("TestVarVector_Multiply7")
	vv1 := m.AddVariableVector(2)
	f2 := 3.1415

	vd3 := mat.NewVecDense(1, []float64{f2})

	// Run Check
	prod, err := vv1.Multiply(optim.KVectorTranspose(*vd3))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	prodAsVLE, tf := prod.(optim.VectorLinearExpr)
	if !tf {
		t.Errorf("expected product to be of type VectorLinearExpr; received %T instead", prod)
	}

	// Investigate components of VLE
	for rowIndex := 0; rowIndex < vv1.Len(); rowIndex++ {
		for colIndex := 0; colIndex < vv1.Len(); colIndex++ {
			if (rowIndex == colIndex) && (prodAsVLE.L.At(rowIndex, colIndex) != 1.0*f2) {
				t.Errorf(
					"prod.L[%v,%v] = %v =/= 1.0 as expected",
					rowIndex, colIndex,
					prodAsVLE.L.At(rowIndex, colIndex),
				)
			}

			if (rowIndex != colIndex) && (prodAsVLE.L.At(rowIndex, colIndex) != 0.0) {
				t.Errorf(
					"prod.L[%v,%v] = %v =/= 0.0 as expected",
					rowIndex, colIndex,
					prodAsVLE.L.At(rowIndex, colIndex),
				)
			}

		}

	}

}

/*
TestVarVector_Multiply8
Description:

	Tests that the error catching behavior works when
	a non-single element KVectorTranspose is given. Should result in VectorLinearExpression
*/
func TestVarVector_Multiply8(t *testing.T) {
	// Constants
	m := optim.NewModel("TestVarVector_Multiply8")
	vv1 := m.AddVariableVector(2)
	f2 := 3.1415

	vd3 := mat.NewVecDense(2, []float64{f2, f2})

	// Run Check
	_, err := vv1.Multiply(optim.KVectorTranspose(*vd3))
	if err == nil {
		t.Errorf("expected error, but received none!")
	} else {
		if !strings.Contains(
			err.Error(),
			fmt.Sprintf("cannot complete multiplication that will create matrix product! Submit an issue if you want this feature!"),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}

}
