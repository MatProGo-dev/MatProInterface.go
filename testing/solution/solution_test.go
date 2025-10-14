package solution_test

import (
	"strings"
	"testing"

	"github.com/MatProGo-dev/MatProInterface.go/problem"
	"github.com/MatProGo-dev/MatProInterface.go/solution"
	solution_status "github.com/MatProGo-dev/MatProInterface.go/solution/status"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

/*
solution_test.go
Description:
	Testing for the solution object.
	(This seems like it is highly representative of the Gurobi solver; is there a reason to make it this way?)
*/

func TestSolution_ToMessage1(t *testing.T) {
	// Constants
	tempSol := solution.DummySolution{
		Values: map[uint64]float64{
			0: 2.1,
			1: 3.14,
		},
		Objective: 2.3,
		Status:    solution_status.NODE_LIMIT,
	}

	// Test the ToMessage() Call on this solution.
	msg, err := tempSol.Status.ToMessage()
	if err != nil {
		t.Errorf("There was an error transforming optimization status to message: %v", err)
	}

	if msg != "Optimization terminated because the total number of branch-and-cut nodes explored exceeded the value specified in the NodeLimit parameter." {
		t.Errorf(
			"Expected for the optimization status to be the NODE_LIMIT; received %v",
			msg,
		)
	}
}

func TestSolution_ToMessage2(t *testing.T) {
	// Constants
	statusMax := 16

	// Test
	for statusIndex := 1; statusIndex < statusMax; statusIndex++ {
		tempStatus := solution_status.SolutionStatus(statusIndex)

		msg, err := tempStatus.ToMessage()
		if err != nil {
			t.Errorf(
				"Provided valid optimization status to ToMessage(), but an error occurred: %v",
				err,
			)
		}

		var expectedMessages = []string{
			"Model is loaded, but no solution information is available.",
			"Model was solved to optimality (subject to tolerances), and an optimal solution is available.",
			"Model was proven to be infeasible.",
			"Model was proven to be either infeasible or unbounded. To obtain a more definitive conclusion, set the DualReductions parameter to 0 and reoptimize.",
			"Model was proven to be unbounded. Important note: an unbounded status indicates the presence of an unbounded ray that allows the objective to improve without limit. It says nothing about whether the model has a feasible solution. If you require information on feasibility, you should set the objective to zero and reoptimize.",
			"Optimal objective for model was proven to be worse than the value specified in the Cutoff parameter. No solution information is available.",
			"Optimization terminated because the total number of simplex iterations performed exceeded the value specified in the IterationLimit parameter, or because the total number of barrier iterations exceeded the value specified in the BarIterLimit parameter.",
			"Optimization terminated because the total number of branch-and-cut nodes explored exceeded the value specified in the NodeLimit parameter.",
			"Optimization terminated because the time expended exceeded the value specified in the TimeLimit parameter.",
			"Optimization terminated because the number of solutions found reached the value specified in the SolutionLimit parameter.",
			"Optimization was terminated by the user.",
			"Optimization was terminated due to unrecoverable numerical difficulties.",
			"Unable to satisfy optimality tolerances; a sub-optimal solution is available.",
			"An asynchronous optimization call was made, but the associated optimization run is not yet complete.",
			"User specified an objective limit (a bound on either the best objective or the best bound), and that limit has been reached.",
			"Optimization terminated because the work expended exceeded the value specified in the WorkLimit parameter.",
		}

		if strings.Compare(expectedMessages[statusIndex], msg) == 0 {
			t.Errorf(
				"Expected message \"%v\" \n for status %v, received: \"%v\"",
				expectedMessages[statusIndex-1],
				statusIndex,
				msg,
			)
		}

	}
}

/*
TestSolution_Value1
Description:

	This function tests whether or not we can properly retrieve values from a solution object.
*/
func TestSolution_Value1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()

	tempSol := solution.DummySolution{
		Values: map[uint64]float64{
			v1.ID: 2.1,
			v2.ID: 3.14,
		},
		Objective: 2.3,
		Status:    solution_status.NODE_LIMIT,
	}

	// Algorithm
	v1Val, err := solution.ExtractValueOfVariable(&tempSol, v1)
	if err != nil {
		t.Errorf("The value of the variable v1 could not be extracted; received error %v", err)
	}

	if v1Val != 2.1 {
		t.Errorf(
			"Expected v1 to have value 2.1; received %v",
			v1Val,
		)
	}

	v2Val, err := solution.ExtractValueOfVariable(&tempSol, v2)
	if err != nil {
		t.Errorf("The value of the variable v2 could not be extracted; received error %v", err)
	}

	if v2Val != 3.14 {
		t.Errorf(
			"Expected v2 to have value 3.14; received %v",
			v2Val,
		)
	}

}

