package optim_test

import (
	"fmt"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"gonum.org/v1/gonum/mat"
	"strings"
	"testing"
)

/*
TestVectorLinearExpressionTranspose_Check1
Description:

	This test will evaluate whether or not the linear expression that has been given is valid.
	In this case, the VectorLinearExpressionTranspose is valid.
*/
func TestVectorLinearExpressionTranspose_Check1(t *testing.T) {
	m := optim.NewModel("Check1")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	L1 := *mat.NewDense(2, 2, []float64{1.0, 2.0, 3.0, 4.0})
	c1 := *mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		vv1, L1, c1,
	}

	// ve1 should pass all checks.
	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was supposed to be valid, but received an error: %v", err)
	}
}

/*
TestVectorLinearExpressionTranspose_Check2
Description:

	This test will evaluate whether or not the linear expression that has been given is valid.
	In this case, the VectorLinearExpressionTranspose is NOT valid. L is too big in rows.
*/
func TestVectorLinearExpressionTranspose_Check2(t *testing.T) {
	m := optim.NewModel("Check2")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	L1 := *mat.NewDense(3, 2, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})
	c1 := *mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		vv1, L1, c1,
	}

	// ve1 should pass all checks.
	err := ve1.Check()
	if err == nil {
		t.Errorf("The vector linear expression was supposed to be invalid, but received no errors!")
	}

	nL, mL := L1.Dims()
	if err.Error() != fmt.Sprintf("Dimension of L (%v x %v) and C (length %v) do not match!", nL, mL, c1.Len()) {
		t.Errorf("The vector linear expression was supposed to have dimension error #2, instead received %v", err)
	}
}

/*
TestVectorLinearExpressionTranspose_Check3
Description:

	This test will evaluate whether or not the linear expression that has been given is valid.
	In this case, the VectorLinearExpressionTranspose is NOT valid. L is too big in columns.
*/
func TestVectorLinearExpressionTranspose_Check3(t *testing.T) {
	m := optim.NewModel("Check3")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	L1 := *mat.NewDense(2, 3, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})
	c1 := *mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		vv1, L1, c1,
	}

	// ve1 should pass all checks.
	err := ve1.Check()
	if err == nil {
		t.Errorf("The vector linear expression was supposed to be invalid, but received no errors!")
	}

	nL, mL := L1.Dims()
	if err.Error() != fmt.Sprintf("Dimensions of L (%v x %v) and x (length %v) do not match appropriately.", nL, mL, vv1.Len()) {
		t.Errorf("The vector linear expression was supposed to have dimension error #1, instead received %v", err)
	}
}

/*
TestVectorLinearExpressionTranspose_VariableIDs1
Description:

	This test the VariableIDs() method when a variable vector with 2 unique vectors.
*/
func TestVectorLinearExpressionTranspose_VariableIDs1(t *testing.T) {
	m := optim.NewModel("VariableIDs1")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	L1 := *mat.NewDense(3, 2, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})
	c1 := *mat.NewVecDense(3, []float64{5.0, 6.0, 7.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		vv1, L1, c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// ve1 should pass all checks.
	extractedIDs := ve1.IDs()
	// Check to see that x and y have ids in extractedIDs
	if foundIndex, _ := optim.FindInSlice(x.ID, extractedIDs); foundIndex == -1 {
		t.Errorf("x was not found in vv1, but it should have been!")
	}

	if foundIndex, _ := optim.FindInSlice(y.ID, extractedIDs); foundIndex == -1 {
		t.Errorf("x was not found in vv1, but it should have been!")
	}
}

/*
TestVectorLinearExpressionTranspose_VariableIDs2
Description:

	This test the VariableIDs() method works for a variable vector with 1 unique vectors.
*/
func TestVectorLinearExpressionTranspose_VariableIDs2(t *testing.T) {
	m := optim.NewModel("VariableIDs2")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, x, x, x},
	}

	L1 := *mat.NewDense(2, 4, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0})
	c1 := *mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		vv1, L1, c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// ve1 should pass all checks.
	extractedIDs := ve1.IDs()
	// Check to see that x has id in extractedIDs (y should not be there)
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
TestVectorLinearExpressionTranspose_NumVars1
Description:

	This test the NumVars() method when a variable vector with a short, unique vectors.
