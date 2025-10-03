package problem_test

import (
	"fmt"
	"testing"

	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"github.com/MatProGo-dev/MatProInterface.go/problem"
)

/*
objective_sense_test.go
Description:

	Tests for the objective sense object.
*/

/*
TestObjectiveSense_ToObjSense1
Description:

	Tests the ToObjSense function with a minimization sense.
*/
func TestObjectiveSense_ToObjSense1(t *testing.T) {
	// Constants
	sense := optim.SenseMinimize

	// Algorithm
	objSense := problem.ToObjSense(sense)

	// Check that the objective sense is as expected.
	if objSense != problem.SenseMinimize {
		t.Errorf("expected the objective sense to be %v; received %v",
			problem.SenseMinimize, objSense)
	}
}

/*
TestObjectiveSense_ToObjSense2
Description:

	Tests the ToObjSense function with a maximization sense.
*/
func TestObjectiveSense_ToObjSense2(t *testing.T) {
	// Constants
	var sense optim.ObjSense = optim.SenseMaximize

	// Algorithm
	objSense := problem.ToObjSense(sense)

	// Check that the objective sense is as expected.
	if objSense != problem.SenseMaximize {
		t.Errorf("expected the objective sense to be %v; received %v",
			problem.SenseMaximize, objSense)
	}
}

/*
TestObjectiveSense_String1
Description:

	Tests that we can extract strings from the three normal ObjSense values
	(Minimize, Maximize, Find).
*/
func TestObjectiveSense_String1(t *testing.T) {
	// Test Minimize
	minSense := problem.SenseMinimize
	if fmt.Sprintf("%v", minSense) != "Minimize" {
		t.Errorf(
			"minSense's string is \"%v\"; expected \"Minimize\"",
			minSense,
		)
	}

	// Test Maximize
	maxSense := problem.SenseMaximize
	if fmt.Sprintf("%v", maxSense) != "Maximize" {
		t.Errorf(
			"maxSense's string is \"%v\"; expected \"Maximize\"",
			maxSense,
		)
	}

	// Test Find
	findSense := problem.SenseFind
	if fmt.Sprintf("%v", findSense) != "Find" {
		t.Errorf(
			"findSense's string is \"%v\"; expected \"Find\"",
			findSense,
		)
	}
}
