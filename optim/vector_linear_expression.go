package optim

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
)

// VectorLinearExpr represents a linear general expression of the form
//
//	L * x + C
//
// where L is an n x m matrix of coefficients that matches the dimension of x, the vector of variables,
// and C is a constant vector.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
type VectorLinearExpr struct {
	X VarVector
	L mat.Dense // Matrix of coefficients. Should match the dimensions of XIndices
	C mat.VecDense
}

// Check checks to see if the VectorLinearExpr is well-defined.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) Check() error {
	// Extract the dimension of the vector x
	m := vle.X.Length()
	nL, mL := vle.L.Dims()
	nC := vle.C.Len()

	// Compare the length of vector x with the appropriate dimension of L
	if m != mL {
		return fmt.Errorf("Dimensions of L (%v x %v) and x (length %v) do not match appropriately.", nL, mL, m)
	}

	// Compare the size of the matrix L with the vector C that it will be compared to.
	if nC != nL {
		return fmt.Errorf("Dimension of L (%v x %v) and C (length %v) do not match!", nL, mL, nC)
	}

	// If all other checks passed, then the VectorLinearExpression seems valid.
	return nil
}

// IDs returns the MatProInterface ID of each variable in the current vector linear expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) IDs() []uint64 {
	return vle.X.IDs()
}

// NumVars returns the number of unique variables in the current vector linear expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) NumVars() int {
	return len(vle.IDs())
}

// LinearCoeff returns the matrix which is applied as a coefficient to the vector X in the expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) LinearCoeff() mat.Dense {

	return vle.L
}

// Constant returns the vector which is given as an offset vector in the linear
// expression (the C in L*x + C).
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) Constant() mat.VecDense {

	return vle.C
}

// GreaterEq creates a VectorConstraint that declares vle is greater than or
// equal to the value to the right hand side rhs.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) GreaterEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return vle.Comparison(rightIn, SenseGreaterThanEqual, errors...)
}

// LessEq creates a VectorConstraint that declares vle is less than or equal to
// the value to the right hand side rhs.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) LessEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return vle.Comparison(rightIn, SenseLessThanEqual, errors...)
}

// Multiply performs multiplication of a VectorLinearExpr with another expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) Multiply(rightIn interface{}, errors ...error) (Expression, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return vle, err
	}

	// Check Dimensions of the input vector
	if IsVectorExpression(rightIn) {
		rightAsE, _ := ToVectorExpression(rightIn)
		err = CheckDimensionsInMultiplication(vle, rightAsE)
		if err != nil {
			return vle, err
		}
	}

	switch right := rightIn.(type) {
	case float64:
		// Create output
		out := vle.Copy()

		// Multiply the elements of the matrices
		out.L.Scale(right, &vle.L)
		out.C.ScaleVec(right, &vle.C)

		return out, nil

	case K:
		return vle.Multiply(float64(right))

	case mat.VecDense:
		// Send warning until we create matrix type.
		return vle, fmt.Errorf(
			"MatProInterface does not currently support operations that result in matrices! if you want this feature, create an issue!",
		)

	case KVector:
		return vle.Multiply(mat.VecDense(right))

	case KVectorTranspose:
		// Send warning until we create matrix type.
		return vle, fmt.Errorf(
			"dimension mismatch! Cannot multiply KVector with a vector of type %T; Try transposing one or the other!",
			right,
		)

	case VectorLinearExpr:
		// Immediately return error.
		return vle, fmt.Errorf(
			"MatProInterface does not currently support operations that result in matrices! if you want this feature, create an issue!",
		)

	case VectorLinearExpressionTranspose:
		// Send warning until we create matrix type.
		return vle, fmt.Errorf(
			"dimension mismatch! Cannot multiply KVector with a vector of type %T; Try transposing one or the other!",
			right,
		)

	default:
		return vle, fmt.Errorf(
			"The input to VectorLinearExpr's Multiply() method (%v) has unexpected type: %T",
			right, right,
		)
	}
}