*/
func TestVectorLinearExpressionTranspose_NumVars1(t *testing.T) {
	m := optim.NewModel("VariableIDs1")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	L1 := *mat.NewDense(3, 2, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})
	c1 := *mat.NewVecDense(3, []float64{5.0, 6.0, 7.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		vv1, L1, c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// ve1 should pass all checks.
	if ve1.NumVars() != 2 {
		t.Errorf("Unexpected number of variables. Received %v; expected 2", ve1.NumVars())
	}
}

/*
TestVectorLinearExpressionTranspose_Coeffs1
Description:

	This test the Coeffs() method which should return the matrix's elements in a prescribed order.
*/
func TestVectorLinearExpressionTranspose_Coeffs1(t *testing.T) {
	m := optim.NewModel("Coeffs1")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	LElts := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}
	L1 := *mat.NewDense(3, 2, LElts)
	c1 := *mat.NewVecDense(3, []float64{5.0, 6.0, 7.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		vv1, L1, c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// ve1 should pass all checks.
	m_L, n_L := L1.Dims()
	extractedCoeffMat := ve1.LinearCoeff()
	for rowIndex := 0; rowIndex < m_L; rowIndex++ {
		for colIndex := 0; colIndex < n_L; colIndex++ {
			// Compare each element of the matrix
			if L1.At(rowIndex, colIndex) != extractedCoeffMat.At(rowIndex, colIndex) {
				t.Errorf(
					"The extracted coefficient at index %v,%v (%v) is not the same as the given one (%v).",
					rowIndex, colIndex,
					extractedCoeffMat.At(rowIndex, colIndex),
					L1.At(rowIndex, colIndex),
				)
			}
		}
	}

}

/*
TestVectorLinearExpressionTranspose_Coeffs2
Description:

	This test the Coeffs() method which should return the matrix's elements in a prescribed order.
*/
func TestVectorLinearExpressionTranspose_Coeffs2(t *testing.T) {
	m := optim.NewModel("Coeffs2")
	x := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, x, x, x},
	}

	LElts := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	L1 := *mat.NewDense(2, 4, LElts)
	c1 := *mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		vv1, L1, c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// ve1 should pass all checks.
	m_L, n_L := L1.Dims()
	extractedCoeffMat := ve1.LinearCoeff()
	for rowIndex := 0; rowIndex < m_L; rowIndex++ {
		for colIndex := 0; colIndex < n_L; colIndex++ {
			// Compare each element of the matrix
			if L1.At(rowIndex, colIndex) != extractedCoeffMat.At(rowIndex, colIndex) {
				t.Errorf(
					"The extracted coefficient at index %v,%v (%v) is not the same as the given one (%v).",
					rowIndex, colIndex,
					extractedCoeffMat.At(rowIndex, colIndex),
					L1.At(rowIndex, colIndex),
				)
			}
		}
	}
}

/*
TestVectorLinearExpressionTranspose_LessEq1
Description:

	This tests that the less than or equal to command works with a constant input.
*/
func TestVectorLinearExpressionTranspose_LessEq1(t *testing.T) {
	// Constants
	m := optim.NewModel("TestVectorLinearExpressionTranspose_LessEq1")
	x := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, x, x, x},
	}

	LElts := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	L1 := mat.NewDense(2, 4, LElts)
	c1 := mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		vv1, *L1, *c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// Algorithm
	constr1, err := ve1.LessEq(
		optim.KVectorTranspose(optim.OnesVector(ve1.Len())),
	)
	if err != nil {
		t.Errorf("There was an error computing the constraint ve1 <= 2.0: %v", err)
	}

	if len(constr1.LeftHandSide.IDs()) != len(vv1.IDs()) {
		t.Errorf("Left Hand")
	}

	if constr1.Sense != optim.SenseLessThanEqual {
		t.Errorf("Expected constraint's sense to be SenseLessThanEqual; received %v", optim.SenseGreaterThanEqual)
	}

}

/*
TestVectorLinearExpressionTranspose_GreaterEq1
Description:

	This tests that the greater than or equal to command works with a constant input.
*/
func TestVectorLinearExpressionTranspose_GreaterEq1(t *testing.T) {
	// Constants
	m := optim.NewModel("TestVectorLinearExpressionTranspose_GreaterEq1")
	x := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, x, x, x},
	}

	LElts := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	L1 := mat.NewDense(2, 4, LElts)
	c1 := mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		vv1, *L1, *c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// Algorithm
	constr1, err := ve1.GreaterEq(
		optim.KVectorTranspose(optim.OnesVector(ve1.Len())),
	)
	if err != nil {
		t.Errorf("There was an error computing the constraint ve1 <= 2.0: %v", err)
	}

	if len(constr1.LeftHandSide.IDs()) != len(vv1.IDs()) {
		t.Errorf("Left Hand")
	}

	if constr1.Sense != optim.SenseGreaterThanEqual {
		t.Errorf("Expected constraint's sense to be SenseLessThanEqual; received %v", optim.SenseGreaterThanEqual)
	}

}

