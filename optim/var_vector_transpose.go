package optim

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
)

// VarVectorTranspose represents a transposed vector of all optimization variables.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
type VarVectorTranspose struct {
	Elements []Variable
}
// Length returns the length of the vector of optimization variables.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) Length() int {
	return len(vvt.Elements)
}

// Len returns the length of the vector of optimization variables.
// This mirrors the GoNum Vector API and does the same thing as Length.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) Len() int {
	return vvt.Length()
}

// AtVec mirrors the gonum API for vectors and extracts the element of the
// variable vector at the given index.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) AtVec(idx int) ScalarExpression {
	// Constants

	// Algorithm
	return vvt.Elements[idx]
}

// IDs returns the unique variable IDs in the transposed variable vector.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) IDs() []uint64 {
	// Algorithm
	var IDSlice []uint64

	for _, elt := range vvt.Elements {
		IDSlice = append(IDSlice, elt.ID)
	}

	return Unique(IDSlice)

}

// NumVars returns the number of unique variables inside the transposed variable vector.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) NumVars() int {
	return len(vvt.IDs())
}

// Constant returns an all-zeros vector as the constant component of the expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) Constant() mat.VecDense {
	zerosOut := ZerosVector(vvt.Len())
	return zerosOut
}

// LinearCoeff returns the matrix which is multiplied by Variables to get the
// current expression. For a single vector, this is an identity matrix.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) LinearCoeff() mat.Dense {
	return Identity(vvt.Len())
}

// Plus computes the addition of the receiver VarVectorTranspose with the
// incoming vector expression eIn.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) Plus(eIn interface{}, errors ...error) (Expression, error) {
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

// Multiply performs multiplication of a VarVectorTranspose with another expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) Multiply(e interface{}, errors ...error) (Expression, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return vvt, err
	}

	if IsVectorExpression(e) {
		// Check dimensions
		e2, _ := ToVectorExpression(e)
		if vvt.Dims()[1] != e2.Dims()[0] {
			return vvt, DimensionError{
				Arg1:      vvt,
				Arg2:      e2,
				Operation: "Multiply",
			}
		}
	}

	//if symbolic.IsMatrixExpression(e) {
	//	// Check Dimensions
	//	e2, _ := symbolic.ToMatrixExpression(e)
	//	if vvt.Dims()[1] != e2.Dims()[0] {
	//		return vvt, DimensionError{
	//			Arg1:      vvt,
	//			Arg2:      e2,
	//			Operation: "Multiply",
	//		}
	//	}
	//}

	// Multiply Algorithms
	switch right := e.(type) {
	case float64:
		// Create scaled identity matrix
		I := Identity(vvt.Len())
		var scaledI mat.Dense
		scaledI.Scale(right, &I)

		// Copy vvt
		vvtCopy := vvt.Copy()

		return VectorLinearExpressionTranspose{
			L: scaledI, X: vvtCopy.Transpose().(VarVector), C: ZerosVector(vvt.Len()),
		}, nil

	case K:
		return vvt.Multiply(float64(right))

	case mat.VecDense:
		// Convert to KVector
		eAsKVector := KVector(right)
		return vvt.Multiply(eAsKVector)

	case KVector:
		// Collect Unique Variables
		vv := VarVector{UniqueVars(vvt.Elements)}

		// Assemble the vector used in the linear expression.
		L := ZerosVector(vv.Len())
		eLen := right.Len()
		for kvIndex := 0; kvIndex < eLen; kvIndex++ {
			// Get coefficient and variable at kvIndex
			kv_i := right.AtVec(kvIndex)
			vvt_i := vvt.AtVec(kvIndex)

			indexOfvvt_i, _ := FindInSlice(vvt_i.(Variable), vv.Elements)
			L.SetVec(indexOfvvt_i, float64(kv_i.(K))+L.AtVec(indexOfvvt_i))
		}
		return ScalarLinearExpr{L: L, X: vv, C: 0}, nil

	case KVectorTranspose:
		// This is only valid if vvt is of length 1.
		kv := right.Transpose().(KVector)

		prod := ScalarLinearExpr{
			L: mat.VecDense(kv),
			X: vvt.Copy().Transpose().(VarVector),
			C: 0.0,
		}

		return prod, nil

	case VectorLinearExpr:
		// Collect Unique Variables
		newVV := VarVector{UniqueVars(append(vvt.Elements, right.X.Elements...))}

		// Rewrite VLE in terms of newVV
		newVLE := right.RewriteInTermsOf(newVV)

		prod0, _ := vvt.Multiply(newVLE.L)
		prod1, _ := prod0.(VectorLinearExpressionTranspose).Multiply(vvt.Transpose())
		prod := prod1

		// Add in Elements for product of vvt with constant C
		prod, _ = prod.(ScalarQuadraticExpression).Plus(
			vvt.Multiply(KVector(right.C)),
		)

		return prod, nil

	case mat.Dense:
		nr, nc := right.Dims()

		// Create product
		var prod VectorLinearExpressionTranspose
		prod.L = ZerosMatrix(nc, nr)
		for rowIndex := 0; rowIndex < nc; rowIndex++ {
			for colIndex := 0; colIndex < nr; colIndex++ {
				prod.L.Set(
					rowIndex, colIndex,
					right.At(colIndex, rowIndex),
				)
			}
		}

		// Create X
		x := vvt.Copy().Transpose()
		prod.X = x.(VarVector) // Convert to VarVector

		// Create C
		prod.C = ZerosVector(nc)

		return prod, nil

	//case symbolic.Constant:
	//	return vvt.Multiply(mat.Dense(right))
	default:
		return vvt, fmt.Errorf(
			"The input to VarVectorTranspose's Multiply() method (%v) has unexpected type: %T.",
			right, e,
		)
	}
}