// Plus returns an expression which adds the expression e to the vector linear expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) Plus(rightIn interface{}, errors ...error) (Expression, error) {
	// Constants
	vleLen := vle.Len()

	// Input Processing
	err := vle.Check()
	if err != nil {
		return vle, err
	}

	err = CheckErrors(errors)
	if err != nil {
		return vle, err
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err = CheckDimensionsInAddition(vle, rightAsE)
		if err != nil {
			return vle, err
		}
	}

	// Algorithm
	switch right := rightIn.(type) {
	case KVector:
		// Check Length
		if right.Len() != vleLen {
			return vle, fmt.Errorf(
				"The length of input KVector (%v) did not match the length of the VectorLinearExpr (%v).",
				right.Len(),
				vleLen,
			)
		}

		// Algorithm
		vleOut := vle
		tempSum, _ := KVector(vle.C).Plus(right)
		//if err != nil {
		//	return vle,
		//		fmt.Errorf(
		//			"There was an issue computing the sum of a KVector with your VectorLinearExpression: %v",
		//			err,
		//		)
		//}
		KSum, _ := tempSum.(KVector)
		vleOut.C = mat.VecDense(KSum)

		// Return
		return vleOut, nil

	case KVectorTranspose:
		return right,
			fmt.Errorf(
				"Cannot add VectorLinearExpr with a transposed vector %v (%T); Try transposing one or the other!",
				right, right,
			)

	case VarVector:
		eAsVLE := VectorLinearExpr{
			L: Identity(right.Len()),
			X: right,
			C: ZerosVector(right.Len()),
		}

		return vle.Plus(eAsVLE)

	case VarVectorTranspose:
		return right,
			fmt.Errorf(
				"Cannot add VectorLinearExpr with a transposed vector %v (%T); Try transposing one or the other!",
				right, right,
			)

	case VectorLinearExpr:
		// Check Lengths
		if right.Len() != vleLen {
			return vle,
				fmt.Errorf(
					"The length of input VectorLinearExpr (%v) did not match the length of the VectorLinearExpr (%v).",
					right.Len(),
					vleLen,
				)
		}

		// Collect VarVectors from expression and vv
		combinedVV := VarVector{append(vle.X.Elements, right.X.Elements...)}
		uniqueVV := VarVector{UniqueVars(combinedVV.Elements)}

		// Create Placeholder vle
		vleOut := vle.RewriteInTermsOf(uniqueVV)
		eRewrittenVLE := right.RewriteInTermsOf(uniqueVV)

		// Add elements of eRewrittenVLE.L to vleOut.L
		nR, nC := vleOut.L.Dims()
		for rowIndex := 0; rowIndex < nR; rowIndex++ {
			for colIndex := 0; colIndex < nC; colIndex++ {
				vleOut.L.Set(
					rowIndex, colIndex,
					vleOut.L.At(rowIndex, colIndex)+eRewrittenVLE.L.At(rowIndex, colIndex),
				)
			}
		}

		// Add elements of eRewrittenVLE.C to vleOut.C
		for rowIndex := 0; rowIndex < nR; rowIndex++ {
			vleOut.C.SetVec(
				rowIndex,
				vleOut.C.AtVec(rowIndex)+eRewrittenVLE.C.AtVec(rowIndex),
			)
		}

		return vleOut, nil

	case VectorLinearExpressionTranspose:
		return right,
			fmt.Errorf(
				"Cannot add VectorLinearExpr with a transposed vector %v (%T); Try transposing one or the other!",
				right, right,
			)
	default:
		return vle, fmt.Errorf("The addition method has not yet been implemented!")
	}
}

/*
LessEq
Description:

	Returns a constraint between the current vector linear expression and the input given
	as the right hand side.
*/
//func (v VectorLinearExpr) LessEq(rhsIn interface{}) (VectorConstraint, error) {
//	// Output depends on the input type
//	switch rhsIn.(type) {
//	case K:
//		// Constant on right hand side.
//		rhsK, _ := rhsIn.(K)
//
//		lhsDim, _ := v.L.Dims()
//
//		onesVec := OnesVector(lhsDim)
//		var rhs KVector
//		rhs.ScaleVec(rhsK.float64, onesVec)
//
//		// Create new VectorExpression
//		return VectorConstraint{
//			LeftHandSide:  v,
//			RightHandSide: rhs,
//			Sense:         SenseLessThanEqual,
//		}, nil
//	}
//
//	return nil, fmt.Errorf("Unexpected type of right hand side %v: %T", rhsIn, rhsIn)
//}

// Eq creates a constraint between the current vector linear expression and
// the rhs given by rhs.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) Eq(rightIn interface{}, errors ...error) (Constraint, error) {
	return vle.Comparison(rightIn, SenseEqual, errors...)
}

// Len returns the number of rows in the vector linear expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) Len() int {
	// Constants

	// Algorithm
	return vle.C.Len()
}