/*
TestVectorLinearExpressionTranspose_Mult1
Description:

	Tests that the unfinished Mult() function is properly returning errors.
*/
func TestVectorLinearExpressionTranspose_Mult1(t *testing.T) {
	// Constants

	// Create model
	m := optim.NewModel("TestVectorLinearExpressionTranspose_Mult1")
	x := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, x, x, x},
	}

	LElts := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	L1 := mat.NewDense(2, 4, LElts)
	c1 := mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		vv1, *L1, *c1,
	}

	// Algorithm
	_, err := ve1.Mult(2.0)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf("The multiplication method has not yet been implemented!"),
	) {
		t.Errorf("Unexpected error after multiplication: %v", err)
	}
}

/*
TestVectorLinearExpressionTranspose_Eq1
Description:

	Tests whether or not an equality constraint between a ones vector and a standard vector variable works well.
	Eq comparison between:
	- Vector Linear Expression, and
	- mat.VecDense
*/
func TestVectorLinearExpressionTranspose_Eq1(t *testing.T) {
	// Constants
	m := optim.NewModel("VLETranspose-Eq1")

	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}
	c := optim.ZerosVector(2)
	vle1 := optim.VectorLinearExpressionTranspose{
		vv1,
		optim.Identity(2),
		c,
	}

	ones0 := optim.KVectorTranspose(optim.OnesVector(2))

	// Create Constraint
	constr, err := vle1.Eq(ones0)
	if err != nil {
		t.Errorf("There was a problem creating the vector constraint using Eq(): %v", err)
	}

	n_R := 2
	for rowIndex := 0; rowIndex < n_R; rowIndex++ {
		lhsConstant := constr.LeftHandSide.Constant()
		vleConstant := vle1.Constant()
		if lhsConstant.AtVec(rowIndex) != vleConstant.AtVec(rowIndex) {
			t.Errorf(
				"The constraint's left hand side has constant value %v at index %v; expected %v!",
				lhsConstant.AtVec(rowIndex),
				rowIndex,
				vleConstant.AtVec(rowIndex),
			)
		}
	}

}

/*
TestVectorLinearExpressionTranspose_Eq2
Description:

	Tests whether or not an equality constraint between a bool and a proper vector variable leads to an error.
	Eq comparison between:
	- Vector Linear Expression, and
	- bool
*/
func TestVectorLinearExpressionTranspose_Eq2(t *testing.T) {
	// Constants
	m := optim.NewModel("Eq2")

	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}
	c := optim.ZerosVector(2)
	vle1 := optim.VectorLinearExpressionTranspose{
		vv1,
		optim.Identity(2),
		c,
	}

	badRHS := false

	// Create Constraint
	_, err := vle1.Eq(badRHS)
	if !strings.Contains(err.Error(), fmt.Sprintf("vector linear expression %v with object of type %T is not currently supported.", vle1, badRHS)) {
		t.Errorf(
			"Expected an error containing \"vector linear expression %v with object of type %T is not currently supported\"; instead received %v",
			vle1,
			badRHS,
			err,
		)
	}

}

/*
TestVectorLinearExpressionTranspose_Eq3
Description:

	Tests whether or not an equality constraint between a KVector and a proper vector variable leads to an error.
	Eq comparison between:
	- Vector Linear Expression, and
	- KVector
*/
func TestVectorLinearExpressionTranspose_Eq3(t *testing.T) {
	// Constants
	m := optim.NewModel("Eq3")

	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}
	c := optim.ZerosVector(2)
	vle1 := optim.VectorLinearExpressionTranspose{
		vv1,
		optim.Identity(2),
		c,
	}

	onesVec1 := optim.OnesVector(2)
	onesVec2 := optim.KVector(onesVec1).Transpose()

	// Create Constraint
	vectorConstraint, err := vle1.Eq(onesVec2)
	if err != nil {
		t.Errorf(
			"There was an issue creating a constraint between %v and %v: %v",
			vle1,
			onesVec2,
			err,
		)
	}

	if vectorConstraint.LeftHandSide.Len() != onesVec2.Len() {
		t.Errorf("The length of lhs (%v) and rhs (%v) should be the same!", vle1.Len(), onesVec2.Len())
	}

}

/*
TestVectorLinearExpressionTranspose_Eq4
Description:

	This test will evaluate how well the Eq() method for the vector of linear constraints works.
	Creates a simple two-dimensional constraint.
	Eq comparison between:
	- Vector Linear Expression, and
	- VarVector
*/
func TestVectorLinearExpressionTranspose_Eq4(t *testing.T) {
	m := optim.NewModel("Eq4")
	dimX := 2
	x := m.AddVariableVector(dimX)

	L1 := optim.Identity(dimX)
	c1 := optim.OnesVector(dimX)

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		x, L1, c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// Create equality comparison.
	_, err = ve1.Eq(x.Transpose())
	if err != nil {
		t.Errorf("There was an issue creating the equality constraint")
	}

}

