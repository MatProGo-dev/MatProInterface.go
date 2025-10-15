package problem

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

func ConstraintIsRedundantGivenOthers(
	constraint symbolic.Constraint,
	constraints []symbolic.Constraint,
) bool {
	// Check if the expression can be derived from the constraints
	for _, c := range constraints {
		if c.ImpliesThisIsAlsoSatisfied(constraint) {
			return true
		}
	}

	return false
}