// Comparison compares the input vector linear expression with respect to the
// expression rhsIn and the sense senseIn.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) Comparison(rightIn interface{}, sense ConstrSense, errors ...error) (Constraint, error) {
	// Constants

	// Check Input
	err := vle.Check()
	if err != nil {
		return VectorConstraint{}, fmt.Errorf(
			"There was an issue in the provided vector linear expression %v: %v",
			vle, err,
		)
	}

	// Algorithm
	switch rhsConverted := rightIn.(type) {
	case KVector:
		// Check length of input and output.
		if rhsConverted.Len() != vle.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two vector inputs to Comparison() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					vle.Len(),
					rhsConverted.Len(),
				)
		}
		return VectorConstraint{vle, rhsConverted, sense}, nil
	case mat.VecDense:
		return vle.Eq(KVector(rhsConverted))

	case KVectorTranspose:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare VectorLinearExpr with a transposed vector %v (%T); Try transposing one or the other!",
				rhsConverted, rhsConverted,
			)

	case VectorLinearExpr:
		// Check length of input and output.
		if rhsConverted.Len() != vle.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two vector inputs to Eq() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					vle.Len(),
					rhsConverted.Len(),
				)
		}
		return VectorConstraint{vle, rhsConverted, sense}, nil

	case VectorLinearExpressionTranspose:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare VectorLinearExpr with a transposed vector %v (%T); Try transposing one or the other!",
				rhsConverted, rhsConverted,
			)

	case VarVector:
		// Check length of input and output.
		if rhsConverted.Len() != vle.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two vector inputs to Eq() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					vle.Len(),
					rhsConverted.Len(),
				)
		}
		return VectorConstraint{vle, rhsConverted, sense}, nil

	case VarVectorTranspose:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare VectorLinearExpr with a transposed vector %v (%T); Try transposing one or the other!",
				rhsConverted, rhsConverted,
			)

	default:
		return VectorConstraint{},
			fmt.Errorf(
				"The comparison of vector linear expression %v with object of type %T is not currently supported.",
				vle, rightIn,
			)
	}
}

// RewriteInTermsOf rewrites the VectorLinearExpr in terms of a new set of
// variables vv. It assumes vv contains all unique variables and all elements
// of vle.X are in vv.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) RewriteInTermsOf(vv VarVector) VectorLinearExpr {
	// Constants

	// Create new empty vle
	vleOut := VectorLinearExpr{
		L: ZerosMatrix(vle.Len(), vv.Len()),
		X: vv,
		C: vle.C,
	}

	// Create new L
	nR, _ := vleOut.L.Dims()
	for xIndex, tempVar := range vle.X.Elements {
		// Identify new index of x
		xIndexInVV, _ := FindInSlice(tempVar, vleOut.X.Elements)

		// Change all columns
		for rowI := 0; rowI < nR; rowI++ {
			vleOut.L.Set(
				rowI, xIndexInVV,
				vle.L.At(rowI, xIndex),
			)
		}
	}

	// Return new vle
	return vleOut

}

// AtVec returns the scalar expression at the given index idx.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) AtVec(idx int) ScalarExpression {
	// Constants
	Li := vle.L.RowView(idx)
	LiAsVecDense := Li.(*mat.VecDense)

	// Cast
	sleOut := ScalarLinearExpr{
		L: *LiAsVecDense,
		X: vle.X,
		C: vle.C.AtVec(idx),
	}

	return sleOut

}

// Transpose creates the transpose of the current vector and returns it.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) Transpose() Expression {
	return VectorLinearExpressionTranspose{
		L: vle.L,
		X: vle.X.Copy(),
		C: vle.C,
	}
}

// Copy returns a copy of the VectorLinearExpr.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) Copy() VectorLinearExpr {
	// Constants
	nRows := vle.Len()
	nX := vle.X.Len()

	// Create output
	out := VectorLinearExpr{
		L: ZerosMatrix(nRows, nX),
		C: ZerosVector(nRows),
	}
	out.L.Copy(&vle.L)
	out.C.CopyVec(&vle.C)
	out.X = vle.X.Copy()
	return out
}

// Dims returns the dimensions of the VectorLinearExpr.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) Dims() []int {
	return []int{vle.Len(), 1}
}

// ToSymbolic returns the symbolic version of the VectorLinearExpr expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (vle VectorLinearExpr) ToSymbolic() (symbolic.Expression, error) {
	// Input Processing
	err := vle.Check()
	if err != nil {
		return nil, err
	}

	// Constants
	L := symbolic.DenseToKMatrix(vle.L)
	C := symbolic.VecDenseToKVector(vle.C)
	X, err := vle.X.ToSymbolic()
	if err != nil {
		return nil, err
	}

	// Create the symbolic expression
	return L.Multiply(X).Plus(C), nil

}
