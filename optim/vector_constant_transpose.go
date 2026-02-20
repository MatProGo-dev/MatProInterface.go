package optim

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
)

// KVectorTranspose is a transposed constant vector expression type for an MIP
// (Mixed Integer Program).
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
type KVectorTranspose mat.VecDense // Inherit all methods from mat.VecDense

// Check verifies that the KVectorTranspose expression is valid. For KVectorTranspose, this always returns nil.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) Check() error {
	return nil
}

// Len computes the length of the KVectorTranspose.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) Len() int {
	kvAsVector := mat.VecDense(kvt)
	return kvAsVector.Len()
}

// AtVec returns the scalar expression value at the given index idx.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) AtVec(idx int) ScalarExpression {
	kvAsVector := mat.VecDense(kvt)
	return K(kvAsVector.AtVec(idx))
}

// NumVars returns the number of variables in the expression. For KVectorTranspose, this is always 0.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) NumVars() int {
	return 0
}

// IDs returns a slice of the Var ids in the expression. For KVectorTranspose, this is always nil.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) IDs() []uint64 {
	return nil
}

// LinearCoeff returns the zero matrix of coefficients. For KVectorTranspose, this is always a zero matrix.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) LinearCoeff() mat.Dense {
	return ZerosMatrix(kvt.Len(), kvt.Len())
}

// Constant returns the constant additive value in the expression, which is
// the vector itself.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) Constant() mat.VecDense {
	return mat.VecDense(kvt)
}

// Plus adds the current expression to another and returns the resulting expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) Plus(rightIn interface{}, errors ...error) (Expression, error) {
	// Constants
	kvLen := kvt.Len()

	// Inpur Processing
	err := CheckErrors(errors)
	if err != nil {
		return kvt, err
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err = CheckDimensionsInAddition(kvt, rightAsE)
		if err != nil {
			return kvt, err
		}
	}

	// Management
	switch right := rightIn.(type) {
	case float64:
		// Create vector
		tempOnes := OnesVector(kvLen)
		var eAsVec mat.VecDense
		eAsVec.ScaleVec(right, &tempOnes)

		// Add the values
		return kvt.Plus(KVectorTranspose(eAsVec))
	case K:
		// Return Addition
		return kvt.Plus(float64(right))
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
			right, right,
		)
	case KVector:
		return kvt, fmt.Errorf(
			"Can not add KVectorTranspose to normal vector %v (type %T); transpose one or the other!",
			right, right,
		)
	case KVectorTranspose:
		// Compute Addition
		var result mat.VecDense
		kvAsVec := mat.VecDense(kvt)
		eAsVec := mat.VecDense(right)
		result.AddVec(&kvAsVec, &eAsVec)

		return KVectorTranspose(result), nil

	case VarVector:
		return kvt, fmt.Errorf(
			"Can not add KVectorTranspose to normal vector %v (type %T); transpose one or the other!",
			right, right,
		)

	case VarVectorTranspose:
		return right.Plus(kvt)

	case VectorLinearExpr:
		return kvt, fmt.Errorf(
			"Can not add KVectorTranspose to normal vector %v (type %T); transpose one or the other!",
			right, right,
		)

	case VectorLinearExpressionTranspose:
		return right.Plus(kvt)

	default:
		errString := fmt.Sprintf("Unrecognized expression type %T for addition of KVectorTranspose kvt.Plus(%v)!", right, right)
		return KVectorTranspose{}, fmt.Errorf(errString)
	}
}

// Mult multiplies the current expression by a scalar float64 and returns
// the resulting expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) Mult(val float64) (VectorExpression, error) {

	// Use mat.Vector's multiplication method
	var result mat.VecDense
	kvAsVec := mat.VecDense(kvt)
	result.ScaleVec(val, &kvAsVec)

	return KVectorTranspose(result), nil
}