/*
TestSolution_FindValueOfExpression1
Description:

	This function tests whether we can evaluate a simple linear expression
	using the solution values.
*/
func TestSolution_FindValueOfExpression1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()

	tempSol := solution.DummySolution{
		Values: map[uint64]float64{
			v1.ID: 2.0,
			v2.ID: 3.0,
		},
		Objective: 2.3,
		Status:    solution_status.OPTIMAL,
	}

	// Create expression: 2*v1 + 3*v2 = 2*2.0 + 3*3.0 = 4.0 + 9.0 = 13.0
	expr := v1.Multiply(symbolic.K(2.0)).Plus(v2.Multiply(symbolic.K(3.0)))

	// Algorithm
	result, err := solution.FindValueOfExpression(&tempSol, expr)
	if err != nil {
		t.Errorf("FindValueOfExpression returned an error: %v", err)
	}

	expected := 13.0
	if result != expected {
		t.Errorf(
			"Expected expression value to be %v; received %v",
			expected,
			result,
		)
	}
}

/*
TestSolution_FindValueOfExpression2
Description:

	This function tests whether we can evaluate a constant expression.
*/
func TestSolution_FindValueOfExpression2(t *testing.T) {
	// Constants
	tempSol := solution.DummySolution{
		Values:    map[uint64]float64{},
		Objective: 2.3,
		Status:    solution_status.OPTIMAL,
	}

	// Create constant expression: 42.0
	expr := symbolic.K(42.0)

	// Algorithm
	result, err := solution.FindValueOfExpression(&tempSol, expr)
	if err != nil {
		t.Errorf("FindValueOfExpression returned an error: %v", err)
	}

	expected := 42.0
	if result != expected {
		t.Errorf(
			"Expected expression value to be %v; received %v",
			expected,
			result,
		)
	}
}

/*
TestSolution_FindValueOfExpression3
Description:

	This function tests whether we can evaluate an expression with a single variable.
*/
func TestSolution_FindValueOfExpression3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()

	tempSol := solution.DummySolution{
		Values: map[uint64]float64{
			v1.ID: 5.5,
		},
		Objective: 2.3,
		Status:    solution_status.OPTIMAL,
	}

	// Create expression: v1 + 10 = 5.5 + 10 = 15.5
	expr := v1.Plus(symbolic.K(10.0))

	// Algorithm
	result, err := solution.FindValueOfExpression(&tempSol, expr)
	if err != nil {
		t.Errorf("FindValueOfExpression returned an error: %v", err)
	}

	expected := 15.5
	if result != expected {
		t.Errorf(
			"Expected expression value to be %v; received %v",
			expected,
			result,
		)
	}
}

/*
TestSolution_FindValueOfExpression4
Description:

	This function tests whether we get an error when a variable is missing
	from the solution.
*/
func TestSolution_FindValueOfExpression4(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()

	tempSol := solution.DummySolution{
		Values: map[uint64]float64{
			v1.ID: 2.0,
			// v2 is missing
		},
		Objective: 2.3,
		Status:    solution_status.OPTIMAL,
	}

	// Create expression: v1 + v2
	expr := v1.Plus(v2)

	// Algorithm
	_, err := solution.FindValueOfExpression(&tempSol, expr)
	if err == nil {
		t.Errorf("Expected FindValueOfExpression to return an error for missing variable, but got nil")
	}
}

