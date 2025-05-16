package mpiErrors

import (
	"fmt"

	"github.com/MatProGo-dev/MatProInterface.go/causeOfProblemNonlinearity"
)

type ProblemNotLinearError struct {
	ProblemName     string
	Cause           causeOfProblemNonlinearity.Cause
	ConstraintIndex int // Index of the constraint that is not linear, if applicable
}

func (e ProblemNotLinearError) Error() string {
	preamble := "The problem " + e.ProblemName + " is not linear"
	switch e.Cause {
	case causeOfProblemNonlinearity.Objective:
		preamble += "; the objective is not linear"
	case causeOfProblemNonlinearity.Constraint:
		preamble += fmt.Sprintf("; constraint #%v is not linear", e.ConstraintIndex)
	case causeOfProblemNonlinearity.NotWellDefined:
		preamble += "; the problem is not well defined"
	default:
		preamble += "; the cause of the problem is not recognized (create an issue on GitHub if you think this is a bug!)"
	}
	// Return the error message
	return preamble
}
