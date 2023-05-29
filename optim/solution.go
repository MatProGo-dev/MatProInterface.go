package optim

import "fmt"

const (
	tinyNum float64 = 0.01
)

// Solution stores the solution of an optimization problem and associated
// metatdata
type Solution struct {
	Values map[uint64]float64

	// The objective for the solution
	Objective float64

	// Whether or not the solution is within the optimality threshold
	Status OptimizationStatus

	// The optimality gap returned from the solver. For many solvers, this is
	// the gap between the best possible solution with integer relaxation and
	// the best integer solution found so far.
	// Gap float64
}

type OptimizationStatus int

// OptimizationStatuses
const (
	OptimizationStatus_LOADED          OptimizationStatus = 1
	OptimizationStatus_OPTIMAL                            = 2
	OptimizationStatus_INFEASIBLE                         = 3
	OptimizationStatus_INF_OR_UNBD                        = 4
	OptimizationStatus_UNBOUNDED                          = 5
	OptimizationStatus_CUTOFF                             = 6
	OptimizationStatus_ITERATION_LIMIT                    = 7
	OptimizationStatus_NODE_LIMIT                         = 8
	OptimizationStatus_TIME_LIMIT                         = 9
	OptimizationStatus_SOLUTION_LIMIT                     = 10
	OptimizationStatus_INTERRUPTED                        = 11
	OptimizationStatus_NUMERIC                            = 12
	OptimizationStatus_SUBOPTIMAL                         = 13
	OptimizationStatus_INPROGRESS                         = 14
	OptimizationStatus_USER_OBJ_LIMIT                     = 15
	OptimizationStatus_WORK_LIMIT                         = 16
)

/*
ToMessage
Description:

	Translates the code to the text meaning.
	This comes from the status codes documentation: https://www.gurobi.com/documentation/9.5/refman/optimization_status_codes.html#sec:StatusCodes
*/
func (os OptimizationStatus) ToMessage() (string, error) {
	// Converts each of the statuses to a text message that is human readable.
	switch os {
	case OptimizationStatus_LOADED:
		return "Model is loaded, but no solution information is available.", nil
	case OptimizationStatus_OPTIMAL:
		return "Model was solved to optimality (subject to tolerances), and an optimal solution is available.", nil
	case OptimizationStatus_INFEASIBLE:
		return "Model was proven to be infeasible.", nil
	case OptimizationStatus_INF_OR_UNBD:
		return "Model was proven to be either infeasible or unbounded. To obtain a more definitive conclusion, set the DualReductions parameter to 0 and reoptimize.", nil
	case OptimizationStatus_UNBOUNDED:
		return "Model was proven to be unbounded. Important note: an unbounded status indicates the presence of an unbounded ray that allows the objective to improve without limit. It says nothing about whether the model has a feasible solution. If you require information on feasibility, you should set the objective to zero and reoptimize.", nil
	case OptimizationStatus_CUTOFF:
		return "Optimal objective for model was proven to be worse than the value specified in the Cutoff parameter. No solution information is available.", nil
	case OptimizationStatus_ITERATION_LIMIT:
		return "Optimization terminated because the total number of simplex iterations performed exceeded the value specified in the IterationLimit parameter, or because the total number of barrier iterations exceeded the value specified in the BarIterLimit parameter.", nil
	case OptimizationStatus_NODE_LIMIT:
		return "Optimization terminated because the total number of branch-and-cut nodes explored exceeded the value specified in the NodeLimit parameter.", nil
	case OptimizationStatus_TIME_LIMIT:
		return "Optimization terminated because the time expended exceeded the value specified in the TimeLimit parameter.", nil
	case OptimizationStatus_SOLUTION_LIMIT:
		return "Optimization terminated because the number of solutions found reached the value specified in the SolutionLimit parameter.", nil
	case OptimizationStatus_INTERRUPTED:
		return "Optimization was terminated by the user.", nil
	case OptimizationStatus_NUMERIC:
		return "Optimization was terminated due to unrecoverable numerical difficulties.", nil
	case OptimizationStatus_SUBOPTIMAL:
		return "Unable to satisfy optimality tolerances; a sub-optimal solution is available.", nil
	case OptimizationStatus_INPROGRESS:
		return "An asynchronous optimization call was made, but the associated optimization run is not yet complete.", nil
	case OptimizationStatus_USER_OBJ_LIMIT:
		return "User specified an objective limit (a bound on either the best objective or the best bound), and that limit has been reached.", nil
	case OptimizationStatus_WORK_LIMIT:
		return "Optimization terminated because the work expended exceeded the value specified in the WorkLimit parameter.", nil
	default:
		return "", fmt.Errorf("The status with value %v is unrecognized.", os)
	}
}

// func newSolution(mipSol solvers.MIPSolution) *Solution {
// 	return &Solution{
// 		vals:      mipSol.GetValues(),
// 		Objective: mipSol.GetObj(),
// 		Optimal:   mipSol.GetOptimal(),
// 		Gap:       mipSol.GetGap(),
// 	}
// }

// Value returns the value assigned to the variable in the solution
func (s *Solution) Value(v Variable) float64 {
	return s.Values[v.ID]
}

//// IsOne returns true if the value assigned to the variable is an integer,
//// and assigned to one. This is a convenience method which should not be
//// super trusted...
//func (s *Solution) IsOne(v Variable) bool {
//	return (v.Vtype == Integer || v.Vtype == Binary) && s.Value(v) > tinyNum
//}
