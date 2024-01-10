package problem_test

import (
	"github.com/MatProGo-dev/MatProInterface.go/problem"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
TestObjective_NewObjective1
Description:

	This test verifies that a new objective can be created
	with the NewObjective function. It also verifies the types
	of the returned value of NewObjective.
*/
func TestObjective_NewObjective1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()

	// Algorithm
	obj1 := problem.NewObjective(v1, problem.SenseMaximize)

	// Check the type of the returned value.
	if obj1.Sense != problem.SenseMaximize {
		t.Errorf(
			"The objective sense is not properly set; received %v!",
			obj1.Sense,
		)
	}

}
