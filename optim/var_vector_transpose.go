package optim

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

/*
var_vector_transpose.go
Description:
	The VarVectorTranspose type will represent a transposed vector of all
	variables.
*/

/*
VarVectorTranspose
Description:

	Represnts a variable in a optimization problem. The variable is
*/
type VarVectorTranspose struct {
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
func (vvt VarVectorTranspose) Length() int {
	return len(vvt.Elements)
}

/*
Len
Description:

	This function is created to mirror the GoNum Vector API. Does the same thing as Length.
*/
func (vvt VarVectorTranspose) Len() int {
	return vvt.Length()
}

/*
At
Description:

	Mirrors the gonum api for vectors. This extracts the element of the variable vector at the index x.
*/
func (vvt VarVectorTranspose) AtVec(idx int) ScalarExpression {
	// Constants

	// Algorithm
	return vvt.Elements[idx]
}

/*
IDs
Description:

	Returns the unique indices
*/
func (vvt VarVectorTranspose) IDs() []uint64 {
	// Algorithm
	var IDSlice []uint64

	for _, elt := range vvt.Elements {
		IDSlice = append(IDSlice, elt.ID)
	}

	return Unique(IDSlice)

}

/*
NumVars
Description:

	The number of unique variables inside the variable vector.
*/
func (vvt VarVectorTranspose) NumVars() int {
	return len(vvt.IDs())
}

/*
Constant
Description:

	Returns an all zeros vector as output from the method.
*/
func (vvt VarVectorTranspose) Constant() mat.VecDense {
	zerosOut := ZerosVector(vvt.Len())
	return zerosOut
}

/*
LinearCoeff
Description:

	Returns the matrix which is multiplied by Variables to get the current "expression".
	For a single vector, this is an identity matrix.
*/
func (vvt VarVectorTranspose) LinearCoeff() mat.Dense {
	return Identity(vvt.Len())
}

/*
Plus
Description:

	This member function computes the addition of the receiver vector var with the
	incoming vector expression ve.
*/
func (vvt VarVectorTranspose) Plus(eIn interface{}, extras ...interface{}) (VectorExpression, error) {
	// Constants
	vvLen := vvt.Len()

	// Processing Extras

	// Algorithm
	switch eAsType := eIn.(type) {
	case KVector:
		return vvt,
			fmt.Errorf(
				"Cannot add VarVectorTranspose to a normal vector %v (%T); Try transposing one or the other!",
				eAsType, eAsType,
			)
	case KVectorTranspose:
		// Check Lengths
		if eAsType.Len() != vvt.Len() {
			return VarVectorTranspose{},
				fmt.Errorf(
					"The lengths of two vectors in Plus must match! VarVectorTranspose has dimension %v, KVector has dimension %v",
					vvt.Len(),
					eAsType.Len(),
				)
		}

		// Algorithm
		return VectorLinearExpressionTranspose{
			L: Identity(vvLen), X: vvt.Transpose().(VarVector), C: mat.VecDense(eAsType),
		}, nil
	case mat.VecDense:
		return vvt,
			fmt.Errorf(
				"Cannot add VarVectorTranspose to a normal vector %v (%T); Try transposing one or the other!",
				eAsType, eAsType,
			)

	case VarVector:
		return vvt,
			fmt.Errorf(
				"Cannot add VarVectorTranspose to a normal vector %v (%T); Try transposing one or the other!",
				eAsType, eAsType,
			)

	case VarVectorTranspose:
		// Use VLE based plus
		eAsVLE := VectorLinearExpressionTranspose{
			L: Identity(eAsType.Len()), X: eAsType.Transpose().(VarVector), C: ZerosVector(eAsType.Len()),
		}

		return vvt.Plus(eAsVLE)

	case VectorLinearExpr:
		return vvt,
			fmt.Errorf(
				"Cannot add VarVectorTranspose with a normal vector %v (%T); Try transposing one or the other!",
				eAsType, eAsType,
			)

	case VectorLinearExpressionTranspose:
		return eAsType.Plus(vvt)

	default:
		errString := fmt.Sprintf("Unrecognized expression type %T for addition of VarVectorTranspose vvt.Plus(%v)!", eAsType, eAsType)
		return VarVectorTranspose{}, fmt.Errorf(errString)
	}
}

/*
Mult
Description:

	This member function computest the multiplication of the receiver vector var with some
	incoming vector expression (may result in quadratic?).
*/
func (vvt VarVectorTranspose) Mult(c float64) (VectorExpression, error) {
	return vvt, fmt.Errorf("The Mult() method for VarVectorTranspose is not implemented yet!")
}

/*
Multiply
Description:

	Multiplication of a VarVectorTranspose with another expression.
*/
func (vvt VarVectorTranspose) Multiply(e interface{}, extras ...interface{}) (Expression, error) {
	// Input Processing
	err := CheckExtras(extras)
	if err != nil {
		return vvt, err
	}

	if IsVectorExpression(e) {
		// Check dimensions
		e2, _ := ToVectorExpression(e)
		if e2.Len() != vvt.Len() {
			return vvt, fmt.Errorf(
				"VarVectorTranspose of length %v can not be multiplied with a %T of different length (%v).",
				vvt.Len(),
				e2,
				e2.Len(),
			)
		}
	}

	// Multiply Algorithms
	switch eConverted := e.(type) {
	case KVector:
		// Collect Unique Variables
		vv := VarVector{UniqueVars(vvt.Elements)}

		// Assemble the vector used in the linear expression.
		L := ZerosVector(vv.Len())
		eLen := eConverted.Len()
		for kvIndex := 0; kvIndex < eLen; kvIndex++ {
			// Get coefficient and variable at kvIndex
			kv_i := eConverted.AtVec(kvIndex)
			vvt_i := vvt.AtVec(kvIndex)

			indexOfvvt_i, _ := FindInSlice(vvt_i.(Variable), vv.Elements)
			L.SetVec(indexOfvvt_i, float64(kv_i.(K))+L.AtVec(indexOfvvt_i))
		}
		return ScalarLinearExpr{L: L, X: vv, C: 0}, nil
	default:
		return vvt, fmt.Errorf(
			"The input to VarVectorTranspose's Multiply() method (%v) has unexpected type: %T.",
			eConverted, e,
		)
	}
}

/*
LessEq
Description:

	This method creates a less than or equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vvt VarVectorTranspose) LessEq(rhs interface{}) (VectorConstraint, error) {
	return vvt.Comparison(rhs, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	This method creates a greater than or equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vvt VarVectorTranspose) GreaterEq(rhs interface{}) (VectorConstraint, error) {
	return vvt.Comparison(rhs, SenseGreaterThanEqual)
}

/*
Eq
Description:

	This method creates an equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vvt VarVectorTranspose) Eq(rhs interface{}) (VectorConstraint, error) {
	return vvt.Comparison(rhs, SenseEqual)

}

/*
Comparison
Description:

	This method creates a constraint of type sense between
	the receiver (as left hand side) and rhs (as right hand side) if both are valid.
*/
func (vvt VarVectorTranspose) Comparison(rhs interface{}, sense ConstrSense) (VectorConstraint, error) {
	// Constants

	// Algorithm
	switch rhs0 := rhs.(type) {
	case KVector:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot commpare VarVectorTranspose with a normal vector %v (%T); Try transposing one or the other!",
				rhs0, rhs0,
			)
	case KVectorTranspose:
		// Check length of input and output.
		if vvt.Len() != rhs0.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two inputs to comparison '%v' must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					sense,
					vvt.Len(),
					rhs0.Len(),
				)
		}
		return VectorConstraint{vvt, rhs0, sense}, nil
	case mat.VecDense:
		// Cast Type to KVectorTranspose. Maybe I shouldn't do this?
		return vvt.Comparison(KVectorTranspose(rhs0), sense)

	case VarVector:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot commpare VarVectorTranspose with a normal vector %v (%T); Try transposing one or the other!",
				rhs0, rhs0,
			)

	case VarVectorTranspose:
		// Check length of input and output.
		if vvt.Len() != rhs0.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two inputs to comparison '%v' must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					sense,
					vvt.Len(),
					rhs0.Len(),
				)
		}
		// Do Computation
		return VectorConstraint{vvt, rhs0, sense}, nil

	case VectorLinearExpr:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot commpare VarVectorTranspose with a normal vector %v (%T); Try transposing one or the other!",
				rhs0, rhs0,
			)

	case VectorLinearExpressionTranspose:
		// Do computation
		return rhs0.Comparison(vvt, sense)

	default:
		return VectorConstraint{}, fmt.Errorf("The Eq() method for VarVectorTranspose is not implemented yet for type %T!", rhs)
	}
}

func (vvt VarVectorTranspose) Copy() VarVectorTranspose {
	// Constants

	// Algorithm
	newVarSlice := []Variable{}
	for varIndex := 0; varIndex < vvt.Len(); varIndex++ {
		// Append to newVar Slice
		newVarSlice = append(newVarSlice, vvt.Elements[varIndex])
	}

	return VarVectorTranspose{newVarSlice}

}

/*
Transpose
Description:

	This method creates the transpose of the current vector and returns it.
*/
func (vvt VarVectorTranspose) Transpose() VectorExpression {
	vvtCopy := vvt.Copy()
	return VarVector{vvtCopy.Elements}
}
