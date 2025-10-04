package solution_test

import (
	"strings"
	"testing"

	"github.com/MatProGo-dev/MatProInterface.go/solution"
	solution_status "github.com/MatProGo-dev/MatProInterface.go/solution/status"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

/*
 */
type DummySolution struct {
	Values map[uint64]float64

	// The objective for the solution
	Objective float64

	// Whether or not the solution is within the optimality threshold
	Status solution_status.SolutionStatus

	// The optimality gap returned from the solver. For many solvers, this is
	// the gap between the best possible solution with integer relaxation and
	// the best integer solution found so far.
	// Gap float64
}

func (ds *DummySolution) GetOptimalValue() float64 {
	return ds.Objective
}

func (ds *DummySolution) GetValueMap() map[uint64]float64 {
	return ds.Values
}

func (ds *DummySolution) GetStatus() solution_status.SolutionStatus {
	return ds.Status
}

/*
solution_test.go
Description:
	Testing for the solution object.
	(This seems like it is highly representative of the Gurobi solver; is there a reason to make it this way?)
*/

func TestSolution_ToMessage1(t *testing.T) {
	// Constants
	tempSol := DummySolution{
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

	tempSol := DummySolution{
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
		t.Errorf("the value of the variable v2 could not be extracted; received error %v", err)
	}

	if v2Val != 3.14 {
		t.Errorf(
			"Expected v2 to have value 3.14; received %v",
			v2Val,
		)
	}

}
