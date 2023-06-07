package optim

// ScalarExpression represents a linear general expression of the form
// c0 * x0 + c1 * x1 + ... + cn * xn + k where ci are coefficients and xi are
// variables and k is a constant. This is a base interface that is implemented
// by single variables, constants, and general linear expressions.
type ScalarExpression interface {
	// Variables returns the variables included in the scalar expression
	Variables() []Variable

	// NumVars returns the number of variables in the expression
	NumVars() int

	// Vars returns a slice of the Var ids in the expression
	IDs() []uint64

	// Coeffs returns a slice of the coefficients in the expression
	Coeffs() []float64

	// Constant returns the constant additive value in the expression
	Constant() float64

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(e interface{}, errors ...error) (ScalarExpression, error)

	// Mult multiplies the current expression to another and returns the
	// resulting expression
	//Mult(c float64, errors ...error) (Expression, error)

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(e ScalarExpression) (ScalarConstraint, error)

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(e ScalarExpression) (ScalarConstraint, error)

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(e ScalarExpression) (ScalarConstraint, error)

	//Comparison
	// Compares the receiver expression rhs with the expression rhs in the sense of sense.
	Comparison(rhs ScalarExpression, sense ConstrSense) (ScalarConstraint, error)

	//Multiply
	// Multiplies the given scalar expression with another expression
	Multiply(term1 interface{}, errors ...error) (Expression, error)
}

// NewExpr returns a new expression with a single additive constant value, c,
// and no variables. Creating an expression like sum := NewExpr(0) is useful
// for creating new empty expressions that you can perform operatotions on
// later
func NewScalarExpression(c float64) ScalarExpression {
	return ScalarLinearExpr{C: c}
}
