package mpiErrors

type NoInequalityConstraintsFoundError struct{}

func (e NoInequalityConstraintsFoundError) Error() string {
	return "No inequality constraints found in the optimization problem; please define some by adding them with \"testProblem1.Constraints = append(testProblem1.Constraints, newConstraint)\""
}
