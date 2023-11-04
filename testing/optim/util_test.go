package optim_test

/*
util_test.go
Description:
	This file tests some of the utilities added in MatProInterface.go's util.go file.
*/

import (
	"fmt"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"strings"
	"testing"
)

/*
TestUtil_OnesVector1
Description:

	Tests that the OnesVector() function works well with a large input size.
*/
func TestUtil_OnesVector1(t *testing.T) {
	// Constants
	length1 := 10

	// Algorithm
	ones1 := optim.OnesVector(length1)
	if ones1.Len() != length1 {
		t.Errorf("Attempted to create ones vector of length %v; received vector of length %v", length1, ones1.Len())
	}

	// Check each element in ones
	for eltIndex := 0; eltIndex < ones1.Len(); eltIndex++ {
		if ones1.AtVec(eltIndex) != 1.0 {
			t.Errorf("Element at index %v of ones1 has value %v; not 1.0", eltIndex, ones1.AtVec(eltIndex))
		}
	}
}

/*
TestUtil_OnesVector2
Description:

	Tests that the OnesVector() function works well with an input size of 1.
*/
func TestUtil_OnesVector2(t *testing.T) {
	// Constants
	length1 := 1

	// Algorithm
	ones1 := optim.OnesVector(length1)
	if ones1.Len() != length1 {
		t.Errorf("Attempted to create ones vector of length %v; received vector of length %v", length1, ones1.Len())
	}

	// Check each element in ones
	for eltIndex := 0; eltIndex < ones1.Len(); eltIndex++ {
		if ones1.AtVec(eltIndex) != 1.0 {
			t.Errorf("Element at index %v of ones1 has value %v; not 1.0", eltIndex, ones1.AtVec(eltIndex))
		}
	}
}

/*
TestUtil_CheckExtras1
Description:

	Tests the CheckExtras function can properly handle cases where extras has nil.
*/
func TestUtil_CheckExtras1(t *testing.T) {
	// Constants
	extras := []interface{}{}

	// Algorithm
	err := optim.CheckExtras(extras)
	if err != nil {
		t.Errorf(
			"When no extras were provided, there should be no errors; instead received %v",
			err,
		)
	}
}

/*
TestUtil_CheckExtras2
Description:

	Tests the CheckExtras function can properly handle cases where extras
	is simply a nil.
*/
func TestUtil_CheckExtras2(t *testing.T) {
	// Constants
	extras := []interface{}{
		nil,
	}

	// Algorithm
	err := optim.CheckExtras(extras)
	if err != nil {
		t.Errorf(
			"When a single nil extras was provided, there should be no errors; instead received %v",
			err,
		)
	}
}

/*
TestUtil_CheckExtras3
Description:

	Tests the CheckExtras function can properly handle cases where extras has
	one element but it is not an error.
*/
func TestUtil_CheckExtras3(t *testing.T) {
	// Constants
	extras := []interface{}{
		false,
	}

	// Algorithm
	err := optim.CheckExtras(extras)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"unexpected type of input as an 'extra': %T",
			false,
		),
	) {
		t.Errorf(
			"unexpected error: %v",
			err,
		)
	}
}

/*
TestUtil_CheckExtras4
Description:

	Tests the CheckExtras function can properly handle cases where extras has
	one element but it is a proper error.
*/
func TestUtil_CheckExtras4(t *testing.T) {
	// Constants
	err0 := fmt.Errorf("test")
	extras := []interface{}{
		err0,
	}

	// Algorithm
	err := optim.CheckExtras(extras)
	if !strings.Contains(
		err.Error(),
		err0.Error(),
	) {
		t.Errorf(
			"unexpected error: %v",
			err,
		)
	}
}

/*
TestUtil_CheckExtras5
Description:

	Tests the CheckExtras function can properly handle cases where extras has
	multiple elements but it is a proper error.
*/
func TestUtil_CheckExtras5(t *testing.T) {
	// Constants
	err0 := fmt.Errorf("test")
	extras := []interface{}{
		err0, false,
	}

	// Algorithm
	err := optim.CheckExtras(extras)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"did not expect to receive more than one element in 'extras' input; received %v",
			len(extras),
		),
	) {
		t.Errorf(
			"unexpected error: %v",
			err,
		)
	}
}

