package optim

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)
// Expression is an interface that should be implemented by any ScalarExpression
// and VectorExpression. It provides methods for arithmetic operations and
// constraint creation.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
type Expression interface {
	// NumVars returns the number of variables in the expression
	NumVars() int

	// Vars returns a slice of the Var ids in the expression
	IDs() []uint64

	// Dims returns a slice describing the true dimensions of a given expression (scalar, vector, or matrix)
	Dims() []int

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(e interface{}, errors ...error) (Expression, error)

	// Multiply multiplies the current expression to another and returns the
	// resulting expression
	Multiply(c interface{}, errors ...error) (Expression, error)

	// Transpose transposes the given expression
	Transpose() Expression

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(rightIn interface{}, errors ...error) (Constraint, error)

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(rightIn interface{}, errors ...error) (Constraint, error)

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(rightIn interface{}, errors ...error) (Constraint, error)

	// Comparison
	Comparison(rightIn interface{}, sense ConstrSense, errors ...error) (Constraint, error)

	//ToSymbolic
	// Converts the expression to a symbolic expression (in SymbolicMath.go)
	ToSymbolic() (symbolic.Expression, error)
}

// IsExpression tests whether or not the input variable is one of the expression types.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func IsExpression(e interface{}) bool {
	return IsScalarExpression(e) || IsVectorExpression(e)
}

// ToExpression converts the input to an Expression, returning an error if the
// input is not a recognized scalar or vector expression type.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func ToExpression(e interface{}) (Expression, error) {
	switch {
	case IsScalarExpression(e):
		return ToScalarExpression(e)
	case IsVectorExpression(e):
		return ToVectorExpression(e)
	default:
		return K(INFINITY), fmt.Errorf("the input expression is not recognized as a scalar or vector expression.")
	}
}

// CheckDimensionsInMultiplication checks that the dimensions of the two
// expressions are compatible for multiplication. It returns an error if the
// number of columns in left does not match the number of rows in right.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func CheckDimensionsInMultiplication(left, right Expression) error {
	// Check that the # of columns in left
	// matches the # of rows in right
	if left.Dims()[1] != right.Dims()[0] {
		return DimensionError{
			Operation: "Multiply",
			Arg1:      left,
			Arg2:      right,
		}
	}
	// If dimensions match, then return nothing.
	return nil
}

// CheckDimensionsInAddition checks that the dimensions of the two expressions
// are compatible for addition. It returns an error if the dimensions do not match,
// unless one of the expressions is a scalar expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func CheckDimensionsInAddition(left, right Expression) error {
	// Check that the size of columns in left and right agree
	dimsAreMatched := (left.Dims()[0] == right.Dims()[0]) && (left.Dims()[1] == right.Dims()[1])
	dimsAreMatched = dimsAreMatched || IsScalarExpression(left)
	dimsAreMatched = dimsAreMatched || IsScalarExpression(right)

	if !dimsAreMatched {
		return DimensionError{
			Operation: "Plus",
			Arg1:      left,
			Arg2:      right,
		}
	}
	// If dimensions match, then return nothing.
	return nil
}
