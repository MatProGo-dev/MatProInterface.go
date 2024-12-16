package mpiErrors

type NoObjectiveDefinedError struct{}

func (e NoObjectiveDefinedError) Error() string {
	return "No objective defined for the optimization problem; please define one with SetObjective()"
}
