package optim

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
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

/*
LessEq
Description:

	This method creates a less than or equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vv VarVector) LessEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return vv.Comparison(rightIn, SenseLessThanEqual, errors...)
}

/*
GreaterEq
Description:

	This method creates a greater than or equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vv VarVector) GreaterEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return vv.Comparison(rightIn, SenseGreaterThanEqual, errors...)
}

/*
Eq
Description:

	This method creates an equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vv VarVector) Eq(rightIn interface{}, errors ...error) (Constraint, error) {
	// Constants

	// Algorithm
	return vv.Comparison(rightIn, SenseEqual, errors...)

}

/*
Comparison
Description:

	This method creates a constraint of type sense between
	the receiver (as left hand side) and rhs (as right hand side) if both are valid.
*/
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
func (vv VarVector) Transpose() Expression {
	vvCopy := vv.Copy()
	return VarVectorTranspose(vvCopy)
}

/*
Dims
Description:

	Dimensions of the variable vector.
*/
func (vv VarVector) Dims() []int {
	return []int{vv.Len(), 1}
}

/*
Check
Description:

	Checks whether or not the VarVector has a sensible initialization.
*/
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

/*
ToSymbolic
Description:

	Converts the variable vector to a symbolic expression.
	(i.e., one that uses the symbolic math toolbox).
*/
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
		eltAsSymExpr, err := elt.ToSymbolic()
		if err != nil {
			return nil, fmt.Errorf(
				"could not convert variable %v to symbolic variable",
				elt,
			)
		}
		eltAsSymVar, ok := eltAsSymExpr.(symbolic.Variable)
		if !ok {
			return nil, fmt.Errorf(
				"could not convert variable %v to symbolic variable",
				elt,
			)
		}
		symVVec = append(symVVec, eltAsSymVar)
	}

	// Return the symbolic vector
	return symVVec, nil
}
