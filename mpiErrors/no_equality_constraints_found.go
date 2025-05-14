package mpiErrors

type NoEqualityConstraintsFoundError struct{}

func (e NoEqualityConstraintsFoundError) Error() string {
	return "No equality constraints found in the optimization problem; please define some by adding them with \"testProblem1.Constraints = append(testProblem1.Constraints, newConstraint)\""
}
