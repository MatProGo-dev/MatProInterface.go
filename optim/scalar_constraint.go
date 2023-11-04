package optim

// ScalarConstraint represnts a linear constraint of the form x <= y, x >= y, or
// x == y. ScalarConstraint uses a left and right hand side expressions along with a
// constraint sense (<=, >=, ==) to represent a generalized linear constraint
type ScalarConstraint struct {
	LeftHandSide  ScalarExpression
	RightHandSide ScalarExpression
	Sense         ConstrSense
}

func (sc ScalarConstraint) Left() Expression {
	return sc.LeftHandSide
}

func (sc ScalarConstraint) Right() Expression {
	return sc.RightHandSide
}

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