/*
TestVectorLinearExpressionTranspose_Eq5
Description:

	This test will evaluate how well the Eq() method for the vector of linear constraints works.
	Creates a simple two-dimensional constraint.
	Eq comparison between:
	- Vector Linear Expression, and
	- Vector Linear Expression
*/
func TestVectorLinearExpressionTranspose_Eq5(t *testing.T) {
	m := optim.NewModel("Eq5")
	dimX := 2
	x := m.AddVariableVector(dimX)

	L1 := optim.Identity(dimX)
	c1 := optim.OnesVector(dimX)

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		x, L1, c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// Create equality comparison.
	_, err = ve1.Eq(ve1)
	if err != nil {
		t.Errorf("There was an issue creating the equality constraint")
	}

}

/*
TestVectorLinearExpressionTranspose_Eq16
Description:

	Tests whether or not an equality constraint between a ones vector and a standard vector variable works well.
	Eq comparison between:
	- Vector Linear Expression, and
	- mat.VecDense
*/
func TestVectorLinearExpressionTranspose_Eq6(t *testing.T) {
	// Constants
	m := optim.NewModel("Eq6")

	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}
	c := optim.ZerosVector(2)
	vle1 := optim.VectorLinearExpressionTranspose{
		vv1,
		optim.Identity(2),
		c,
	}

	ones0 := optim.OnesVector(2)

	// Create Constraint
	_, err := vle1.Eq(ones0)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot compare VectorLinearExpressionTranspose with a normal vector %v (%T); Try transposing one or the other!",
			ones0, ones0,
		),
	) {
		t.Errorf("The wrong error was thrown while using Eq(): %v", err)
	}

}

/*
TestVectorLinearExpressionTranspose_Len1
Description:

	This test will evaluate how well the Len() method for the vector of linear constraints works.
	A constraint between two vectors of length 2
*/
func TestVectorLinearExpressionTranspose_Len1(t *testing.T) {
	m := optim.NewModel("Len1")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	L1 := *mat.NewDense(2, 2, []float64{1.0, 2.0, 3.0, 4.0})
	c1 := *mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		vv1, L1, c1,
	}

	// ve1 should pass all checks.
	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was supposed to be valid, but received an error: %v", err)
	}

	if ve1.Len() != 2 {
		t.Errorf("Len() of vector linear expression was %v; expeted 2", ve1.Len())
	}
}

/*
TestVectorLinearExpressionTranspose_Len2
Description:

	This test will evaluate how well the Len() method for the vector of linear constraints works.
	A constraint between two vectors of length 10
*/
func TestVectorLinearExpressionTranspose_Len2(t *testing.T) {
	m := optim.NewModel("Len2")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y, x, y, x, y, x, y, x, y},
	}

	dimX := 10
	L1 := optim.Identity(dimX)
	c1 := optim.OnesVector(dimX)

	// Use these to create expression.
	ve1 := optim.VectorLinearExpressionTranspose{
		vv1, L1, c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// ve1 should pass all checks.
	if ve1.Len() != dimX {
		t.Errorf("Len() of vector linear expression was %v; expeted %v", ve1.Len(), dimX)
	}

}

/*
TestVectorLinearExpressionTranspose_Plus1
Description:

	Add VectorLinearExpressionTranspose to a KVector of appropriate length.
*/
func TestVectorLinearExpressionTranspose_Plus1(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("Plus1")

	kv1 := optim.KVectorTranspose(
		optim.OnesVector(n),
	)
	vle2 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}

	// Compute Sum
	tempSum, err := vle2.Plus(kv1)
	if err != nil {
		t.Errorf("There was an issue computing this good addition: %v", err)
	}

	sumAsVLE, ok := tempSum.(optim.VectorLinearExpressionTranspose)
	if !ok {
		t.Errorf("Expecting sum to be of type VectorLinearExpressionTranspose; received %T", tempSum)
	}

	// Verify the values of C
	for dimIndex := 0; dimIndex < n; dimIndex++ {
		if float64(kv1.AtVec(dimIndex).(optim.K)) != sumAsVLE.C.AtVec(dimIndex) {
			t.Errorf("kv1[%v] = %v != %v = sumAsVLE.C[%v]",
				dimIndex,
				kv1.AtVec(dimIndex),
				sumAsVLE.C.AtVec(dimIndex),
				dimIndex,
			)
		}
	}

	// Verify the values of L
	nR, nC := sumAsVLE.L.Dims()
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			if rowIndex == colIndex {
				if sumAsVLE.L.At(rowIndex, colIndex) != 1.0 {
					t.Errorf(
						"Expected L[%v,%v] = 1.0; received %v",
						rowIndex, colIndex,
						sumAsVLE.L.At(rowIndex, colIndex),
					)
				}
			} else {
				if sumAsVLE.L.At(rowIndex, colIndex) != 0.0 {
					t.Errorf(
						"Expected L[%v,%v] = 0.0; received %v",
						rowIndex, colIndex,
						sumAsVLE.L.At(rowIndex, colIndex),
					)
				}
			}
		}
	}
}

