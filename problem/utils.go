package problem

import "github.com/MatProGo-dev/SymbolicMath.go/symbolic"

// ConstraintIsRedundantGivenOthers returns true if one of the provided
// constraints implies the input constraint is already satisfied.
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
