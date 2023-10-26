package optim

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

/*
var_vector.go
Description:
	The VarVector type will represent a
*/

/*
VarVector
Description:

	Represnts a variable in a optimization problem. The variable is
*/
type VarVector struct {
	Elements []Variable
}

// =========
// Functions
// =========

/*
Length
Description:

	Returns the length of the vector of optimization variables.
*/
func (vv VarVector) Length() int {
	return len(vv.Elements)
}

/*
Len
Description:

	This function is created to mirror the GoNum Vector API. Does the same thing as Length.
*/
func (vv VarVector) Len() int {
	return vv.Length()
}

/*
At
Description:

	Mirrors the gonum api for vectors. This extracts the element of the variable vector at the index x.
*/
func (vv VarVector) AtVec(idx int) ScalarExpression {
	// Constants

	// Algorithm
	return vv.Elements[idx]
}

/*
IDs
Description:

	Returns the unique indices
*/
func (vv VarVector) IDs() []uint64 {
	// Algorithm
	var IDSlice []uint64

	for _, elt := range vv.Elements {
		IDSlice = append(IDSlice, elt.ID)
	}

	return Unique(IDSlice)

}

/*
NumVars
Description:

	The number of unique variables inside the variable vector.
*/
func (vv VarVector) NumVars() int {
	return len(vv.IDs())
}

/*
Constant
Description:

	Returns an all zeros vector as output from the method.
*/
func (vv VarVector) Constant() mat.VecDense {
	zerosOut := ZerosVector(vv.Len())
	return zerosOut
}

/*
LinearCoeff
Description:

	Returns the matrix which is multiplied by Variables to get the current "expression".
	For a single vector, this is an identity matrix.
*/
func (vv VarVector) LinearCoeff() mat.Dense {
	return Identity(vv.Len())
}

/*
Plus
Description:

	This member function computes the addition of the receiver vector var with the
	incoming vector expression ve.
*/
func (vv VarVector) Plus(e interface{}, errors ...error) (VectorExpression, error) {
	// Constants
	vvLen := vv.Len()

	// Processing Extras

	// Algorithm
	switch eAsType := e.(type) {
	case KVector:
		// Check Lengths
		if eAsType.Len() != vv.Len() {
			return VarVector{},
				fmt.Errorf(
					"The lengths of two vectors in Plus must match! VarVector has dimension %v, KVector has dimension %v",
					vv.Len(),
					eAsType.Len(),
				)
		}

		// Algorithm
		return VectorLinearExpr{
			L: Identity(vvLen),
			X: vv,
			C: mat.VecDense(eAsType),
		}, nil
	case mat.VecDense:
		// Call KVector version
		return vv.Plus(KVector(eAsType))

	case KVectorTranspose:
		return vv,
			fmt.Errorf(
				"Cannot add VarVector with a transposed vector %v (%T); Try transposing one or the other!",
				eAsType, eAsType,
			)

	case VarVector:
		// Use VLE based plus
		eAsVLE := VectorLinearExpr{
			L: Identity(eAsType.Len()),
			X: eAsType,
			C: ZerosVector(eAsType.Len()),
		}

		return vv.Plus(eAsVLE)

	case VarVectorTranspose:
		return vv,
			fmt.Errorf(
				"Cannot add VarVector with a transposed vector %v (%T); Try transposing one or the other!",
				eAsType, eAsType,
			)

	case VectorLinearExpr:
		return eAsType.Plus(vv)

	case VectorLinearExpressionTranspose:
		return vv,
			fmt.Errorf(
				"Cannot add VarVector with a transposed vector %v (%T); Try transposing one or the other!",
				eAsType, eAsType,
			)

	default:
		errString := fmt.Sprintf("Unrecognized expression type %T for addition of VarVector vv.Plus(%v)!", e, e)
		return VarVector{}, fmt.Errorf(errString)
	}
}

/*
Mult
Description:

	This member function computest the multiplication of the receiver vector var with some
	incoming vector expression (may result in quadratic?).
*/
func (vv VarVector) Mult(c float64) (VectorExpression, error) {
	return vv, fmt.Errorf("The Mult() method for VarVector is not implemented yet!")
}