/*
TestSolution_FindValueOfExpression5
Description:

	This function tests whether we can evaluate a more complex expression
	with multiple operations.
*/
func TestSolution_FindValueOfExpression5(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	v3 := symbolic.NewVariable()

	tempSol := solution.DummySolution{
		Values: map[uint64]float64{
			v1.ID: 1.0,
			v2.ID: 2.0,
			v3.ID: 3.0,
		},
		Objective: 2.3,
		Status:    solution_status.OPTIMAL,
	}

	// Create expression: (v1 + v2) * v3 + 5 = (1.0 + 2.0) * 3.0 + 5 = 3.0 * 3.0 + 5 = 9.0 + 5 = 14.0
	expr := v1.Plus(v2).Multiply(v3).Plus(symbolic.K(5.0))

	// Algorithm
	result, err := solution.FindValueOfExpression(&tempSol, expr)
	if err != nil {
		t.Errorf("FindValueOfExpression returned an error: %v", err)
	}

	expected := 14.0
	if result != expected {
		t.Errorf(
			"Expected expression value to be %v; received %v",
			expected,
			result,
		)
	}
}

/*
TestSolution_GetProblem1
Description:

	This function tests whether we can retrieve the problem from a solution.
*/
func TestSolution_GetProblem1(t *testing.T) {
	// Constants
	p := problem.NewProblem("TestProblem1")
	v1 := p.AddVariable()

	tempSol := solution.DummySolution{
		Values: map[uint64]float64{
			v1.ID: 2.1,
		},
		Objective: 2.3,
		Status:    solution_status.OPTIMAL,
		Problem:   p,
	}

	// Algorithm
	retrievedProblem := tempSol.GetProblem()

	// Verify the problem is the same
	if retrievedProblem != p {
		t.Errorf("Expected GetProblem to return the same problem pointer")
	}

	if retrievedProblem.Name != "TestProblem1" {
		t.Errorf("Expected problem name to be 'TestProblem1'; received %v", retrievedProblem.Name)
	}
}

/*
TestSolution_GetProblem2
Description:

	This function tests whether GetProblem returns nil when no problem is set.
*/
func TestSolution_GetProblem2(t *testing.T) {
	// Constants
	tempSol := solution.DummySolution{
		Values: map[uint64]float64{
			0: 2.1,
		},
		Objective: 2.3,
		Status:    solution_status.OPTIMAL,
		Problem:   nil,
	}

	// Algorithm
	retrievedProblem := tempSol.GetProblem()

	// Verify the problem is nil
	if retrievedProblem != nil {
		t.Errorf("Expected GetProblem to return nil when no problem is set")
	}
}

/*
TestSolution_GetOptimalObjectiveValue1
Description:

	This function tests whether we can compute the objective value at the solution point
	for a simple linear objective.
*/
func TestSolution_GetOptimalObjectiveValue1(t *testing.T) {
	// Constants
	p := problem.NewProblem("TestProblem")
	v1 := p.AddVariable()
	v2 := p.AddVariable()

	// Set objective: 2*v1 + 3*v2
	objectiveExpr := v1.Multiply(symbolic.K(2.0)).Plus(v2.Multiply(symbolic.K(3.0)))
	err := p.SetObjective(objectiveExpr, problem.SenseMinimize)
	if err != nil {
		t.Errorf("Failed to set objective: %v", err)
	}

	// Create solution with v1=2.0, v2=3.0
	// Expected objective value: 2*2.0 + 3*3.0 = 4.0 + 9.0 = 13.0
	tempSol := solution.DummySolution{
		Values: map[uint64]float64{
			v1.ID: 2.0,
			v2.ID: 3.0,
		},
		Objective: 13.0,
		Status:    solution_status.OPTIMAL,
		Problem:   p,
	}

	// Algorithm
	objectiveValue, err := solution.GetOptimalObjectiveValue(&tempSol)
	if err != nil {
		t.Errorf("GetOptimalObjectiveValue returned an error: %v", err)
	}

	expected := 13.0
	if objectiveValue != expected {
		t.Errorf(
			"Expected objective value to be %v; received %v",
			expected,
			objectiveValue,
		)
	}
}

