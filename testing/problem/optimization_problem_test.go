package problem_test

/*
optimization_problem_test.go
Description:

	Tests for all functions and objects defined in the optimization_problem.go file.
*/

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"github.com/MatProGo-dev/MatProInterface.go/problem"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
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
	name := "TestProblem1"

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

/*
TestOptimizationProblem_AddVariable1
Description:

	Tests the AddVariable function with a simple problem.
*/
func TestOptimizationProblem_AddVariable1(t *testing.T) {
	// Constants
	problem := problem.NewProblem("TestProblem1")

	// Algorithm
	problem.AddVariable()

	// Check that the number of variables is one.
	if len(problem.Variables) != 1 {
		t.Errorf("expected the number of variables to be 1; received %v",
			len(problem.Variables))
	}

	// Verify that the type of the variable is as expected.
	if problem.Variables[0].Type != symbolic.Continuous {
		t.Errorf("expected the type of the variable to be %v; received %v",
			symbolic.Continuous, problem.Variables[0].Type)
	}
}

/*
TestOptimizationProblem_AddRealVariable1
Description:

	Tests the AddRealVariable function with a simple problem.
*/
func TestOptimizationProblem_AddRealVariable1(t *testing.T) {
	// Constants
	problem := problem.NewProblem("TestProblem1")

	// Algorithm
	problem.AddRealVariable()

	// Check that the number of variables is one.
	if len(problem.Variables) != 1 {
		t.Errorf("expected the number of variables to be 1; received %v",
			len(problem.Variables))
	}

	// Verify that the type of the variable is as expected.
	if problem.Variables[0].Type != symbolic.Continuous {
		t.Errorf("expected the type of the variable to be %v; received %v",
			symbolic.Continuous, problem.Variables[0].Type)
	}
}

/*
TestOptimizationProblem_AddVariableClassic1
Description:

	Tests the AddVariableClassic function with a simple problem.
*/
func TestOptimizationProblem_AddVariableClassic1(t *testing.T) {
	// Constants
	problem := problem.NewProblem("TestProblem1")

	// Algorithm
	problem.AddVariableClassic(0, 1, symbolic.Binary)

	// Check that the number of variables is one.
	if len(problem.Variables) != 1 {
		t.Errorf("expected the number of variables to be 1; received %v",
			len(problem.Variables))
	}

	// Verify that the type of the variable is as expected.
	if problem.Variables[0].Type != symbolic.Binary {
		t.Errorf("expected the type of the variable to be %v; received %v",
			symbolic.Binary, problem.Variables[0].Type)
	}
}

/*
TestOptimizationProblem_AddBinaryVariable1
Description:

	Tests the AddBinaryVariable function with a simple problem.
*/
func TestOptimizationProblem_AddBinaryVariable1(t *testing.T) {
	// Constants
	problem := problem.NewProblem("TestProblem1")

	// Algorithm
	problem.AddBinaryVariable()

	// Check that the number of variables is one.
	if len(problem.Variables) != 1 {
		t.Errorf("expected the number of variables to be 1; received %v",
			len(problem.Variables))
	}

	// Verify that the type of the variable is as expected.
	if problem.Variables[0].Type != symbolic.Binary {
		t.Errorf("expected the type of the variable to be %v; received %v",
			symbolic.Binary, problem.Variables[0].Type)
	}
}

/*
TestOptimizationProblem_AddVariableVector1
Description:

	Tests the AddVariableVector function with a simple problem.
*/
func TestOptimizationProblem_AddVariableVector1(t *testing.T) {
	// Constants
	problem := problem.NewProblem("TestProblem1")
	dim := 5

	// Algorithm
	problem.AddVariableVector(dim)

	// Check that the number of variables is as expected.
	if len(problem.Variables) != dim {
		t.Errorf("expected the number of variables to be %v; received %v",
			dim, len(problem.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range problem.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf("expected the type of the variable to be %v; received %v",
				symbolic.Continuous, v.Type)
		}
	}
}

/*
TestOptimizationProblem_AddVariableVectorClassic1
Description:

	Tests the AddVariableVectorClassic function with a simple problem.
*/
func TestOptimizationProblem_AddVariableVectorClassic1(t *testing.T) {
	// Constants
	problem := problem.NewProblem("TestProblem1")
	dim := 5

	// Algorithm
	problem.AddVariableVectorClassic(dim, 0, 1, symbolic.Binary)

	// Check that the number of variables is as expected.
	if len(problem.Variables) != dim {
		t.Errorf("expected the number of variables to be %v; received %v",
			dim, len(problem.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range problem.Variables {
		if v.Type != symbolic.Binary {
			t.Errorf("expected the type of the variable to be %v; received %v",
				symbolic.Binary, v.Type)
		}
	}
}

/*
TestOptimizationProblem_From1
Description:

	Tests the From function with a simple
	model that doesn't have an objective.
*/
func TestOptimizationProblem_From1(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From1",
	)

	N := 5
	for ii := 0; ii < N; ii++ {
		model.AddVariable()
	}

	// Algorithm
	_, err := problem.From(*model)
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			"the input model has no objective function!",
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_From2
Description:

	Tests the From function with a simple
	model that doesn't have an objective.
*/
func TestOptimizationProblem_From2(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From2",
	)

	N := 5
	var tempVar optim.Variable
	for ii := 0; ii < N; ii++ {
		tempVar = model.AddVariable()
	}

	model.SetObjective(tempVar, optim.SenseMaximize)

	// Algorithm
	problem1, err := problem.From(*model)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	if len(problem1.Variables) != 5 {
		t.Errorf("expected the number of variables to be %v; received %v",
			5, len(problem1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range problem1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf(
				"expected the type of the variable to be %v; received %v",
				symbolic.Continuous,
				v.Type,
			)
		}
	}
}
