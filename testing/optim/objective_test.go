package optim

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"testing"
)

/*
objective_test.go
Description:

	Tests for the Objective object.
*/
func TestObjective_NewObjective1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-newobjective1")
	v := m.AddVariable()

	// Create objective expression
	obj0 := optim.NewObjective(v, optim.SenseMaximize)
	if obj0.Sense != optim.SenseMaximize {
		t.Errorf(
			"The sense of the objective is expected to be maximize; received %v",
			obj0.Sense,
		)
	}
}
