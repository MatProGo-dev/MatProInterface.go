package optim

import "github.com/MatProGo-dev/SymbolicMath.go/symbolic"

// ConstrSense represents if the constraint x <= y, x >= y, or x == y. For easy
// integration with Gurobi, the senses have been encoding using a byte in
// the same way Gurobi encodes the constraint senses.
type ConstrSense byte

// Different constraint senses conforming to Gurobi's encoding.
const (
	SenseEqual            ConstrSense = '='
	SenseLessThanEqual                = '<'
	SenseGreaterThanEqual             = '>'
)

/*
ToSymbolic
Description:

	Converts a constraint sense to a the appropriate representation
	in the symbolic math toolbox.
*/
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
