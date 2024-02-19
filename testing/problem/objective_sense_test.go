package problem_test

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"github.com/MatProGo-dev/MatProInterface.go/problem"
	"testing"
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
