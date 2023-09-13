package optim

/*
vector_expression.go
Description:
	An improvement/successor to the scalar expr interface.
*/

import "gonum.org/v1/gonum/mat"

/*
VectorExpression
Description:

	This interface represents any expression written in terms of a
	vector of represents a linear general expression of the form
		c0 * x0 + c1 * x1 + ... + cn * xn + k where ci are coefficients and xi are
	variables and k is a constant. This is a base interface that is implemented
	by single variables, constants, and general linear expressions.
*/
type VectorExpression interface {
	// NumVars returns the number of variables in the expression
	NumVars() int

	// IDs returns a slice of the Var ids in the expression
	IDs() []uint64

	// Coeffs returns a slice of the coefficients in the expression
	LinearCoeff() mat.Dense

	// Constant returns the constant additive value in the expression
	Constant() mat.VecDense

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(e interface{}, extras ...interface{}) (VectorExpression, error)

	// Mult multiplies the current expression with another and returns the
	// resulting expression
	Multiply(e interface{}, extras ...interface{}) (Expression, error)

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(rhs interface{}) (VectorConstraint, error)

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(rhs interface{}) (VectorConstraint, error)

	// Comparison
	// Returns a constraint with respect to the sense (senseIn) between the
	// current expression and another.
	Comparison(rhs interface{}, sense ConstrSense) (VectorConstraint, error)

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(rhs interface{}) (VectorConstraint, error)

	// Len returns the length of the vector expression.
	Len() int

	//AtVec returns the expression at a given index
	AtVec(idx int) ScalarExpression

	//Transpose returns the transpose of the given vector expression
	Transpose() VectorExpression
}

/*
NewVectorExpression
Description:

	NewExpr returns a new expression with a single additive constant value, c,
	and no variables. Creating an expression like sum := NewVectorExpr(0) is useful
	for creating new empty expressions that you can perform operatotions on later
*/
func NewVectorExpression(c mat.VecDense) VectorLinearExpr {
	return VectorLinearExpr{C: c}
}

//func (e VectorExpression) getVarsPtr() *uint64 {
//
//	if e.NumVars() > 0 {
//		return &e.IDs()[0]
//	}
//
//	return nil
//}
//
//func (e VectorExpression) getCoeffsPtr() *float64 {
//	if e.NumVars() > 0 {
//		return &e.Coeffs()[0]
//	}
//
//	return nil
//}
