package optim

// Constraint is an interface for use with the ScalarConstraint and
// VectorConstraint objects.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
type Constraint interface {
	Left() Expression
	Right() Expression
	ConstrSense() ConstrSense
	Check() error
}

// IsConstraint returns true if the input is a valid Constraint type
// (ScalarConstraint or VectorConstraint).
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func IsConstraint(c interface{}) bool {
	switch c.(type) {
	case ScalarConstraint:
		return true
	case *ScalarConstraint:
		return true
	case VectorConstraint:
		return true
	case *VectorConstraint:
		return true
	}

	// Return false, if the constraint is not a scalar or vector constraint.
	return false
}
