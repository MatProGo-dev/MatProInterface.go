package optim

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
)

// VarVector represents a vector of optimization variables.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
type VarVector struct {
	Elements []Variable
}

// Length returns the length of the vector of optimization variables.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) Length() int {
	return len(vv.Elements)
}

// Len returns the length of the vector of optimization variables.
// This mirrors the GoNum Vector API and does the same thing as Length.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) Len() int {
	return vv.Length()
}

// AtVec mirrors the gonum API for vectors and extracts the element of the
// variable vector at the given index.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) AtVec(idx int) ScalarExpression {
	// Constants

	// Algorithm
	return vv.Elements[idx]
}

// IDs returns the unique variable IDs in the variable vector.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) IDs() []uint64 {
	// Algorithm
	var IDSlice []uint64

	for _, elt := range vv.Elements {
		IDSlice = append(IDSlice, elt.ID)
	}

	return Unique(IDSlice)

}

// NumVars returns the number of unique variables inside the variable vector.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) NumVars() int {
	return len(vv.IDs())
}

// Constant returns an all-zeros vector as the constant component of the expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) Constant() mat.VecDense {
	zerosOut := ZerosVector(vv.Len())
	return zerosOut
}

// LinearCoeff returns the matrix which is multiplied by Variables to get the
// current expression. For a single vector, this is an identity matrix.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) LinearCoeff() mat.Dense {
	return Identity(vv.Len())
}

// Plus computes the addition of the receiver VarVector with the incoming
// vector expression e.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) Plus(e interface{}, errors ...error) (Expression, error) {
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

// Mult computes the multiplication of the receiver VarVector with a scalar
// float64 value. This method is not yet implemented.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) Mult(c float64) (VectorExpression, error) {
	return vv, fmt.Errorf("The Mult() method for VarVector is not implemented yet!")
}

// Multiply performs multiplication of a VarVector with another expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) Multiply(rightIn interface{}, errors ...error) (Expression, error) {
	//Input Processing
	err := vv.Check()
	if err != nil {
		return vv, err
	}

	err = CheckErrors(errors)
	if err != nil {
		return vv, err
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err = CheckDimensionsInMultiplication(vv, rightAsE)
		if err != nil {
			return vv, err
		}
	}

	switch right := rightIn.(type) {
	case float64:
		// Multiply all elements of vector with this.
		var prod VectorLinearExpr
		prod.X = vv.Copy()

		prod.L = ZerosMatrix(vv.Len(), vv.Len())
		tempIdentity := Identity(vv.Len())
		prod.L.Scale(right, &tempIdentity)

		prod.C = ZerosVector(vv.Len())
		return prod, nil

	case K:
		return vv.Multiply(float64(right))

	case KVector:
		//KVector must be a scalar.
		rightAsVD := mat.VecDense(right)
		k0 := rightAsVD.AtVec(0)

		return vv.Multiply(k0)

	case KVectorTranspose:
		// KVectorTranspose must be a scalar. Do the same thing as for KVector.
		if right.Len() == 1 {
			rightAsVD := mat.VecDense(right)
			k0 := rightAsVD.AtVec(0)

			return vv.Multiply(k0)
		}
		// Otherwise, throw this error!
		return vv, fmt.Errorf("cannot complete multiplication that will create matrix product! Submit an issue if you want this feature!")

	default:
		return vv, fmt.Errorf(
			"The input to VarVector's Multiply() method (%v) has unexpected type: %T",
			right, rightIn,
		)
	}
}

// LessEq creates a less than or equal to vector constraint using the receiver
// as the left hand side and the input rhs as the right hand side.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) LessEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return vv.Comparison(rightIn, SenseLessThanEqual, errors...)
}

// GreaterEq creates a greater than or equal to vector constraint using the
// receiver as the left hand side and the input rhs as the right hand side.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) GreaterEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return vv.Comparison(rightIn, SenseGreaterThanEqual, errors...)
}

// Eq creates an equal to vector constraint using the receiver as the left hand
// side and the input rhs as the right hand side.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) Eq(rightIn interface{}, errors ...error) (Constraint, error) {
	return vv.Comparison(rightIn, SenseEqual, errors...)

}

// Comparison creates a constraint of type sense between the receiver (as left
// hand side) and rhs (as right hand side).
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) Comparison(rhs interface{}, sense ConstrSense, errors ...error) (Constraint, error) {
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
		return VectorConstraint{}, fmt.Errorf(
			"The VarVector.Comparison (%v) method is not implemented yet for type %T!",
			sense,
			rhs,
		)
	}
}

// Copy creates a copy of the VarVector.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
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

// Transpose creates the transpose of the current vector and returns it.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) Transpose() Expression {
	vvCopy := vv.Copy()
	return VarVectorTranspose(vvCopy)
}

// Dims returns the dimensions of the variable vector.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) Dims() []int {
	return []int{vv.Len(), 1}
}

// Check checks whether or not the VarVector has a sensible initialization,
// returning an error if any element is not properly defined.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) Check() error {
	// Check that each variable is properly defined
	for ii, element := range vv.Elements {
		err := element.Check()
		if err != nil {
			return fmt.Errorf(
				"element %v has an issue: %v",
				ii, err,
			)
		}
	}

	// If nothing was thrown, then return nil!
	return nil
}

// ToSymbolic converts the variable vector to a symbolic expression
// (i.e., one that uses the symbolic math toolbox).
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vv VarVector) ToSymbolic() (symbolic.Expression, error) {
	// Input Checking
	err := vv.Check()
	if err != nil {
		return nil, err
	}

	// Algorithm
	// Create the symbolic vector
	symVVec := symbolic.VariableVector{}

	// Add each variable to the vector
	for _, elt := range vv.Elements {
		eltAsSymExpr, _ := elt.ToSymbolic()                // No errors should occur because elts were checked above
		eltAsSymVar, _ := eltAsSymExpr.(symbolic.Variable) // No errors should occur because elt must be a variable
		symVVec = append(symVVec, eltAsSymVar)
	}

	// Return the symbolic vector
	return symVVec, nil
}
