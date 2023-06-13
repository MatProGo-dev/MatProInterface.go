package optim

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

/*
operators.go
Description:
	Defines the operators that transform variables and expressions into expressions or constraints.
*/

/*
Eq
Description:

	Returns a constraint representing lhs == rhs
*/
func Eq(lhs, rhs interface{}) (Constraint, error) {
	return Comparison(lhs, rhs, SenseEqual)
}

// LessEq returns a constraint representing lhs <= rhs
func LessEq(lhs, rhs interface{}) (Constraint, error) {
	return Comparison(lhs, rhs, SenseLessThanEqual)
}

// GreaterEq returns a constraint representing lhs >= rhs
func GreaterEq(lhs, rhs ScalarExpression) (Constraint, error) {
	return Comparison(lhs, rhs, SenseGreaterThanEqual)
}

/*
Comparison
Description:

	Compares the two inputs lhs (Left Hand Side) and rhs (Right Hand Side) in the sense provided in sense.

Usage:

	constr, err := Comparison(expr1, expr2, SenseGreaterThanEqual)
*/
func Comparison(lhs, rhs interface{}, sense ConstrSense) (Constraint, error) {
	// Constants

	// Algorithm
	switch lhs.(type) {
	case float64:
		// Convert lhs to K
		lhsAsFloat64, _ := lhs.(float64)
		lhsAsK := K(lhsAsFloat64)

		// Create constraint
		return Comparison(lhsAsK, rhs, sense)
	case mat.VecDense:
		// Convert lhs to KVector.
		lhsAsVecDense, _ := lhs.(mat.VecDense)
		lhsAsKVector := KVector(lhsAsVecDense)

		// Create constraint
		return lhsAsKVector.Comparison(rhs, sense)
	case ScalarExpression:
		lhsAsScalarExpression, _ := lhs.(ScalarExpression)
		rhsAsScalarExpression, _ := rhs.(ScalarExpression)
		return ScalarConstraint{
			lhsAsScalarExpression,
			rhsAsScalarExpression,
			sense,
		}, nil
	case VectorExpression:
		lhsAsVecExpr, _ := lhs.(VectorExpression)
		return lhsAsVecExpr.Comparison(rhs, sense)
	default:
		return nil, fmt.Errorf("Comparison in sense '%v' is not defined for lhs type %T and rhs type %T!", sense, lhs, rhs)
	}

}

/*
Multiply
Description:

	Defines the multiplication between two objects.
*/
func Multiply(term1, term2 interface{}) (Expression, error) {
	// Constants

	// Algorithm
	switch term1.(type) {
	case float64:
		// Convert lhs to K
		term1AsFloat64, _ := term1.(float64)
		term1AsK := K(term1AsFloat64)

		// Create constraint
		return Multiply(term1AsK, term2)
	case mat.VecDense:
		// Convert lhs to KVector.
		term1AsVecDense, _ := term1.(mat.VecDense)
		term1AsKVector := KVector(term1AsVecDense)

		// Create constraint
		return term1AsKVector.Multiply(term2)
	case ScalarExpression:
		term1AsScalarExpression, _ := term1.(ScalarExpression)

		// Create Constraint
		return term1AsScalarExpression.Multiply(term2)

	//case VectorExpression:
	//	lhsAsVecExpr, _ := lhs.(VectorExpression)
	//	return lhsAsVecExpr.Multiply(term2)
	default:
		return nil, fmt.Errorf("Multiply of %T term with %T term is not yet defined!", term1, term2)
	}
}

// TODO[Kwesi]: Define dot product logic for vectors.

// Dot returns the dot product of a vector of variables and slice of floats.
//func Dot(vs []Variable, coeffs []float64) ScalarExpression {
//	//if len(vs) != len(coeffs) {
//	//	log.WithFields(log.Fields{
//	//		"num_vars":   len(vs),
//	//		"num_coeffs": len(coeffs),
//	//	}).Panic("Number of vars and coeffs mismatch")
//	//}
//
//	newExpr := NewExpr(0)
//	for i := range vs {
//		newExpr.Plus(vs[i].Multiply(coeffs[i]))
//	}
//
//	return newExpr
//}

// Sum returns the sum of the given expressions. It creates a new empty
// expression and adds to it the given expressions.
func Sum(exprs ...interface{}) (Expression, error) {
	// Constants

	// Input Processing
	// ================

	if !IsExpression(exprs[0]) {
		return ScalarLinearExpr{}, fmt.Errorf("The first input to Sum must be an expression! Received type %T", exprs[0])
	}
	e0, _ := ToExpression(exprs[0])

	if len(exprs) == 1 { // If only one expression was given, then return that.
		return ToExpression(exprs[0])
	}

	// Check whether or not the second argument is an error or not.
	var (
		e1        Expression
		exprIndex int
		tf        bool
	)
	switch secondElt := exprs[1].(type) {
	case error:
		if len(exprs) < 3 {
			return e0, nil
		}

		if secondElt != nil {
			return ScalarLinearExpr{}, fmt.Errorf("An error occurred in the sum: %v", secondElt)
		}
		tf = IsExpression(exprs[2])
		if !tf {
			return ScalarLinearExpr{}, fmt.Errorf("Expected third expression in sum to be an Expression; received %T (%v)", exprs[2], exprs[2])
		}

		e1 = exprs[2].(Expression)

		exprIndex = 3
	case Expression:
		e1 = secondElt
		exprIndex = 2
	case nil:
		if len(exprs) < 3 {
			return e0, nil
		}

		tf = IsExpression(exprs[2])
		if !tf {
			return ScalarLinearExpr{}, fmt.Errorf("Expected third expression in sum to be an Expression; received %T (%v)", exprs[2], exprs[2])
		}
		e1 = exprs[2].(Expression)
		exprIndex = 3
	case float64:
		// Cast variable value
		e1 = K(secondElt)

		// Set next variable to check here.
		exprIndex = 2
	case mat.VecDense:
		e1 = KVector(secondElt)
		exprIndex = 2
	default:
		e1 = ScalarLinearExpr{}
		return ScalarLinearExpr{}, fmt.Errorf("Unexpected input to Sum %v of type %T", exprs[1], exprs[1])
	}

	// Recursive call to sum
	if len(exprs) > exprIndex {
		tempSum, err := Sum(e0, e1)
		if err != nil {
			return e0, fmt.Errorf("Error computing sum between %v and %v: %v", e0, e1, err)
		}

		var tempInter []interface{} = []interface{}{tempSum}
		tempInter = append(tempInter, exprs[exprIndex:]...)
		return Sum(tempInter...)
	}

	// Collect Expression
	// ==================

	switch first := e0.(type) {
	case ScalarExpression:
		return first.Plus(e1)
	case VectorExpression:
		return first.Plus(e1)
	default:
		return e0, fmt.Errorf("Unexpected type input to Sum: %T", e0)
	}
}
