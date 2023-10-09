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
KVector

	A type which is built on top of the KVector()
	a constant expression type for an MIP. K for short ¯\_(ツ)_/¯
*/
type KVector mat.VecDense // Inherit all methods from mat.VecDense

/*
Len

	Computes the length of the KVector given.
*/
func (kv KVector) Len() int {
	kvAsVector := mat.VecDense(kv)
	return kvAsVector.Len()
}

/*
AtVec
Description:

	This function returns the value at the k index.
*/
func (kv KVector) AtVec(idx int) ScalarExpression {
	kvAsVector := mat.VecDense(kv)
	return K(kvAsVector.AtVec(idx))
}

/*
NumVars
Description:

	This returns the number of variables in the expression. For constants, this is 0.
*/
func (kv KVector) NumVars() int {
	return 0
}

/*
Vars
Description:

	This function returns a slice of the Var ids in the expression. For constants, this is always nil.
*/
func (kv KVector) IDs() []uint64 {
	return nil
}

/*
LinearCoeff
Description:

	This function returns a slice of the coefficients in the expression. For constants, this is always nil.
*/
func (kv KVector) LinearCoeff() mat.Dense {
	return ZerosMatrix(kv.Len(), kv.Len())
}

/*
Constant

	Returns the constant additive value in the expression. For constants, this is just the constants value
*/
func (kv KVector) Constant() mat.VecDense {
	return mat.VecDense(kv)
}

/*
Plus
Description:

	Adds the current expression to another and returns the resulting expression
*/
func (kv KVector) Plus(eIn interface{}, errors ...error) (VectorExpression, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return kv, err
	}

	// Constants
	kvLen := kv.Len()

	// Management
	switch e := eIn.(type) {
	case float64:
		// Create vector
		tempOnes := OnesVector(kvLen)
		var eAsVec mat.VecDense
		eAsVec.ScaleVec(e, &tempOnes)

		// Add the values
		return kv.Plus(KVector(eAsVec))
	case K:
		// Return Addition
		return kv.Plus(float64(e))
	case mat.VecDense:
		// Input Checking
		if kvLen != e.Len() {
			return kv, fmt.Errorf(
				"Length of vectors in sum do not match! Vectors have lengths %v and %v!",
				kv.Len(), e.Len(),
			)
		}
		// Return Sum
		var result mat.VecDense
		kv2 := mat.VecDense(kv)
		result.AddVec(&kv2, &e)

		return KVector(result), nil
	case KVector:
		// Input Checking
		if kvLen != e.Len() {
			return kv, fmt.Errorf(
				"Length of vectors in sum do not match! Vectors have lengths %v and %v!",
				kv.Len(), e.Len(),
			)
		}

		// Compute Addition
		var result mat.VecDense
		kvAsVec := mat.VecDense(kv)
		eAsVec := mat.VecDense(e)
		result.AddVec(&kvAsVec, &eAsVec)

		return KVector(result), nil

	case KVectorTranspose:
		return kv,
			fmt.Errorf(
				"Cannot add KVector with a transposed vector %T; Try transposing one or the other!",
				e,
			)

	case VarVector:
		return e.Plus(kv)

	case VarVectorTranspose:
		return kv,
			fmt.Errorf(
				"Cannot add KVector with a transposed vector %T; Try transposing one or the other!",
				e,
			)

	case VectorLinearExpr:
		return e.Plus(kv)

	case VectorLinearExpressionTranspose:
		return kv,
			fmt.Errorf(
				"Cannot add KVector with a transposed vector %T; Try transposing one or the other!",
				e,
			)

	default:
		errString := fmt.Sprintf("Unrecognized expression type %T for addition of KVector kv.Plus(%v)!", e, e)
		return KVector{}, fmt.Errorf(errString)
	}
}

/*
Mult
Description:

	This method multiplies the current expression to another and returns the resulting expression.
*/
func (kv KVector) Mult(val float64) (VectorExpression, error) {

	// Use mat.Vector's multiplication method
	var result mat.VecDense
	kvAsVec := mat.VecDense(kv)
	result.ScaleVec(val, &kvAsVec)

	return KVector(result), nil
}