/*
TestVectorLinearExpressionTranspose_Plus2
Description:

	Add VectorLinearExpressionTranspose to a KVector of inappropriate length.
*/
func TestVectorLinearExpressionTranspose_Plus2(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("Plus2")

	kv1 := optim.KVector(
		optim.OnesVector(n + 1),
	).Transpose()
	vle2 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}

	// Compute Sum
	_, err := vle2.Plus(kv1)
	if err == nil {
		t.Errorf("There should have been an issue adding together these two vector expressions of different dimension, but none was received!")
	}

	if !strings.Contains(err.Error(), fmt.Sprintf("The length of input KVector (%v) did not match the length of the VectorLinearExpressionTranspose (%v).", kv1.Len(), vle2.Len())) {
		t.Errorf("Unexpected error: %v", err)
	}

}

/*
TestVectorLinearExpressionTranspose_Plus3
Description:

	Add VectorLinearExpressionTranspose to a KVector of appropriate length.
	Nonzero offset in VectorLinearExpressionTranspose.
*/
func TestVectorLinearExpressionTranspose_Plus3(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLETranspose-Plus3")

	kv1 := optim.KVectorTranspose(
		optim.OnesVector(n),
	)
	vle2 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.OnesVector(n),
	}

	// Compute Sum
	tempSum, err := vle2.Plus(kv1)
	if err != nil {
		t.Errorf("There was an issue computing this good addition: %v", err)
	}

	sumAsVLE, ok := tempSum.(optim.VectorLinearExpressionTranspose)
	if !ok {
		t.Errorf("Expecting sum to be of type VectorLinearExpressionTranspose; received %T", tempSum)
	}

	// Verify the values of C
	for dimIndex := 0; dimIndex < n; dimIndex++ {
		if float64(kv1.AtVec(dimIndex).(optim.K))+1.0 != sumAsVLE.C.AtVec(dimIndex) {
			t.Errorf("kv1[%v] + 1.0 = %v != %v = sumAsVLE.C[%v]",
				dimIndex,
				float64(kv1.AtVec(dimIndex).(optim.K))+1.0,
				sumAsVLE.C.AtVec(dimIndex),
				dimIndex,
			)
		}
	}

	// Verify the values of L
	nR, nC := sumAsVLE.L.Dims()
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			if rowIndex == colIndex {
				if sumAsVLE.L.At(rowIndex, colIndex) != 1.0 {
					t.Errorf(
						"Expected L[%v,%v] = 1.0; received %v",
						rowIndex, colIndex,
						sumAsVLE.L.At(rowIndex, colIndex),
					)
				}
			} else {
				if sumAsVLE.L.At(rowIndex, colIndex) != 0.0 {
					t.Errorf(
						"Expected L[%v,%v] = 0.0; received %v",
						rowIndex, colIndex,
						sumAsVLE.L.At(rowIndex, colIndex),
					)
				}
			}
		}
	}
}

/*
TestVectorLinearExpressionTranspose_Plus4
Description:

	Add VectorLinearExpressionTranspose to a VarVector of appropriate length.
*/
func TestVectorLinearExpressionTranspose_Plus4(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLETranspose-Plus4")

	vv1 := m.AddVariableVector(n).Transpose()
	vle2 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.OnesVector(n),
	}

	// Compute Sum
	tempSum, err := vle2.Plus(vv1)
	if err != nil {
		t.Errorf("There was an issue computing this good addition: %v", err)
	}

	sumAsVLE, ok := tempSum.(optim.VectorLinearExpressionTranspose)
	if !ok {
		t.Errorf("Expecting sum to be of type VectorLinearExpressionTranspose; received %T", tempSum)
	}

	// Verify the values of C
	for dimIndex := 0; dimIndex < n; dimIndex++ {
		if vle2.C.AtVec(dimIndex) != sumAsVLE.C.AtVec(dimIndex) {
			t.Errorf("kv1[%v] = %v != %v = sumAsVLE.C[%v]",
				dimIndex,
				vle2.C.AtVec(dimIndex),
				sumAsVLE.C.AtVec(dimIndex),
				dimIndex,
			)
		}
	}

	// Verify the values of L
	nR, nC := sumAsVLE.L.Dims()
	if nC != 2*n {
		t.Errorf("Expected for the number of columns in sum.L (%v) to match 2*n = %v.", nC, 2*n)
	}

	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			switch {
			case rowIndex == colIndex:
				if sumAsVLE.L.At(rowIndex, colIndex) != 1.0 {
					t.Errorf(
						"Expected L[%v,%v] = 1.0; received %v",
						rowIndex, colIndex,
						sumAsVLE.L.At(rowIndex, colIndex),
					)
				}
			case rowIndex+n == colIndex:
				if sumAsVLE.L.At(rowIndex, colIndex) != 1.0 {
					t.Errorf(
						"Expected L[%v,%v] = 1.0; received %v",
						rowIndex, colIndex,
						sumAsVLE.L.At(rowIndex, colIndex),
					)
				}
			default:
				if sumAsVLE.L.At(rowIndex, colIndex) != 0.0 {
					t.Errorf(
						"Expected L[%v,%v] = 0.0; received %v",
						rowIndex, colIndex,
						sumAsVLE.L.At(rowIndex, colIndex),
					)
				}
			}
		}
	}
}

