package optim

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

/*
expression.go
Description:
	This file holds all of the functions and methods related to the Expression
	interface.
*/

/*
Expression
Description:

	This interface should be implemented by and ScalarExpression and VectorExpression
*/
type Expression interface {
	// NumVars returns the number of variables in the expression
	NumVars() int

	// Vars returns a slice of the Var ids in the expression
	IDs() []uint64

	//// Plus adds the current expression to another and returns the resulting
	//// expression
	// Plus(e Expression, extras interface{}) (Expression, error)
	//
	//// Mult multiplies the current expression to another and returns the
	//// resulting expression
	//Multiply(c interface{}) (Expression, error)
	//
	//// LessEq returns a less than or equal to (<=) constraint between the
	//// current expression and another
	//LessEq(e Expression) Constraint
	//
	//// GreaterEq returns a greater than or equal to (>=) constraint between the
	//// current expression and another
	//GreaterEq(e Expression) Constraint
	//
	//// Eq returns an equality (==) constraint between the current expression
	//// and another
	//Eq(e ScalarExpression) *ScalarConstraint
}

func ToExpression(eIn interface{}) (Expression, error) {
	// Constants

	// Algorithm

	// Attempt conversion to float64
	if eAsFloat, tf := eIn.(float64); tf {
		return K(eAsFloat), nil
	}

	if eAsK, tf := eIn.(K); tf {
		return eAsK, nil
	}

	if eAsVariable, tf := eIn.(Variable); tf {
		return eAsVariable, nil
	}

	switch eIn.(type) {
	case float64:
		eAsFloat, _ := eIn.(float64)
		return K(eAsFloat), nil
	case K:
		eAsK, _ := eIn.(K)
		return eAsK, nil
	case Variable:
		eAsVar, _ := eIn.(Variable)
		return eAsVar, nil
	case ScalarLinearExpr:
		eAsSLE, _ := eIn.(ScalarLinearExpr)
		return eAsSLE, nil
	case ScalarQuadraticExpression:
		eAsSQE, _ := eIn.(ScalarQuadraticExpression)
		return eAsSQE, nil
	case mat.VecDense:
		eAsVD, _ := eIn.(mat.VecDense)
		return ToExpression(KVector(eAsVD))
	case KVector:
		eAsKV, _ := eIn.(KVector)
		return eAsKV, nil
	case VarVector:
		eAsVV, _ := eIn.(VarVector)
		return eAsVV, nil
	case VectorLinearExpr:
		eAsVLE, _ := eIn.(VectorLinearExpr)
		return eAsVLE, nil
	default:
		return K(-1.0), fmt.Errorf("Unexpected type input to ToExpression(): %T", eIn)
	}
}

/*
IsExpression
Description:

	Tests whether or not the input variable is one of the expression types.
*/
func IsExpression(e interface{}) bool {
	// Constants

	// Checks
	_, isScalarExpression := e.(ScalarExpression)
	_, isVectorExpression := e.(VectorExpression)

	return isScalarExpression || isVectorExpression
}
