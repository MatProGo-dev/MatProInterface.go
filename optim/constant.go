package optim

import (
	"fmt"
)

// Integer constants represnting commonly used numbers. Makes for better
// readability
const (
	Zero = K(0)
	One  = K(1)
)

// K is a constant expression type for an MIP. K for short ¯\_(ツ)_/¯
type K float64

/*
Variables
Description:

	Shares all variables included in the expression that is K.
	It is a constant, so there are none.
*/
func (c K) Variables() []Variable {
	return []Variable{}
}

// NumVars returns the number of variables in the expression. For constants,
// this is always 0
func (c K) NumVars() int {
	return 0
}

// Vars returns a slice of the Var ids in the expression. For constants,
// this is always nil
func (c K) IDs() []uint64 {
	return nil
}

// Coeffs returns a slice of the coefficients in the expression. For constants,
// this is always nil
func (c K) Coeffs() []float64 {
	return nil
}

// Constant returns the constant additive value in the expression. For
// constants, this is just the constants value
func (c K) Constant() float64 {
	return float64(c)
}

// Plus adds the current expression to another and returns the resulting
// expression
func (c K) Plus(e interface{}, errors ...error) (ScalarExpression, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return c, err
	}

	// Switching based on input type
	switch e.(type) {
	case K:
		eAsK, _ := e.(K)
		return K(c.Constant() + eAsK.Constant()), nil
	case Variable:
		eAsVar := e.(Variable)
		return eAsVar.Plus(c)
	case ScalarLinearExpr:
		eAsSLE := e.(ScalarLinearExpr)
		return eAsSLE.Plus(c)
	case ScalarQuadraticExpression:
		return e.(ScalarQuadraticExpression).Plus(c) // Very compact, but potentially confusing to read?
	default:
		return c, fmt.Errorf("Unexpected type in K.Plus() for constant %v: %T", e, e)
	}
}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (c K) LessEq(rhsIn interface{}, errors ...error) (ScalarConstraint, error) {
	return c.Comparison(rhsIn, SenseLessThanEqual, errors...)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (c K) GreaterEq(rhsIn interface{}, errors ...error) (ScalarConstraint, error) {
	return c.Comparison(rhsIn, SenseGreaterThanEqual, errors...)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (c K) Eq(rhsIn interface{}, errors ...error) (ScalarConstraint, error) {
	return c.Comparison(rhsIn, SenseEqual, errors...)
}

/*
Comparison
Description:

	This method compares the receiver with expression rhs in the sense provided by sense.
*/
func (c K) Comparison(rhsIn interface{}, sense ConstrSense, errors ...error) (ScalarConstraint, error) {
	// InputProcessing
	err := CheckErrors(errors)
	if err != nil {
		return ScalarConstraint{}, err
	}

	rhs, err := ToScalarExpression(rhsIn)
	if err != nil {
		return ScalarConstraint{}, err
	}

	// Constants

	// Algorithm
	return ScalarConstraint{c, rhs, sense}, nil
}

/*
Multiply
Description:

	This method multiplies the input constant by another expression.
*/
func (c K) Multiply(term1 interface{}, errors ...error) (Expression, error) {
	// Constants

	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return c, err
	}

	if IsVectorExpression(term1) {
		// Check dimensions
		term1AsVE, _ := ToVectorExpression(term1)
		if term1AsVE.Len() != 1 { // Dimension mismatch
			return c, fmt.Errorf(
				"cannot multiply constant %v of shape (1,1) with vector of shape (%v,1)",
				c,
				term1AsVE.Len(),
			)
		}
	}

	// Algorithm
	switch term1Converted := term1.(type) {
	case float64:
		return c.Multiply(K(term1Converted))
	case K:
		return c * term1Converted, nil
	case Variable:
		// Algorithm
		term1AsSLE := term1Converted.ToScalarLinearExpression()

		return c.Multiply(term1AsSLE)
	case ScalarLinearExpr:
		// Scale all vectors and constants
		sleOut := term1Converted.Copy()
		sleOut.L.ScaleVec(float64(c), &sleOut.L)
		sleOut.C = term1Converted.C * float64(c)

		return sleOut, nil
	case ScalarQuadraticExpression:
		// Scale all matrices and constants
		var sqeOut ScalarQuadraticExpression
		sqeOut.Q.Scale(float64(c), &term1Converted.Q)
		sqeOut.L.ScaleVec(float64(c), &term1Converted.L)
		sqeOut.C = float64(c) * term1Converted.C

		return sqeOut, nil
	case KVector:
		return term1Converted.Multiply(c)

	case VectorLinearExpressionTranspose:
		var vletOut VectorLinearExpressionTranspose
		vletOut.L.Scale(float64(c), &term1Converted.L)
		vletOut.C.ScaleVec(float64(c), &term1Converted.C)

		return vletOut, nil
	default:
		return K(0), fmt.Errorf("Unexpected type of term1 in the Multiply() method: %T (%v)", term1, term1)

	}
}

func (c K) Dims() []uint {
	return []uint{1, 1} // Signifies scalar
}