/*
TestVectorLinearExpressionTranspose_Plus5
Description:

	Add VectorLinearExpressionTranspose to a VarVector of appropriate length.
*/
func TestVectorLinearExpressionTranspose_Plus5(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("Plus5")

	vv1 := m.AddVariableVector(n)
	vle1 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: vv1,
		C: optim.ZerosVector(n),
	}
	vle2 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: vv1,
		C: optim.OnesVector(n),
	}

	// Compute Sum
	tempSum, err := vle2.Plus(vle1)
	if err != nil {
		t.Errorf("There was an issue computing this good addition: %v", err)
	}

	sumAsVLE, ok := tempSum.(optim.VectorLinearExpressionTranspose)
	if !ok {
		t.Errorf("Expecting sum to be of type VectorLinearExpressionTranspose; received %T", tempSum)
	}

	// Verify the values of C
	for dimIndex := 0; dimIndex < n; dimIndex++ {
		if vle2.C.AtVec(dimIndex)+vle1.C.AtVec(dimIndex) != sumAsVLE.C.AtVec(dimIndex) {
			t.Errorf("vle1[%v] + vle2[%v] = %v != %v = sumAsVLE.C[%v]",
				dimIndex, dimIndex,
				vle2.C.AtVec(dimIndex)+vle1.C.AtVec(dimIndex),
				sumAsVLE.C.AtVec(dimIndex),
				dimIndex,
			)
		}
	}

	// Verify the values of L
	nR, nC := sumAsVLE.L.Dims()
	if nC != n {
		t.Errorf("Expected for the number of columns in sum.L (%v) to match 2*n = %v.", nC, 2*n)
	}

	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			switch {
			case rowIndex == colIndex:
				if sumAsVLE.L.At(rowIndex, colIndex) != 2.0 {
					t.Errorf(
						"Expected L[%v,%v] = 1.0; received %v",
						rowIndex, colIndex,
						sumAsVLE.L.At(rowIndex, colIndex),
					)
				}
			default:
				if sumAsVLE.L.At(rowIndex, colIndex) != 0.0 {
					t.Errorf(
						"Expected L[%v,%v] = 0.0; received %v",
						rowIndex, colIndex,
						sumAsVLE.L.At(rowIndex, colIndex),
					)
				}
			}
		}
	}
}

/*
TestVectorLinearExpressionTranspose_Plus6
Description:

	Add VectorLinearExpressionTranspose to a VectorLinearExpressionTranspose of appropriate length. (But different variables)
*/
func TestVectorLinearExpressionTranspose_Plus6(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("Plus6")

	vv1 := m.AddVariableVector(n)
	vle1 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: vv1,
		C: optim.ZerosVector(n),
	}
	vle2 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.OnesVector(n),
	}

	// Compute Sum
	tempSum, err := vle2.Plus(vle1)
	if err != nil {
		t.Errorf("There was an issue computing this good addition: %v", err)
	}

	sumAsVLE, ok := tempSum.(optim.VectorLinearExpressionTranspose)
	if !ok {
		t.Errorf("Expecting sum to be of type VectorLinearExpressionTranspose; received %T", tempSum)
	}

	// Verify the values of C
	for dimIndex := 0; dimIndex < n; dimIndex++ {
		if vle2.C.AtVec(dimIndex)+vle1.C.AtVec(dimIndex) != sumAsVLE.C.AtVec(dimIndex) {
			t.Errorf("vle1[%v] + vle2[%v] = %v != %v = sumAsVLE.C[%v]",
				dimIndex, dimIndex,
				vle2.C.AtVec(dimIndex)+vle1.C.AtVec(dimIndex),
				sumAsVLE.C.AtVec(dimIndex),
				dimIndex,
			)
		}
	}

	// Verify the values of L
	nR, nC := sumAsVLE.L.Dims()
	if nC != 2*n {
		t.Errorf("Expected for the number of columns in sum.L (%v) to match 2*n = %v.", nC, 2*n)
	}

	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			switch {
			case rowIndex == colIndex:
				if sumAsVLE.L.At(rowIndex, colIndex) != 1.0 {
					t.Errorf(
						"Expected L[%v,%v] = 1.0; received %v",
						rowIndex, colIndex,
						sumAsVLE.L.At(rowIndex, colIndex),
					)
				}
			case rowIndex+n == colIndex:
				if sumAsVLE.L.At(rowIndex, colIndex) != 1.0 {
					t.Errorf(
						"Expected L[%v,%v] = 1.0; received %v",
						rowIndex, colIndex,
						sumAsVLE.L.At(rowIndex, colIndex),
					)
				}
			default:
				if sumAsVLE.L.At(rowIndex, colIndex) != 0.0 {
					t.Errorf(
						"Expected L[%v,%v] = 0.0; received %v",
						rowIndex, colIndex,
						sumAsVLE.L.At(rowIndex, colIndex),
					)
				}
			}
		}
	}
}

