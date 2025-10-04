package solution_status

import "fmt"

type SolutionStatus int

// OptimizationStatuses
const (
	LOADED          SolutionStatus = 1
	OPTIMAL         SolutionStatus = 2
	INFEASIBLE      SolutionStatus = 3
	INF_OR_UNBD     SolutionStatus = 4
	UNBOUNDED       SolutionStatus = 5
	CUTOFF          SolutionStatus = 6
	ITERATION_LIMIT SolutionStatus = 7
	NODE_LIMIT      SolutionStatus = 8
	TIME_LIMIT      SolutionStatus = 9
	SOLUTION_LIMIT  SolutionStatus = 10
	INTERRUPTED     SolutionStatus = 11
	NUMERIC         SolutionStatus = 12
	SUBOPTIMAL      SolutionStatus = 13
	INPROGRESS      SolutionStatus = 14
	USER_OBJ_LIMIT  SolutionStatus = 15
	WORK_LIMIT      SolutionStatus = 16
)

/*
ToMessage
Description:

	Translates the code to the text meaning.
	This comes from the status codes documentation: https://www.gurobi.com/documentation/9.5/refman/optimization_status_codes.html#sec:StatusCodes
*/
func (os SolutionStatus) ToMessage() (string, error) {
	// Converts each of the statuses to a text message that is human readable.
	switch os {
	case LOADED:
		return "Model is loaded, but no solution information is available.", nil
	case OPTIMAL:
		return "Model was solved to optimality (subject to tolerances), and an optimal solution is available.", nil
	case INFEASIBLE:
		return "Model was proven to be infeasible.", nil
	case INF_OR_UNBD:
		return "Model was proven to be either infeasible or unbounded. To obtain a more definitive conclusion, set the DualReductions parameter to 0 and reoptimize.", nil
	case UNBOUNDED:
		return "Model was proven to be unbounded. Important note: an unbounded status indicates the presence of an unbounded ray that allows the objective to improve without limit. It says nothing about whether the model has a feasible solution. If you require information on feasibility, you should set the objective to zero and reoptimize.", nil
	case CUTOFF:
		return "Optimal objective for model was proven to be worse than the value specified in the Cutoff parameter. No solution information is available.", nil
	case ITERATION_LIMIT:
		return "Optimization terminated because the total number of simplex iterations performed exceeded the value specified in the IterationLimit parameter, or because the total number of barrier iterations exceeded the value specified in the BarIterLimit parameter.", nil
	case NODE_LIMIT:
		return "Optimization terminated because the total number of branch-and-cut nodes explored exceeded the value specified in the NodeLimit parameter.", nil
	case TIME_LIMIT:
		return "Optimization terminated because the time expended exceeded the value specified in the TimeLimit parameter.", nil
	case SOLUTION_LIMIT:
		return "Optimization terminated because the number of solutions found reached the value specified in the SolutionLimit parameter.", nil
	case INTERRUPTED:
		return "Optimization was terminated by the user.", nil
	case NUMERIC:
		return "Optimization was terminated due to unrecoverable numerical difficulties.", nil
	case SUBOPTIMAL:
		return "Unable to satisfy optimality tolerances; a sub-optimal solution is available.", nil
	case INPROGRESS:
		return "An asynchronous optimization call was made, but the associated optimization run is not yet complete.", nil
	case USER_OBJ_LIMIT:
		return "User specified an objective limit (a bound on either the best objective or the best bound), and that limit has been reached.", nil
	case WORK_LIMIT:
		return "Optimization terminated because the work expended exceeded the value specified in the WorkLimit parameter.", nil
	default:
		return "", fmt.Errorf("The status with value %v is unrecognized.", os)
	}
}