/*
Multiply
Description:

	Multiplication of a VarVector with another expression.
*/
func (vv VarVector) Multiply(rightIn interface{}, errors ...error) (Expression, error) {
	//Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return vv, err
	}

	if IsVectorExpression(rightIn) {
		rightAsVE, _ := ToVectorExpression(rightIn)
		if vv.Dims()[1] != rightAsVE.Dims()[0] {
			return vv, DimensionError{
				Operation: "Multiply",
				Arg1:      vv,
				Arg2:      rightAsVE,
			}
		}
	}

	switch right := rightIn.(type) {
	case KVector:
		//KVector must be a scalar.
		rightAsVD := mat.VecDense(right)
		k0 := rightAsVD.AtVec(0)

		// Multiply all elements of vector with this.
		var prod VectorLinearExpr
		prod.X = vv.Copy()

		prod.L = ZerosMatrix(vv.Len(), vv.Len())
		tempIdentity := Identity(vv.Len())
		prod.L.Scale(k0, &tempIdentity)

		prod.C = ZerosVector(vv.Len())

		return prod, nil

	case KVectorTranspose:
		// KVector must be a scalar. Do the same thing as for KVector.
		return vv.Multiply(right.Transpose())

	default:
		return vv, fmt.Errorf(
			"The input to VarVector's Multiply() method (%v) has unexpected type: %T",
			right, rightIn,
		)
	}
}

/*
LessEq
Description:

	This method creates a less than or equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vv VarVector) LessEq(rhs interface{}) (VectorConstraint, error) {
	return vv.Comparison(rhs, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	This method creates a greater than or equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vv VarVector) GreaterEq(rhs interface{}) (VectorConstraint, error) {
	return vv.Comparison(rhs, SenseGreaterThanEqual)
}

/*
Eq
Description:

	This method creates an equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vv VarVector) Eq(rhs interface{}) (VectorConstraint, error) {
	// Constants

	// Algorithm
	return vv.Comparison(rhs, SenseEqual)

}

/*
Comparison
Description:

	This method creates a constraint of type sense between
	the receiver (as left hand side) and rhs (as right hand side) if both are valid.
*/
func (vv VarVector) Comparison(rhs interface{}, sense ConstrSense) (VectorConstraint, error) {
	// Constants

	// Algorithm
	switch rhsConverted := rhs.(type) {
	case KVector:
		// Check length of input and output.
		if vv.Len() != rhsConverted.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two inputs to comparison '%v' must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					sense,
					vv.Len(),
					rhsConverted.Len(),
				)
		}
		return VectorConstraint{vv, rhsConverted, sense}, nil
	case mat.VecDense:
		rhsAsKVector := KVector(rhsConverted)

		return vv.Comparison(rhsAsKVector, sense)

	case KVectorTranspose:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare VarVector with a transposed vector %v (%T); Try transposing one or the other!",
				rhsConverted, rhsConverted,
			)

	case VarVector:
		// Check length of input and output.
		if vv.Len() != rhsConverted.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two inputs to comparison '%v' must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					sense,
					vv.Len(),
					rhsConverted.Len(),
				)
		}
		// Do Computation
		return VectorConstraint{vv, rhsConverted, sense}, nil

	case VarVectorTranspose:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare VarVector with a transposed vector %v (%T); Try transposing one or the other!",
				rhsConverted, rhsConverted,
			)

	case VectorLinearExpr:
		// Cast type
		rhsAsVLE, _ := rhs.(VectorLinearExpr)

		// Do computation
		return rhsAsVLE.Comparison(vv, sense)

	case VectorLinearExpressionTranspose:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare VarVector with a transposed vector %v (%T); Try transposing one or the other!",
				rhsConverted, rhsConverted,
			)

	default:
		return VectorConstraint{}, fmt.Errorf("The Eq() method for VarVector is not implemented yet for type %T!", rhs)
	}
}

func (vv VarVector) Copy() VarVector {
	// Constants

	// Algorithm
	newVarSlice := []Variable{}
	for varIndex := 0; varIndex < vv.Len(); varIndex++ {
		// Append to newVar Slice
		newVarSlice = append(newVarSlice, vv.Elements[varIndex])
	}

	return VarVector{newVarSlice}

}

/*
Transpose
Description:

	This method creates the transpose of the current vector and returns it.
*/
func (vv VarVector) Transpose() VectorExpression {
	vvCopy := vv.Copy()
	return VarVectorTranspose(vvCopy)
}

/*
Dims
Description:

	Dimensions of the variable vector.
*/
func (vv VarVector) Dims() []uint {
	return []uint{uint(vv.Len()), 1}
}
