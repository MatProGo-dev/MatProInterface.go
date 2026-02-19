package optim

import "github.com/MatProGo-dev/SymbolicMath.go/symbolic"

// ConstrSense represents if the constraint x <= y, x >= y, or x == y. For easy
// integration with Gurobi, the senses have been encoded using a byte in
// the same way Gurobi encodes the constraint senses.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
type ConstrSense byte

// Different constraint senses conforming to Gurobi's encoding.
const (
	// SenseEqual represents the equality constraint sense (==).
	//
	// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
	SenseEqual ConstrSense = '='
	// SenseLessThanEqual represents the less-than-or-equal constraint sense (<=).
	//
	// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
	SenseLessThanEqual = '<'
	// SenseGreaterThanEqual represents the greater-than-or-equal constraint sense (>=).
	//
	// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
	SenseGreaterThanEqual = '>'
)

// ToSymbolic converts a constraint sense to the appropriate representation
// in the symbolic math toolbox.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (cs ConstrSense) ToSymbolic() symbolic.ConstrSense {
	switch cs {
	case SenseEqual:
		return symbolic.SenseEqual
	case SenseLessThanEqual:
		return symbolic.SenseLessThanEqual
	case SenseGreaterThanEqual:
		return symbolic.SenseGreaterThanEqual
	}
	return '1'
}

// String returns the string representation of the constraint sense.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (cs ConstrSense) String() string {
	switch cs {
	case SenseEqual:
		return "=="
	case SenseLessThanEqual:
		return "<="
	case SenseGreaterThanEqual:
		return ">="
	}
	return "UNRECOGNIZED ConstrSense"
}