/*
LessEq
Description:

	Returns a less than or equal to (<=) constraint between the current expression and another
*/
func (kv KVector) LessEq(rhsIn interface{}) (VectorConstraint, error) {
	return kv.Comparison(rhsIn, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	This method returns a greater than or equal to (>=) constraint between the current expression and another
*/
func (kv KVector) GreaterEq(rhsIn interface{}) (VectorConstraint, error) {
	return kv.Comparison(rhsIn, SenseGreaterThanEqual)
}

/*
Eq
Description:

	This method returns an equality (==) constraint between the current expression and another
*/
func (kv KVector) Eq(rhsIn interface{}) (VectorConstraint, error) {
	return kv.Comparison(rhsIn, SenseEqual)
}

func (kv KVector) Comparison(rhs interface{}, sense ConstrSense) (VectorConstraint, error) {
	switch rhsConverted := rhs.(type) {
	case KVector:
		// Check Lengths
		if kv.Len() != rhsConverted.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The left hand side's dimension (%v) and the left hand side's dimension (%v) do not match!",
					kv.Len(),
					rhsConverted.Len(),
				)
		}

		// Return constraint
		return VectorConstraint{
			LeftHandSide:  kv,
			RightHandSide: rhsConverted,
			Sense:         sense,
		}, nil
	case KVectorTranspose:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare KVector with a transposed vector %T; Try transposing one or the other!",
				rhsConverted,
			)
	case VarVector:
		// Return constraint
		return rhsConverted.Comparison(kv, sense)
	case VarVectorTranspose:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare KVector with a transposed vector %T; Try transposing one or the other!",
				rhsConverted,
			)
	case VectorLinearExpr:
		// Return constraint
		return rhsConverted.Comparison(kv, sense)
	case VectorLinearExpressionTranspose:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare KVector with a transposed vector %T; Try transposing one or the other!",
				rhsConverted,
			)
	default:
		// Return an error
		return VectorConstraint{}, fmt.Errorf("The input to KVector's '%v' comparison (%v) has unexpected type: %T", sense, rhs, rhs)

	}
}

/*
Multiply
Description:

	This method is used to compute the multiplication of the input vector constant with another term.
*/
func (kv KVector) Multiply(e interface{}, errors ...error) (Expression, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return kv, err
	}

	if IsVectorExpression(e) {
		// Check dimensions
		e2, _ := ToVectorExpression(e)
		if e2.Len() != kv.Len() {
			return kv, fmt.Errorf(
				"KVector of length %v can not be multiplied with a %T of different length (%v).",
				kv.Len(),
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
		kvAsVec := mat.VecDense(kv)
		result.ScaleVec(eConverted, &kvAsVec)

		return KVector(result), nil
	case K:
		// Convert to float64
		eAsFloat := float64(eConverted)

		return kv.Multiply(eAsFloat)
	case Variable:
		// Create a VectorLinearExpression Output
		vv := VarVector{[]Variable{eConverted}}
		L := ZerosMatrix(kv.Len(), 1)
		for kIndex := 0; kIndex < kv.Len(); kIndex++ {
			L.Set(kIndex, 0, float64(kv.AtVec(kIndex).(K)))
		}

		return VectorLinearExpr{
			L: L,
			X: vv,
			C: ZerosVector(kv.Len()),
		}, nil

	case ScalarLinearExpr:
		// Create a VectorLinearExpression to Output
		L := ZerosMatrix(kv.Len(), eConverted.X.Len())
		for rowIndex := 0; rowIndex < kv.Len(); rowIndex++ {
			for colIndex := 0; colIndex < eConverted.X.Len(); colIndex++ {
				L.Set(
					rowIndex, colIndex,
					eConverted.L.AtVec(colIndex)*float64(kv.AtVec(rowIndex).(K)),
				)
			}
		}
		C := ZerosVector(kv.Len())
		for rowIndex := 0; rowIndex < kv.Len(); rowIndex++ {
			C.SetVec(
				rowIndex,
				eConverted.C*float64(kv.AtVec(rowIndex).(K)),
			)
		}
		return VectorLinearExpr{
			L: L,
			C: C,
			X: eConverted.X.Copy(),
		}, nil

	case mat.VecDense:
		// Send warning until we create matrix type.
		return kv, fmt.Errorf(
			"MatProInterface does not currently support operations that result in matrices! if you want this feature, create an issue!",
		)

	case KVector:
		// Immediately return error.
		return kv, fmt.Errorf(
			"dimension mismatch! Cannot multiply KVector with a vector of type %T; Try transposing one or the other!",
			eConverted,
		)

	case KVectorTranspose:
		// Send warning until we create matrix type.
		return kv, fmt.Errorf(
			"MatProInterface does not currently support operations that result in matrices! if you want this feature, create an issue!",
		)

	case VarVector:
		// Immediately return error.
		return kv, fmt.Errorf(
			"dimension mismatch! Cannot multiply KVector with a vector of type %T; Try transposing one or the other!",
			eConverted,
		)

	case VarVectorTranspose:
		// Send warning until we create matrix type.
		return kv, fmt.Errorf(
			"MatProInterface does not currently support operations that result in matrices! if you want this feature, create an issue!",
		)

	case VectorLinearExpr:
		// Immediately return error.
		return kv, fmt.Errorf(
			"dimension mismatch! Cannot multiply KVector with a vector of type %T; Try transposing one or the other!",
			eConverted,
		)

	case VectorLinearExpressionTranspose:
		// Send warning until we create matrix type.
		return kv, fmt.Errorf(
			"MatProInterface does not currently support operations that result in matrices! if you want this feature, create an issue!",
		)

	default:
		return kv, fmt.Errorf(
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
func (kv KVector) Transpose() VectorExpression {
	return KVectorTranspose(kv)
}
