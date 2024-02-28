package optim

import "fmt"

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

func (sc ScalarConstraint) ConstrSense() ConstrSense {
	return sc.Sense
}

/*
IsLinear
Description:

	Describes whether or not a given linear constraint is
	linear or not.
*/
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

/*
Simplify
Description:

	Moves all of the variables of the ScalarConstraint to its
	left hand side.
*/
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

/*
Check
Description:

	Checks the validity of the ScalarConstraint, this makes sure that:
	- The Sense if either SenseEqual, SenseLessThanEqual, or SenseGreaterThanEqual
*/
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