/*
TestVectorLinearExpressionTranspose_Plus7
Description:

	Add VectorLinearExpressionTranspose to a KVector.
*/
func TestVectorLinearExpressionTranspose_Plus7(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLET Plus7")

	kv1 := optim.KVector(
		optim.OnesVector(n),
	)
	vle2 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}

	// Compute Sum
	_, err := vle2.Plus(kv1)
	if err == nil {
		t.Errorf("There should have been an issue adding together these two vector expressions of different dimension, but none was received!")
	}

	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot add VectorLinearExpressioinTranspose with a normal vector %v (%T); Try transposing one or the other!",
			kv1, kv1,
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}

}

/*
TestVectorLinearExpressionTranspose_Plus8
Description:

	Add VectorLinearExpressionTranspose to a VarVector.
*/
func TestVectorLinearExpressionTranspose_Plus8(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLET Plus7")

	vle2 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}

	// Compute Sum
	_, err := vle2.Plus(vle2.X)
	if err == nil {
		t.Errorf("There should have been an issue adding together these two vector expressions of different dimension, but none was received!")
	}

	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot add VectorLinearExpressioinTranspose with a normal vector %v (%T); Try transposing one or the other!",
			vle2.X, vle2.X,
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}

}

/*
TestVectorLinearExpressionTranspose_Plus9
Description:

	Add VectorLinearExpressionTranspose to a VectorLinearExpression.
*/
func TestVectorLinearExpressionTranspose_Plus9(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLET Plus9")

	vle2 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}

	// Compute Sum
	_, err := vle2.Plus(true)
	if err == nil {
		t.Errorf("There should have been an issue adding together these two vector expressions of different dimension, but none was received!")
	}

	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"The VectorLinearExpressionTranspose.Plus method has not yet been implemented for type %T!",
			true,
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}

}

/*
TestVectorLinearExpressionTranspose_Plus10
Description:

	Add VectorLinearExpressionTranspose to a VectorLinearExpression.
*/
func TestVectorLinearExpressionTranspose_Plus10(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLET Plus10")

	vle2 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}

	// Compute Sum
	_, err := vle2.Plus(vle2.Transpose())
	if err == nil {
		t.Errorf("There should have been an issue adding together these two vector expressions of different dimension, but none was received!")
	}

	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot add VectorLinearExpressioinTranspose with a normal vector %v (%T); Try transposing one or the other!",
			vle2.Transpose(), vle2.Transpose(),
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}

}

/*
TestVectorLinearExpressionTranspose_Plus11
Description:

	Add VectorLinearExpressionTranspose to a VectorLinearExpressionTranspose
	that is not of the right size.
*/
func TestVectorLinearExpressionTranspose_Plus11(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLET Plus11")

	vle1 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n + 1),
		X: m.AddVariableVector(n + 1),
		C: optim.ZerosVector(n + 1),
	}

	vle2 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}

	// Compute Sum
	_, err := vle1.Plus(vle2)
	if err == nil {
		t.Errorf("There should have been an issue adding together these two vector expressions of different dimension, but none was received!")
	}

	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"The length of input VectorLinearExpressionTranspose (%v) did not match the length of the VectorLinearExpressionTranspose (%v).",
			vle2.Len(),
			vle1.Len(),
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}

}

/*
TestVectorLinearExpressionTranspose_Comparison1
Description:

	Compares a VectorLinearExpressionTranspose with a constant vector.
*/
func TestVectorLinearExpressionTranspose_Comparison1(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLET Comparison1")

	vle1 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}
	kv1 := optim.KVector(optim.OnesVector(n))

	// Algorithm
	_, err := vle1.Comparison(kv1, optim.SenseLessThanEqual)
	if err == nil {
		t.Errorf("Expected for there to be an error, but there was none.")
	}

	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot compare VectorLinearExpressionTranspose with a normal vector %v (%T); Try transposing one or the other!",
			kv1, kv1,
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}
}

