package optim

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// Var represnts a variable in a optimization problem. The variable is
// identified with an uint64.
type Variable struct {
	ID    uint64
	Lower float64
	Upper float64
	Vtype VarType
}

/*
Variables
Description:

	This function returns a slice containing all unique variables in the variable expression v.
*/
func (v Variable) Variables() []Variable {
	return []Variable{v}
}

// NumVars returns the number of variables in the expression. For a variable, it
// always returns one.
func (v Variable) NumVars() int {
	return 1
}

// Vars returns a slice of the Var ids in the expression. For a variable, it
// always returns a singleton slice with the given variable ID.
func (v Variable) IDs() []uint64 {
	return []uint64{v.ID}
}

// Coeffs returns a slice of the coefficients in the expression. For a variable,
// it always returns a singleton slice containing the value one.
func (v Variable) Coeffs() []float64 {
	return []float64{1}
}

// Constant returns the constant additive value in the expression. For a
// variable, it always returns zero.
func (v Variable) Constant() float64 {
	return 0
}

// Plus adds the current expression to another and returns the resulting
// expression.
func (v Variable) Plus(e interface{}, errors ...error) (ScalarExpression, error) {
	// Input Processing??

	// Algorithm
	switch e.(type) {
	case K:
		eAsK := e.(K)

		// Organize vector variables
		vv := VarVector{
			UniqueVars(append([]Variable{v}, eAsK.Variables()...)),
		}

		// Return
		return ScalarLinearExpr{
			L: OnesVector(1),
			X: vv,
			C: float64(eAsK),
		}, nil
	case Variable:
		// Convert
		eAsV := e.(Variable)

		vv := VarVector{
			UniqueVars(append([]Variable{v}, eAsV.Variables()...)),
		}

		// Check to see if this is the same Variable or a different one
		if eAsV.ID == v.ID {
			return ScalarLinearExpr{
				X: vv,
				L: *mat.NewVecDense(1, []float64{2.0}),
				C: 0.0,
			}, nil
		} else {
			return ScalarLinearExpr{
				X: vv,
				L: OnesVector(2),
				C: 0.0,
			}, nil
		}
	case ScalarLinearExpr:
		// Convert
		eAsSLE := e.(ScalarLinearExpr)

		vv := VarVector{
			UniqueVars(append([]Variable{v}, eAsSLE.Variables()...)),
		}

		// Convert SLE to new form
		e2, _ := eAsSLE.RewriteInTermsOf(vv)
		vIndex, _ := FindInSlice(v, vv.Elements)
		e2.L.SetVec(vIndex, e2.L.AtVec(vIndex)+1.0)

		return e2, nil

	case ScalarQuadraticExpression:
		// Convert
		eAsQE := e.(ScalarQuadraticExpression)

		vv := VarVector{
			UniqueVars(append([]Variable{v}, eAsQE.Variables()...)),
		}

		// Convert QE to new form
		e2, _ := eAsQE.RewriteInTermsOf(vv)
		vIndex, _ := FindInSlice(v, vv.Elements)
		e2.L.SetVec(vIndex, e2.L.AtVec(vIndex)+1.0)

		return e2, nil

	default:
		return v, fmt.Errorf("There was an unexpected type (%T) given to Variable.Plus()!", e)
	}
}

//// Mult multiplies the current expression to another and returns the
//// resulting expression
//func (v Variable) Mult(m float64) (ScalarExpression, error) {
//	// Constants
//	// switch m.(type) {
//	// case float64:
//
//	vars := []Variable{v}
//	coeffs := []float64{m * v.Coeffs()[0]}
//
//	// Algorithm
//	newExpr := ScalarLinearExpr{
//		X: VarVector{vars},
//		L: *mat.NewVecDense(1, coeffs),
//		C: 0,
//	}
//	return newExpr, nil
//	// case *Variable:
//	// 	return nil
//	// }
//}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (v Variable) LessEq(other ScalarExpression) (ScalarConstraint, error) {
	return v.Comparison(other, SenseLessThanEqual)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (v Variable) GreaterEq(other ScalarExpression) (ScalarConstraint, error) {
	return v.Comparison(other, SenseGreaterThanEqual)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (v Variable) Eq(other ScalarExpression) (ScalarConstraint, error) {
	return v.Comparison(other, SenseEqual)
}

/*
Comparison
Description:

	This method compares the receiver with expression rhs in the sense provided by sense.

Usage:

	constr, err := v.Comparison(expr1,SenseGreaterThanEqual)
*/
func (v Variable) Comparison(rhs ScalarExpression, sense ConstrSense) (ScalarConstraint, error) {
	// Constants

	// Algorithm
	return ScalarConstraint{v, rhs, sense}, nil
}

/*
// ID returns the ID of the variable
func (v *Variable) ID() uint64 {
	return v.ID
}

// Lower returns the lower value limit of the variable
func (v *Variable) Lower() float64 {
	return v.Lower
}

// Upper returns the upper value limit of the variable
func (v *Variable) Upper() float64 {
	return v.Upper
}

// Type returns the type of variable (continuous, binary, integer, etc)
func (v *Variable) Type() VarType {
	return v.Vtype
}
*/

// VarType represents the type of the variable (continuous, binary,
// integer, etc) and uses Gurobi's encoding.
type VarType byte

// Multiple common variable types have been included as constants that conform
// to Gurobi's encoding.
const (
	Continuous VarType = 'C'
	Binary             = 'B'
	Integer            = 'I'
)

/*
UniqueVars
Description:

	This function creates a slice of unique variables from the slice given in
	varsIn
*/
func UniqueVars(varsIn []Variable) []Variable {
	// Constants

	// Algorithm
	var varsOut []Variable
	for _, v := range varsIn {
		if vIndex, _ := FindInSlice(v, varsOut); vIndex == -1 { // If v is not yet in varsOut, then add it
			varsOut = append(varsOut, v)
		}
	}

	return varsOut

}

/*
Multiply
Description:

	multiplies the current expression to another and returns the resulting expression
*/
func (v Variable) Multiply(val interface{}, errors ...error) (Expression, error) {
	// Input Processing
	// TODO: Finish input processing!

	// Constants
	switch val.(type) {
	case float64:
		// Cast
		valAsFloat := val.(float64)

		// Algorithm
		vars := []Variable{v}
		coeffs := []float64{valAsFloat * v.Coeffs()[0]}

		// Algorithm
		newExpr := ScalarLinearExpr{
			X: VarVector{vars},
			L: *mat.NewVecDense(1, coeffs),
			C: 0,
		}
		return newExpr, nil
	default:
		return v, fmt.Errorf("Unexpected input to v.Multiply(): %T", val)
	}
}

/*
ToScalarLinearExpression
Description:

	Converting the variable into a scalar linear Expression.
*/
func (v Variable) ToScalarLinearExpression() ScalarLinearExpr {
	// Constants

	// Create components
	vars := []Variable{v}
	coeffs := []float64{v.Coeffs()[0]}

	// Create sle
	return ScalarLinearExpr{
		X: VarVector{vars},
		L: *mat.NewVecDense(1, coeffs),
		C: 0,
	}
}
