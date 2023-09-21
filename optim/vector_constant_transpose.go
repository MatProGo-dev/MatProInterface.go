package optim

import "gonum.org/v1/gonum/mat"

/*
vector_constant_test.go
Description:
	Creates a vector extension of the constant type K from the original goop.
*/

import (
	"fmt"
)

/*
KVectorTranspose

	A type which is built on top of the KVector()
	a constant expression type for an MIP. K for short ¯\_(ツ)_/¯
*/
type KVectorTranspose mat.VecDense // Inherit all methods from mat.VecDense

/*
Len

	Computes the length of the KVector given.
*/
func (kvt KVectorTranspose) Len() int {
	kvAsVector := mat.VecDense(kvt)
	return kvAsVector.Len()
}

/*
AtVec
Description:

	This function returns the value at the k index.
*/
func (kvt KVectorTranspose) AtVec(idx int) ScalarExpression {
	kvAsVector := mat.VecDense(kvt)
	return K(kvAsVector.AtVec(idx))
}

/*
NumVars
Description:

	This returns the number of variables in the expression. For constants, this is 0.
*/
func (kvt KVectorTranspose) NumVars() int {
	return 0
}

/*
Vars
Description:

	This function returns a slice of the Var ids in the expression. For constants, this is always nil.
*/
func (kvt KVectorTranspose) IDs() []uint64 {
	return nil
}

/*
LinearCoeff
Description:

	This function returns a slice of the coefficients in the expression. For constants, this is always nil.
*/
func (kvt KVectorTranspose) LinearCoeff() mat.Dense {
	return ZerosMatrix(kvt.Len(), kvt.Len())
}

/*
Constant

	Returns the constant additive value in the expression. For constants, this is just the constants value
*/
func (kvt KVectorTranspose) Constant() mat.VecDense {
	return mat.VecDense(kvt)
}

/*
Plus
Description:

	Adds the current expression to another and returns the resulting expression
*/
func (kvt KVectorTranspose) Plus(eIn interface{}, extras ...interface{}) (VectorExpression, error) {
	// Constants
	kvLen := kvt.Len()

	// Extras Management

	// Management
	switch e := eIn.(type) {
	case float64:
		// Create vector
		tempOnes := OnesVector(kvLen)
		var eAsVec mat.VecDense
		eAsVec.ScaleVec(e, &tempOnes)

		// Add the values
		return kvt.Plus(KVectorTranspose(eAsVec))
	case K:
		// Return Addition
		return kvt.Plus(float64(e))
	//case mat.VecDense:
	//	// Input Checking
	//	if kvLen != e.Len() {
	//		return kvt, fmt.Errorf(
	//			"Length of vectors in sum do not match! Vectors have lengths %v and %v!",
	//			kvt.Len(), e.Len(),
	//		)
	//	}
	//	// Return Sum
	//	var result mat.VecDense
	//	kv2 := mat.VecDense(kvt)
	//	result.AddVec(&kv2, &e)
	//
	//	return KVectorTranspose(result), nil
	case mat.VecDense:
		return kvt, fmt.Errorf(
			"Can not add KVectorTranspose to normal vector %v (type %T); transpose one or the other!",
			e, e,
		)
	case KVector:
		return kvt, fmt.Errorf(
			"Can not add KVectorTranspose to normal vector %v (type %T); transpose one or the other!",
			e, e,
		)
	case KVectorTranspose:
		// Input Checking
		if kvLen != e.Len() {
			return kvt, fmt.Errorf(
				"Length of vectors in sum do not match! Vectors have lengths %v and %v!",
				kvt.Len(), e.Len(),
			)
		}

		// Compute Addition
		var result mat.VecDense
		kvAsVec := mat.VecDense(kvt)
		eAsVec := mat.VecDense(e)
		result.AddVec(&kvAsVec, &eAsVec)

		return KVectorTranspose(result), nil

	case VarVector:
		return kvt, fmt.Errorf(
			"Can not add KVectorTranspose to normal vector %v (type %T); transpose one or the other!",
			e, e,
		)

	case VarVectorTranspose:
		return e.Plus(kvt)

	case VectorLinearExpr:
		return kvt, fmt.Errorf(
			"Can not add KVectorTranspose to normal vector %v (type %T); transpose one or the other!",
			e, e,
		)

	case VectorLinearExpressionTranspose:
		return e.Plus(kvt)

	default:
		errString := fmt.Sprintf("Unrecognized expression type %T for addition of KVectorTranspose kvt.Plus(%v)!", e, e)
		return KVectorTranspose{}, fmt.Errorf(errString)
	}
}

/*
Mult
Description:

	This method multiplies the current expression to another and returns the resulting expression.
*/
func (kvt KVectorTranspose) Mult(val float64) (VectorExpression, error) {

	// Use mat.Vector's multiplication method
	var result mat.VecDense
	kvAsVec := mat.VecDense(kvt)
	result.ScaleVec(val, &kvAsVec)

	return KVectorTranspose(result), nil
}