/*
TestVectorLinearExpressionTranspose_Comparison2
Description:

	Compares a VectorLinearExpressionTranspose with a constant vector transpose.
*/
func TestVectorLinearExpressionTranspose_Comparison2(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLET Comparison2")

	vle1 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}
	kv1 := optim.KVectorTranspose(optim.OnesVector(n + 1))

	// Algorithm
	_, err := vle1.Comparison(kv1, optim.SenseLessThanEqual)
	if err == nil {
		t.Errorf("Expected for there to be an error, but there was none.")
	}

	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"The two vector inputs to Eq() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
			vle1.Len(),
			kv1.Len(),
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}
}

/*
TestVectorLinearExpressionTranspose_Comparison3
Description:

	Compares a VectorLinearExpressionTranspose with a var vector.
*/
func TestVectorLinearExpressionTranspose_Comparison3(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLET Comparison1")

	vle1 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}
	X1 := m.AddVariableVector(n)

	// Algorithm
	_, err := vle1.Comparison(X1, optim.SenseLessThanEqual)
	if err == nil {
		t.Errorf("Expected for there to be an error, but there was none.")
	}

	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot compare VectorLinearExpressionTranspose with a normal vector %v (%T); Try transposing one or the other!",
			X1, X1,
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}
}

/*
TestVectorLinearExpressionTranspose_Comparison4
Description:

	Compares a VectorLinearExpressionTranspose with a varVectorTranspose
	of the wrong shape.
*/
func TestVectorLinearExpressionTranspose_Comparison4(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLET Comparison4")

	vle1 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}
	X1 := m.AddVariableVector(n + 1)

	// Algorithm
	_, err := vle1.Comparison(X1.Transpose(), optim.SenseLessThanEqual)
	if err == nil {
		t.Errorf("Expected for there to be an error, but there was none.")
	}

	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"The two vector inputs to Eq() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
			vle1.Len(),
			X1.Len(),
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}
}

/*
TestVectorLinearExpressionTranspose_Comparison5
Description:

	Compares a VectorLinearExpressionTranspose with a vector linear expression.
*/
func TestVectorLinearExpressionTranspose_Comparison5(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLET Comparison5")

	vle1 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}

	// Algorithm
	_, err := vle1.Comparison(vle1.Transpose(), optim.SenseLessThanEqual)
	if err == nil {
		t.Errorf("Expected for there to be an error, but there was none.")
	}

	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"Cannot compare VectorLinearExpressionTranspose with a normal vector %v (%T); Try transposing one or the other!",
			vle1.Transpose(), vle1.Transpose(),
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}
}

/*
TestVectorLinearExpressionTranspose_Comparison6
Description:

	Compares a VectorLinearExpressionTranspose with a vector linear expression transpose of
	the wrong shape.
*/
func TestVectorLinearExpressionTranspose_Comparison6(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLET Comparison6")

	vle1 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}

	vle2 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n + 1),
		X: m.AddVariableVector(n + 1),
		C: optim.ZerosVector(n + 1),
	}

	// Algorithm
	_, err := vle1.Comparison(vle2, optim.SenseLessThanEqual)
	if err == nil {
		t.Errorf("Expected for there to be an error, but there was none.")
	}

	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"The two vector inputs to Eq() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
			vle1.Len(),
			vle2.Len(),
		),
	) {
		t.Errorf("Unexpected error: %v", err)
	}
}

/*
TestVectorLinearExpressionTranspose_AtVec1
Description:

	Tests to make sure that AtVec properly works.
*/
func TestVectorLinearExpressionTranspose_AtVec1(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLET AtVec1")

	vle1 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}

	// Expect
	sle1 := vle1.AtVec(n - 1)
	if _, ok := sle1.(optim.ScalarLinearExpr); !ok {
		t.Errorf("Expected for vle1[n-1] to be a ScalarLinearExpr; received %T", sle1)
	}

}

/*
TestVectorLinearExpressionTranspose_Transpose1
Description:

	Tests to make sure that AtVec properly works.
*/
func TestVectorLinearExpressionTranspose_Transpose1(t *testing.T) {
	// Constants
	n := 5
	m := optim.NewModel("VLET Transpose1")

	vle1 := optim.VectorLinearExpressionTranspose{
		L: optim.Identity(n),
		X: m.AddVariableVector(n),
		C: optim.ZerosVector(n),
	}

	// Expect
	vle2 := vle1.Transpose()
	if _, ok := vle2.(optim.VectorLinearExpr); !ok {
		t.Errorf("Expected for vle1.Transpose() to be a VectorLinearExpr; received %T", vle2)
	}

}
