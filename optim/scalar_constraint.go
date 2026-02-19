package optim

import "fmt"

// ScalarConstraint represents a linear constraint of the form x <= y, x >= y, or
// x == y. ScalarConstraint uses a left and right hand side expression along with a
// constraint sense (<=, >=, ==) to represent a generalized linear constraint.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
type ScalarConstraint struct {
	LeftHandSide  ScalarExpression
	RightHandSide ScalarExpression
	Sense         ConstrSense
}

// Left returns the left hand side expression of the scalar constraint.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (sc ScalarConstraint) Left() Expression {
	return sc.LeftHandSide
}

// Right returns the right hand side expression of the scalar constraint.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (sc ScalarConstraint) Right() Expression {
	return sc.RightHandSide
}

// ConstrSense returns the sense of the scalar constraint.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (sc ScalarConstraint) ConstrSense() ConstrSense {
	return sc.Sense
}

// IsLinear returns true if the scalar constraint is linear (i.e., neither side
// is a ScalarQuadraticExpression).
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (sc ScalarConstraint) IsLinear() (bool, error) {
	// Check left and right side.
	if _, tf := sc.LeftHandSide.(ScalarQuadraticExpression); tf {
		return false, nil
	}

	// If left side has degree less than two, then this only depends
	// on the right side.
	if _, tf := sc.RightHandSide.(ScalarQuadraticExpression); tf {
		return false, nil
	}

	// Otherwise return true
	return true, nil
}

// Simplify moves all of the variables of the ScalarConstraint to its
// left hand side.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (sc ScalarConstraint) Simplify() (ScalarConstraint, error) {
	// Create LHS
	newLHS := sc.LeftHandSide

	// Algorithm
	switch right := sc.RightHandSide.(type) {
	case K:
		return sc, nil
	case Variable:
		newLHS, err := newLHS.Plus(right.Multiply(-1.0))
		if err != nil {
			return sc, err
		}
		newLHSAsSE, _ := ToScalarExpression(newLHS)

		return ScalarConstraint{
			LeftHandSide:  newLHSAsSE,
			RightHandSide: K(0),
			Sense:         sc.Sense,
		}, nil
	case ScalarLinearExpr:
		rightWithoutConstant := right
		rightWithoutConstant.C = 0.0

		newLHS, err := newLHS.Plus(rightWithoutConstant.Multiply(-1.0))
		if err != nil {
			return sc, err
		}
		newLHSAsSE, _ := ToScalarExpression(newLHS)

		return ScalarConstraint{
			LeftHandSide:  newLHSAsSE,
			RightHandSide: K(right.C),
			Sense:         sc.Sense,
		}, nil
	case ScalarQuadraticExpression:
		rightWithoutConstant := right
		rightWithoutConstant.C = 0.0

		newLHS, err := newLHS.Plus(rightWithoutConstant.Multiply(-1.0))
		if err != nil {
			return sc, err
		}
		newLHSAsSE, _ := ToScalarExpression(newLHS)

		return ScalarConstraint{
			LeftHandSide:  newLHSAsSE,
			RightHandSide: K(right.C),
			Sense:         sc.Sense,
		}, nil

	default:
		return sc, fmt.Errorf("unexpected type of right hand side: %T", right)
	}

}

// Check checks the validity of the ScalarConstraint, ensuring the sense is
// recognized and both hand sides are valid.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (sc ScalarConstraint) Check() error {
	// Check sense
	switch sc.Sense {
	case SenseEqual:
		break
	case SenseLessThanEqual:
		break
	case SenseGreaterThanEqual:
		break
	default:
		return fmt.Errorf("the constraint sense is not recognized.")
	}

	// Check left and right hand sides
	err := sc.LeftHandSide.Check()
	if err != nil {
		return fmt.Errorf("left hand side of the constraint is not valid: %v", err)
	}

	// Check right hand side
	err = sc.RightHandSide.Check()
	if err != nil {
		return fmt.Errorf("right hand side of the constraint is not valid: %v", err)
	}

	// Return
	return nil
}