/*
LessEq
Description:

	Returns a less than or equal to (<=) constraint between the current expression and another
*/
func (kvt KVectorTranspose) LessEq(rhsIn interface{}) (VectorConstraint, error) {
	return kvt.Comparison(rhsIn, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	This method returns a greater than or equal to (>=) constraint between the current expression and another
*/
func (kvt KVectorTranspose) GreaterEq(rhsIn interface{}) (VectorConstraint, error) {
	return kvt.Comparison(rhsIn, SenseGreaterThanEqual)
}

/*
Eq
Description:

	This method returns an equality (==) constraint between the current expression and another
*/
func (kvt KVectorTranspose) Eq(rhsIn interface{}) (VectorConstraint, error) {
	return kvt.Comparison(rhsIn, SenseEqual)
}

func (kvt KVectorTranspose) Comparison(rhs interface{}, sense ConstrSense) (VectorConstraint, error) {
	switch rhs0 := rhs.(type) {
	case KVector:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare KVectorTranspose with a normal vector %T; Try transposing one or the other!",
				rhs0,
			)

	case KVectorTranspose:
		// Check Lengths
		if kvt.Len() != rhs0.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The left hand side's dimension (%v) and the left hand side's dimension (%v) do not match!",
					kvt.Len(),
					rhs0.Len(),
				)
		}

		// Return constraint
		return VectorConstraint{
			LeftHandSide:  kvt,
			RightHandSide: rhs0,
			Sense:         sense,
		}, nil

	case VarVector:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare KVectorTranspose with a normal vector of type %T; Try transposing one or the other!",
				rhs0,
			)
	case VarVectorTranspose:
		// Return constraint
		return rhs0.Comparison(kvt, sense)

	case VectorLinearExpr:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare KVectorTranspose with a normal vector %T; Try transposing one or the other!",
				rhs0,
			)

	case VectorLinearExpressionTranspose:
		// Return constraint
		return rhs0.Comparison(kvt, sense)
	default:
		// Return an error
		return VectorConstraint{}, fmt.Errorf("The input to KVectorTranspose's '%v' comparison (%v) has unexpected type: %T", sense, rhs, rhs)

	}
}

/*
Multiply
Description:

	This method is used to compute the multiplication of the input vector constant with another term.
*/
func (kvt KVectorTranspose) Multiply(e interface{}, extras ...interface{}) (Expression, error) {
	// Input Processing
	err := CheckExtras(extras)
	if err != nil {
		return kvt, err
	}

	if IsVectorExpression(e) {
		// Check dimensions
		e2, _ := ToVectorExpression(e)
		if e2.Len() != kvt.Len() {
			return kvt, fmt.Errorf(
				"KVectorTranspose of length %v can not be multiplied with a %T of different length (%v).",
				kvt.Len(),
				e2,
				e2.Len(),
			)
		}
	}

	// Compute Multiplication
	switch eConverted := e.(type) {
	case float64:
		// Use mat.Vector's multiplication method
		var result mat.VecDense
		kvAsVec := mat.VecDense(kvt)
		result.ScaleVec(eConverted, &kvAsVec)

		return KVectorTranspose(result), nil
	case K:
		// Convert to float64
		eAsFloat := float64(eConverted)

		return kvt.Multiply(eAsFloat)

	case mat.VecDense:
		// Do the dot product
		var result float64
		kvtAsVec := mat.VecDense(kvt)
		result = mat.Dot(&kvtAsVec, &eConverted)

		return K(result), nil

	case KVector:
		// Convert to mat.VecDense
		eAsVecDense := mat.VecDense(eConverted)

		return kvt.Multiply(eAsVecDense)

	case KVectorTranspose:
		// Immediately return error.
		return kvt, fmt.Errorf(
			"dimension mismatch! Cannot multiply KVectorTranspose with a transposed vector of type %T; Try transposing one or the other!",
			eConverted,
		)

	case VarVector:
		return eConverted.Transpose().Multiply(kvt.Transpose())

	case VarVectorTranspose:
		// Immediately return error.
		return kvt, fmt.Errorf(
			"dimension mismatch! Cannot multiply KVectorTranspose with a transposed vector of type %T; Try transposing one or the other!",
			eConverted,
		)
	case VectorLinearExpr:
		return eConverted.Multiply(kvt)

	case VectorLinearExpressionTranspose:
		// Immediately return error.
		return kvt, fmt.Errorf(
			"dimension mismatch! Cannot multiply KVectorTranspose with a transposed vector of type %T; Try transposing one or the other!",
			eConverted,
		)

	default:
		return kvt, fmt.Errorf(
			"The input to KVectorTranspose's Multiply method (%v) has unexpected type: %T",
			e, e,
		)

	}
}

/*
Transpose
Description:

	This method creates the transpose of the current vector and returns it.
*/
func (kvt KVectorTranspose) Transpose() VectorExpression {
	return KVector(kvt)
}