// LessEq returns a less than or equal to (<=) constraint between the current
// expression and another.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) LessEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return kvt.Comparison(rightIn, SenseLessThanEqual, errors...)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) GreaterEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return kvt.Comparison(rightIn, SenseGreaterThanEqual, errors...)
}

// Eq returns an equality (==) constraint between the current expression and another.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) Eq(rightIn interface{}, errors ...error) (Constraint, error) {
	return kvt.Comparison(rightIn, SenseEqual, errors...)
}

// Comparison compares the KVectorTranspose with the given expression in the given sense.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) Comparison(rightIn interface{}, sense ConstrSense, errors ...error) (Constraint, error) {
	switch rhs0 := rightIn.(type) {
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
		return VectorConstraint{},
			fmt.Errorf(
				"The input to KVectorTranspose's '%v' comparison (%v) has unexpected type: %T",
				sense, rightIn, rightIn,
			)

	}
}

// Multiply computes the multiplication of the KVectorTranspose with another term.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) Multiply(rightIn interface{}, errors ...error) (Expression, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return kvt, err
	}

	if IsExpression(rightIn) {
		// Check dimensions
		rightAsE, _ := ToExpression(rightIn)
		err = CheckDimensionsInMultiplication(kvt, rightAsE)
		if err != nil {
			return kvt, err
		}
	}

	// Compute Multiplication
	switch right := rightIn.(type) {
	case float64:
		// Use mat.Vector's multiplication method
		var result mat.VecDense
		kvAsVec := mat.VecDense(kvt)
		result.ScaleVec(right, &kvAsVec)

		return KVectorTranspose(result), nil
	case K:
		// Convert to float64
		eAsFloat := float64(right)

		return kvt.Multiply(eAsFloat)

	case mat.VecDense:
		// Do the dot product
		var result float64
		kvtAsVec := mat.VecDense(kvt)
		result = mat.Dot(&kvtAsVec, &right)

		return K(result), nil

	case KVector:
		// Convert to mat.VecDense
		eAsVecDense := mat.VecDense(right)

		return kvt.Multiply(eAsVecDense)

	case KVectorTranspose:
		// Immediately return error.
		return kvt, fmt.Errorf(
			"dimension mismatch! Cannot multiply KVectorTranspose with a transposed vector of type %T; Try transposing one or the other!",
			right,
		)

	case VarVector:
		return right.Transpose().Multiply(kvt.Transpose())

	case VarVectorTranspose:
		// Immediately return error.
		return kvt, fmt.Errorf(
			"dimension mismatch! Cannot multiply KVectorTranspose with a transposed vector of type %T; Try transposing one or the other!",
			right,
		)
	case VectorLinearExpr:
		return right.Multiply(kvt)

	case VectorLinearExpressionTranspose:
		// Immediately return error.
		return kvt, fmt.Errorf(
			"dimension mismatch! Cannot multiply KVectorTranspose with a transposed vector of type %T; Try transposing one or the other!",
			right,
		)

	default:
		return kvt, fmt.Errorf(
			"The input to KVectorTranspose's Multiply method (%v) has unexpected type: %T",
			right, right,
		)

	}
}

// Transpose creates the transpose of the current vector and returns it.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) Transpose() Expression {
	return KVector(kvt)
}

// Dims returns the dimensions of the KVectorTranspose object.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) Dims() []int {
	return []int{1, kvt.Len()}
}

// ToSymbolic returns the symbolic version of the KVectorTranspose expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (kvt KVectorTranspose) ToSymbolic() (symbolic.Expression, error) {
	// Constants
	kvLen := kvt.Len()

	// Create the symbolic expression
	km := symbolic.KMatrix{}

	// Add the constant values
	km = append(km, make([]symbolic.K, kvLen))
	for i := 0; i < kvLen; i++ {
		kvtAsVD := mat.VecDense(kvt)
		km[0][i] = symbolic.K(kvtAsVD.AtVec(i))
	}

	return km, nil
}
