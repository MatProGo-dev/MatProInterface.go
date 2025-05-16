package causeOfProblemNonlinearity

type Cause string

const (
	Objective      Cause = "Objective"
	Constraint           = "Constraint"
	NotWellDefined       = "NotWellDefined"
)