// LessEq creates a less than or equal to vector constraint using the receiver
// as the left hand side and the input rhs as the right hand side.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) LessEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return vvt.Comparison(rightIn, SenseLessThanEqual, errors...)
}

// GreaterEq creates a greater than or equal to vector constraint using the
// receiver as the left hand side and the input rhs as the right hand side.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) GreaterEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return vvt.Comparison(rightIn, SenseGreaterThanEqual, errors...)
}

// Eq creates an equal to vector constraint using the receiver as the left hand
// side and the input rhs as the right hand side.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) Eq(rightIn interface{}, errors ...error) (Constraint, error) {
	return vvt.Comparison(rightIn, SenseEqual, errors...)

}

// Comparison creates a constraint of type sense between the receiver (as left
// hand side) and rhs (as right hand side).
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) Comparison(rightIn interface{}, sense ConstrSense, errors ...error) (Constraint, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return VectorConstraint{}, err
	}

	// Algorithm
	switch rhs0 := rightIn.(type) {
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
				"Cannot compare VarVectorTranspose with a normal vector %v (%T); Try transposing one or the other!",
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
				"cannot compare VarVectorTranspose with a normal vector %v (%T); try transposing one or the other!",
				rhs0, rhs0,
			)

	case VectorLinearExpressionTranspose:
		// Do computation
		return rhs0.Comparison(vvt, sense)

	default:
		return VectorConstraint{}, fmt.Errorf("The Eq() method for VarVectorTranspose is not implemented yet for type %T!", rhs0)
	}
}

// Copy creates a copy of the VarVectorTranspose.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) Copy() VarVectorTranspose {
	// Constants

	// Algorithm
	var newVarSlice []Variable
	for varIndex := 0; varIndex < vvt.Len(); varIndex++ {
		// Append to newVar Slice
		newVarSlice = append(newVarSlice, vvt.Elements[varIndex])
	}

	return VarVectorTranspose{newVarSlice}

}

// Transpose creates the transpose of the current vector and returns it.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) Transpose() Expression {
	vvtCopy := vvt.Copy()
	return VarVector{vvtCopy.Elements}
}

// Dims returns the dimension of the VarVectorTranspose object.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) Dims() []int {
	return []int{1, vvt.Len()}
}

// Check checks whether or not the VarVectorTranspose has a sensible
// initialization, returning an error if any element is not properly defined.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) Check() error {
	// Check that each variable is properly defined
	for ii, element := range vvt.Elements {
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

// ToSymbolic converts the VarVectorTranspose to a symbolic expression
// (i.e., an expression made using SymbolicMath.go).
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vvt VarVectorTranspose) ToSymbolic() (symbolic.Expression, error) {
	// Input Processing
	err := vvt.Check()
	if err != nil {
		return nil, err
	}

	// Constants
	vm := symbolic.VariableMatrix{}

	// Algorithm
	vm = append(vm, make([]symbolic.Variable, vvt.Len()))
	for ii, v := range vvt.Elements {
		// Convert to symbolic
		tempV, err := v.ToSymbolic()
		if err != nil {
			return nil, err
		}
		vm[0][ii] = tempV.(symbolic.Variable)
	}

	// Return
	return vm, nil

}
