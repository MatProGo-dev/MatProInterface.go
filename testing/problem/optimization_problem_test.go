package problem_test

/*
optimization_problem_test.go
Description:

	Tests for all functions and objects defined in the optimization_problem.go file.
*/

import (
	"github.com/MatProGo-dev/MatProInterface.go/problem"
	"testing"
)

/*
TestOptimizationProblem_NewProblem1
Description:

	Tests the NewProblem function with a simple name.
	Verifies that the name is set correctly and
	that zero variables and constraints exist in the fresh
	problem.
*/
func TestOptimizationProblem_NewProblem1(t *testing.T) {
	// Constants
	name := "TestProblem"

	// New Problem
	problem := problem.NewProblem(name)

	// Check that the name is as expected in the problem.
	if problem.Name != name {
		t.Errorf("expected the name of the problem to be %v; received %v",
			name, problem.Name)
	}

	// Check that the number of variables is zero.
	if len(problem.Variables) != 0 {
		t.Errorf("expected the number of variables to be 0; received %v",
			len(problem.Variables))
	}

	// Check that the number of constraints is zero.
	if len(problem.Constraints) != 0 {
		t.Errorf("expected the number of constraints to be 0; received %v",
			len(problem.Constraints))
	}
}
