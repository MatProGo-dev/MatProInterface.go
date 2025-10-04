package solution

import solution_status "github.com/MatProGo-dev/MatProInterface.go/solution/status"

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
