package optimization_problem_errors

type NotWellDefinedError struct {
	ProblemName string
	ErrorSource error
}

func (e NotWellDefinedError) Error() string {
	return "the problem " + e.ProblemName + " is not well defined: " + e.ErrorSource.Error()
}