/*
TestSolution_GetOptimalObjectiveValue2
Description:

	This function tests whether GetOptimalObjectiveValue returns an error
	when the solution has no associated problem.
*/
func TestSolution_GetOptimalObjectiveValue2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()

	tempSol := solution.DummySolution{
		Values: map[uint64]float64{
			v1.ID: 2.0,
		},
		Objective: 2.3,
		Status:    solution_status.OPTIMAL,
		Problem:   nil,
	}

	// Algorithm
	_, err := solution.GetOptimalObjectiveValue(&tempSol)
	if err == nil {
		t.Errorf("Expected GetOptimalObjectiveValue to return an error for nil problem, but got nil")
	}
}

/*
TestSolution_GetOptimalObjectiveValue3
Description:

	This function tests whether we can compute the objective value
	for a constant objective function.
*/
func TestSolution_GetOptimalObjectiveValue3(t *testing.T) {
	// Constants
	p := problem.NewProblem("TestProblem")
	v1 := p.AddVariable()

	// Set constant objective: 42.0
	objectiveExpr := symbolic.K(42.0)
	err := p.SetObjective(objectiveExpr, problem.SenseMaximize)
	if err != nil {
		t.Errorf("Failed to set objective: %v", err)
	}

	// Create solution
	tempSol := solution.DummySolution{
		Values: map[uint64]float64{
			v1.ID: 1.0,
		},
		Objective: 42.0,
		Status:    solution_status.OPTIMAL,
		Problem:   p,
	}

	// Algorithm
	objectiveValue, err := solution.GetOptimalObjectiveValue(&tempSol)
	if err != nil {
		t.Errorf("GetOptimalObjectiveValue returned an error: %v", err)
	}

	expected := 42.0
	if objectiveValue != expected {
		t.Errorf(
			"Expected objective value to be %v; received %v",
			expected,
			objectiveValue,
		)
	}
}

/*
TestSolution_GetOptimalObjectiveValue4
Description:

	This function tests whether we can compute the objective value
	for a more complex objective with multiple variables and operations.
*/
func TestSolution_GetOptimalObjectiveValue4(t *testing.T) {
	// Constants
	p := problem.NewProblem("TestProblem")
	v1 := p.AddVariable()
	v2 := p.AddVariable()
	v3 := p.AddVariable()

	// Set objective: (v1 + v2) * v3 + 5
	objectiveExpr := v1.Plus(v2).Multiply(v3).Plus(symbolic.K(5.0))
	err := p.SetObjective(objectiveExpr, problem.SenseMinimize)
	if err != nil {
		t.Errorf("Failed to set objective: %v", err)
	}

	// Create solution with v1=1.0, v2=2.0, v3=3.0
	// Expected objective: (1.0 + 2.0) * 3.0 + 5 = 3.0 * 3.0 + 5 = 14.0
	tempSol := solution.DummySolution{
		Values: map[uint64]float64{
			v1.ID: 1.0,
			v2.ID: 2.0,
			v3.ID: 3.0,
		},
		Objective: 14.0,
		Status:    solution_status.OPTIMAL,
		Problem:   p,
	}

	// Algorithm
	objectiveValue, err := solution.GetOptimalObjectiveValue(&tempSol)
	if err != nil {
		t.Errorf("GetOptimalObjectiveValue returned an error: %v", err)
	}

	expected := 14.0
	if objectiveValue != expected {
		t.Errorf(
			"Expected objective value to be %v; received %v",
			expected,
			objectiveValue,
		)
	}
}
