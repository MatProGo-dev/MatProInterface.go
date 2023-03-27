package optim_test

/*
util_test.go
Description:
	This file tests some of the utilities added in MatProInterface.go's util.go file.
*/

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
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