/*
TestUtils_SumVars1
Description:

	This function tests the rewritten version of SumVars.
*/
func TestUtils_SumRow1(t *testing.T) {
	// Constants
	N := 9
	m := optim.NewModel("SumRow1")
	vv1 := m.AddVariableVector(N)

	// Create matrix
	fakeMatrix := [][]optim.Variable{
		[]optim.Variable{vv1.Elements[0], vv1.Elements[1], vv1.Elements[2]},
		[]optim.Variable{vv1.Elements[3], vv1.Elements[4], vv1.Elements[5]},
		[]optim.Variable{vv1.Elements[6], vv1.Elements[7], vv1.Elements[8]},
	}

	// Add Sum
	sum := optim.SumRow(fakeMatrix, 1)

	sumAsSLE, ok1 := sum.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf(
			"Expected sum to be ScalarLinearExpr; received %T",
			sum,
		)
	}

	if sumAsSLE.X.Len() != 3 {
		t.Errorf(
			"Expected sum to have 3 elements; it contained %v",
			sumAsSLE.X.Len(),
		)
	}

	expectedVarIndices := []uint64{3, 4, 5}
	for colIndex := 0; colIndex < 3; colIndex++ {
		X_i := sumAsSLE.X.AtVec(colIndex).(optim.Variable)
		if foundIndex, _ := optim.FindInSlice(X_i.ID, expectedVarIndices); foundIndex == -1 {
			t.Errorf(
				"Expected X_i ID %v to be in column %v. It was not!",
				X_i.ID,
				1,
			)
		}
	}
}

/*
TestUtils_SumCol1
Description:

	This function tests the rewritten version of SumCol.
*/
func TestUtils_SumCol1(t *testing.T) {
	// Constants
	N := 9
	m := optim.NewModel("SumRow1")
	vv1 := m.AddVariableVector(N)

	// Create matrix
	fakeMatrix := [][]optim.Variable{
		[]optim.Variable{vv1.Elements[0], vv1.Elements[1], vv1.Elements[2]},
		[]optim.Variable{vv1.Elements[3], vv1.Elements[4], vv1.Elements[5]},
		[]optim.Variable{vv1.Elements[6], vv1.Elements[7], vv1.Elements[8]},
	}

	// Add Sum
	sum := optim.SumCol(fakeMatrix, 1)

	sumAsSLE, ok1 := sum.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf(
			"Expected sum to be ScalarLinearExpr; received %T",
			sum,
		)
	}

	if sumAsSLE.X.Len() != 3 {
		t.Errorf(
			"Expected sum to have 3 elements; it contained %v",
			sumAsSLE.X.Len(),
		)
	}

	expectedVarIndices := []uint64{1, 4, 7}
	for rowIndex := 0; rowIndex < 3; rowIndex++ {
		X_i := sumAsSLE.X.AtVec(rowIndex).(optim.Variable)
		if foundIndex, _ := optim.FindInSlice(X_i.ID, expectedVarIndices); foundIndex == -1 {
			t.Errorf(
				"Expected X_i ID %v to be in column %v. It was not!",
				X_i.ID,
				1,
			)
		}
	}
}

/*
TestUtils_FindInSlice1
Description:

	Tests whether or not FindInSlice collectly handles erorrs when a bad slice is given.
*/
func TestUtils_FindInSlice1(t *testing.T) {
	// Constants
	m := optim.NewModel("FindInSlice1")
	v1 := m.AddVariable()
	slice1 := []int{1, 2, 3}

	// Find
	_, err := optim.FindInSlice(v1, slice1)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"the input slice is of type %T, but the element we're searching for is of type %T",
			slice1,
			v1,
		),
	) {
		t.Errorf(
			"Unexpected error: %v",
			err,
		)
	}
}

/*
TestUtils_FindInSlice2
Description:

	Tests whether or not FindInSlice collectly handles erorrs when an unknown type is given.
*/
func TestUtils_FindInSlice2(t *testing.T) {
	// Constants
	b1 := false
	slice1 := []bool{false, true, false}

	allowedTypes := []string{"string", "int", "uint64", "Variable"}

	// Find
	_, err := optim.FindInSlice(b1, slice1)
	if !strings.Contains(
		err.Error(),
		fmt.Sprintf(
			"the FindInSlice() function was only defined for types %v, not type %T:",
			allowedTypes,
			b1,
		),
	) {
		t.Errorf(
			"Unexpected error: %v",
			err,
		)
	}
}
