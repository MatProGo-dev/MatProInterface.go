package testing

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"testing"
)

func TestSumVars(t *testing.T) {
	numVars := 3
	m := optim.NewModel("TestSumVars")
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()
	z := m.AddBinaryVariable()
	expr := optim.SumVars(x, y, z)

	// t.Errorf("%v", expr.(optim.ScalarLinearExpr))

	for _, coeff := range expr.Coeffs() {
		if coeff != 1 {
			t.Errorf("Coeff mismatch: %v != 1", coeff)
		}
	}

	if expr.NumVars() != numVars {
		t.Errorf("NumVars mismatch: %v != %v", expr.NumVars(), numVars)
	}

	if expr.Constant() != 0 {
		t.Errorf("Constant mismatch: %v != 0", expr.Constant())
	}
}

/*
TestUtil_Identity1
Description:

	Create identity matrix of dimension 1 (scalar?).
*/
func TestUtil_Identity1(t *testing.T) {
	// Constants
	n := 1
	// Algorithm
	identMat1 := optim.Identity(n)

	nX, nY := identMat1.Dims()
	if (nX != n) || (nY != n) {
		t.Errorf("The identity matrix created has dimension %v x %v; Expected %v x %v.",
			nX, nY,
			n, n,
		)
	}

}

/*
TestUtil_Identity2
Description:

	Create identity matrix of dimension 10.
*/
func TestUtil_Identity2(t *testing.T) {
	// Constants
	n := 10
	// Algorithm
	identMat1 := optim.Identity(n)

	nX, nY := identMat1.Dims()
	if (nX != n) || (nY != n) {
		t.Errorf("The identity matrix created has dimension %v x %v; Expected %v x %v.",
			nX, nY,
			n, n,
		)
	}

}

/*
TestUtil_FindInSlice1
Description:

	Tests the find in slice function for strings!
	(string in slice)
*/
func TestUtil_FindInSlice1(t *testing.T) {
	// Constant
	slice0 := []string{"Why", "test", "this", "?"}

	// Find!
	foundIndex, err := optim.FindInSlice("?", slice0)
	if err != nil {
		t.Errorf("Received an error while trying to find valid string: %v", err)
	}

	if foundIndex != 3 {
		t.Errorf(
			"Expected index of string to be 3; received %v.", foundIndex,
		)
	}
}

/*
TestUtil_FindInSlice2
Description:

	Tests the find in slice function for strings!
	(string NOT in slice)
*/
func TestUtil_FindInSlice2(t *testing.T) {
	// Constant
	slice0 := []string{"Why", "test", "this", "?"}

	// Find!
	foundIndex, err := optim.FindInSlice("llama", slice0)
	if err != nil {
		t.Errorf("Received an error while trying to find valid string: %v", err)
	}

	if foundIndex != -1 {
		t.Errorf(
			"Expected index of string to be -1; received %v.", foundIndex,
		)
	}
}

/*
TestUtil_FindInSlice3
Description:

	Tests the find in slice function for ints!
	(int in slice)
*/
func TestUtil_FindInSlice3(t *testing.T) {
	// Constant
	slice0 := []int{1, 3, 7, 11}

	// Find!
	foundIndex, err := optim.FindInSlice(1, slice0)
	if err != nil {
		t.Errorf("Received an error while trying to find valid string: %v", err)
	}

	if foundIndex != 0 {
		t.Errorf(
			"Expected index of string to be 0; received %v.", foundIndex,
		)
	}
}

/*
TestUtil_FindInSlice4
Description:

	Tests the find in slice function for ints!
	(int NOT in slice)
*/
func TestUtil_FindInSlice4(t *testing.T) {
	// Constant
	slice0 := []int{1, 3, 7, 11}

	// Find!
	foundIndex, err := optim.FindInSlice(21, slice0)
	if err != nil {
		t.Errorf("Received an error while trying to find valid string: %v", err)
	}

	if foundIndex != -1 {
		t.Errorf(
			"Expected index of string to be -1; received %v.", foundIndex,
		)
	}
}

/*
TestUtil_FindInSlice5
Description:

	Tests the find in slice function for uint64!
	(uint64 in slice)
*/
func TestUtil_FindInSlice5(t *testing.T) {
	// Constant
	slice0 := []uint64{1, 3, 7, 11}
	var x uint64 = 1

	// Find!
	foundIndex, err := optim.FindInSlice(x, slice0)
	if err != nil {
		t.Errorf("Received an error while trying to find valid string: %v", err)
	}

	if foundIndex != 0 {
		t.Errorf(
			"Expected index of string to be 0; received %v.", foundIndex,
		)
	}
}

/*
TestUtil_FindInSlice6
Description:

	Tests the find in slice function for uint64!
	(uint64 NOT in slice)
*/
func TestUtil_FindInSlice6(t *testing.T) {
	// Constant
	slice0 := []uint64{1, 3, 7, 11}
	var x uint64 = 21

	// Find!
	foundIndex, err := optim.FindInSlice(x, slice0)
	if err != nil {
		t.Errorf("Received an error while trying to find valid string: %v", err)
	}

	if foundIndex != -1 {
		t.Errorf(
			"Expected index of string to be -1; received %v.", foundIndex,
		)
	}
}

/*
TestUtil_FindInSlice7
Description:

	Tests the find in slice function for strings!
	(Variable in slice)
*/
func TestUtil_FindInSlice7(t *testing.T) {
	// Constant
	m := optim.NewModel("test-util-findinslice7")
	vv1 := m.AddVariableVector(100)
	slice0 := vv1.Elements

	// Find!
	indexUnderTest := 2
	foundIndex, err := optim.FindInSlice(vv1.Elements[indexUnderTest], slice0)
	if err != nil {
		t.Errorf("Received an error while trying to find valid string: %v", err)
	}

	if foundIndex != indexUnderTest {
		t.Errorf(
			"Expected index of string to be %v; received %v.",
			indexUnderTest, foundIndex,
		)
	}
}

/*
TestUtil_FindInSlice4
Description:

	Tests the find in slice function for strings!
	(int NOT in slice)
*/
func TestUtil_FindInSlice8(t *testing.T) {
	// Constant
	m := optim.NewModel("test-util-findinslice7")
	vv1 := m.AddVariableVector(100)
	slice0 := vv1.Elements

	// Find!
	newVar := m.AddVariable()
	foundIndex, err := optim.FindInSlice(newVar, slice0)
	if err != nil {
		t.Errorf("Received an error while trying to find valid string: %v", err)
	}

	if foundIndex != -1 {
		t.Errorf(
			"Expected index of string to be %v; received %v.",
			-1, foundIndex,
		)
	}
}
