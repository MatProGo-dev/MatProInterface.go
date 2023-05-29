package optim

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"strings"
	"testing"
)

/*
solution_test.go
Description:
	Testing for the solution object.
	(This seems like it is highly representative of the Gurobi solver; is there a reason to make it this way?)
*/

func TestSolution_ToMessage1(t *testing.T) {
	// Constants
	tempSol := optim.Solution{
		Values: map[uint64]float64{
			0: 2.1,
			1: 3.14,
		},
		Objective: 2.3,
		Status:    optim.OptimizationStatus_NODE_LIMIT,
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
		tempStatus := optim.OptimizationStatus(statusIndex)

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
	tempSol := optim.Solution{
		Values: map[uint64]float64{
			0: 2.1,
			1: 3.14,
		},
		Objective: 2.3,
		Status:    optim.OptimizationStatus_NODE_LIMIT,
	}
	v1 := optim.Variable{
		ID: 0, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}
	v2 := optim.Variable{
		ID: 1, Lower: -optim.INFINITY, Upper: optim.INFINITY, Vtype: optim.Continuous,
	}

	// Algorithm
	if tempSol.Value(v1) != 2.1 {
		t.Errorf(
			"Expected v1 to have value 2.1; received %v",
			tempSol.Value(v1),
		)
	}

	if tempSol.Value(v2) != 3.14 {
		t.Errorf(
			"Expected v2 to have value 3.14; received %v",
			tempSol.Value(v2),
		)
	}

}
