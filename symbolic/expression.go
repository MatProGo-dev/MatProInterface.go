package symbolic

/*
matrix_expression.go
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

	// Dims returns a slice describing the true dimensions of a given expression (scalar, vector, or matrix)
	Dims() []int

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(e interface{}, errors ...error) (Expression, error)

	// Multiply multiplies the current expression to another and returns the
	// resulting expression
	Multiply(c interface{}, errors ...error) (Expression, error)
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

/*
IsExpression
Description:

	Tests whether or not the input variable is one of the expression types.
*/
//func IsExpression(e interface{}) bool {
//	return IsScalarExpression(e) || IsVectorExpression(e)
//}
//
//func ToExpression(e interface{}) (Expression, error) {
//	switch {
//	case IsScalarExpression(e):
//		return ToScalarExpression(e)
//	case IsVectorExpression(e):
//		return ToVectorExpression(e)
//	default:
//		return K(INFINITY), fmt.Errorf("the input expression is not recognized as a scalar or vector expression.")
//	}
//}
//
//func CheckDimensionsInMultiplication(left, right Expression) error {
//	// Check that the # of columns in left
//	// matches the # of rows in right
//	if left.Dims()[1] != right.Dims()[0] {
//		return DimensionError{
//			Operation: "Multiply",
//			Arg1:      left,
//			Arg2:      right,
//		}
//	}
//	// If dimensions match, then return nothing.
//	return nil
//}
//
//func CheckDimensionsInAddition(left, right Expression) error {
//	// Check that the size of columns in left and right agree
//	if (left.Dims()[0] != right.Dims()[0]) || (left.Dims()[1] != right.Dims()[1]) {
//		return DimensionError{
//			Operation: "Plus",
//			Arg1:      left,
//			Arg2:      right,
//		}
//	}
//	// If dimensions match, then return nothing.
//	return nil
//}
