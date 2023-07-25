package optim

/*
vector_linear_expression.go
Description:

*/

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// VectorLinearExpr represents a linear general expression of the form
//
//	L' * x + C
//
// where L is an n x m matrix of coefficients that matches the dimension of x, the vector of variables
// and C is a constant vector
type VectorLinearExpr struct {
	X VarVector
	L mat.Dense // Matrix of coefficients. Should match the dimensions of XIndices
	C mat.VecDense
}

/*
Check
Description:

	Checks to see if the VectorLinearExpression is well-defined.
*/
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

/*
IDs
Description:

	Returns the MatProInterface ID of each variable in the current vector linear expression.
*/
func (vle VectorLinearExpr) IDs() []uint64 {
	return vle.X.IDs()
}

/*
NumVars
Description:

	Returns the goop2 ID of each variable in the current vector linear expression.
*/
func (vle VectorLinearExpr) NumVars() int {
	return len(vle.IDs())
}

/*
LinearCoeff
Description:

	Returns the matrix which is applied as a coefficient to the vector X in our expression.
*/
func (vle VectorLinearExpr) LinearCoeff() mat.Dense {

	return vle.L
}

/*
Constant
Description:

	Returns the vector which is given as an offset vector in the linear expression represented by v
	(the c in the above expression).
*/
func (vle VectorLinearExpr) Constant() mat.VecDense {

	return vle.C
}

/*
GreaterEq
Description:

	Creates a VectorConstraint that declares vle is greater than or equal to the value to the right hand side rhs.
*/
func (vle VectorLinearExpr) GreaterEq(rhs interface{}) (VectorConstraint, error) {
	return vle.Comparison(rhs, SenseGreaterThanEqual)
}

/*
LessEq
Description:

	Creates a VectorConstraint that declares vle is less than or equal to the value to the right hand side rhs.
*/
func (vle VectorLinearExpr) LessEq(rhs interface{}) (VectorConstraint, error) {
	return vle.Comparison(rhs, SenseLessThanEqual)
}

/*
Mult
Description:

	Returns an expression which scales every dimension of the vector linear expression by the input.
*/
func (vle VectorLinearExpr) Mult(c float64) (VectorExpression, error) {
	return vle, fmt.Errorf("The multiplication method has not yet been implemented!")
}

/*
Plus
Description:

	Returns an expression which adds the expression e to the vector linear expression at hand.
*/
func (vle VectorLinearExpr) Plus(e interface{}, extras ...interface{}) (VectorExpression, error) {
	// Constants
	vleLen := vle.Len()

	// Input Processing

	// Algorithm
	switch eConverted := e.(type) {
	case KVector:
		// Check Length
		if eConverted.Len() != vleLen {
			return vle, fmt.Errorf(
				"The length of input KVector (%v) did not match the length of the VectorLinearExpr (%v).",
				eConverted.Len(),
				vleLen,
			)
		}

		// Algorithm
		vleOut := vle
		tempSum, err := KVector(vle.C).Plus(eConverted)
		if err != nil {
			return vle,
				fmt.Errorf(
					"There was an issue computing the sum of a KVector with your VectorLinearExpression: %v",
					err,
				)
		}
		KSum, _ := tempSum.(KVector)
		vleOut.C = mat.VecDense(KSum)

		// Return
		return vleOut, nil

	case KVectorTranspose:
		return eConverted,
			fmt.Errorf(
				"Cannot add VectorLinearExpr with a transposed vector %v (%T); Try transposing one or the other!",
				eConverted, eConverted,
			)

	case VarVector:
		eAsVLE := VectorLinearExpr{
			L: Identity(eConverted.Len()),
			X: eConverted,
			C: ZerosVector(eConverted.Len()),
		}

		return vle.Plus(eAsVLE)

	case VarVectorTranspose:
		return eConverted,
			fmt.Errorf(
				"Cannot add VectorLinearExpr with a transposed vector %v (%T); Try transposing one or the other!",
				eConverted, eConverted,
			)

	case VectorLinearExpr:
		// Check Lengths
		if eConverted.Len() != vleLen {
			return vle,
				fmt.Errorf(
					"The length of input VectorLinearExpr (%v) did not match the length of the VectorLinearExpr (%v).",
					eConverted.Len(),
					vleLen,
				)
		}

		// Collect VarVectors from expression and vv
		combinedVV := VarVector{append(vle.X.Elements, eConverted.X.Elements...)}
		uniqueVV := VarVector{UniqueVars(combinedVV.Elements)}

		// Create Placeholder vle
		vleOut := vle.RewriteInTermsOf(uniqueVV)
		eRewrittenVLE := eConverted.RewriteInTermsOf(uniqueVV)

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
		return eConverted,
			fmt.Errorf(
				"Cannot add VectorLinearExpr with a transposed vector %v (%T); Try transposing one or the other!",
				eConverted, eConverted,
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

/*
Eq
Description:

	Creates a constraint between the current vector linear expression v and the
	rhs given by rhs.
*/
func (vle VectorLinearExpr) Eq(rhs interface{}) (VectorConstraint, error) {
	return vle.Comparison(rhs, SenseEqual)
}

/*
Len
Description:

	The size of the constraint.
*/
func (vle VectorLinearExpr) Len() int {
	// Constants

	// Algorithm
	return vle.C.Len()
}

/*
Comparison
Description:

	Compares the input vector linear expression with respect to the expression rhsIn and the sense
	senseIn.
*/
func (vle VectorLinearExpr) Comparison(rhs interface{}, sense ConstrSense) (VectorConstraint, error) {
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
	switch rhsConverted := rhs.(type) {
	case KVector:
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
		return VectorConstraint{}, fmt.Errorf("The comparison of vector linear expression %v with object of type %T is not currently supported.", vle, rhs)
	}
}

/*
RewriteInTermsOf
Description:

	Rewrites the VectorLinearExpression in terms of a new set of variables vv

Assumes:

	vv contains all unique variables.
	All elements of vle.X are in vv.
*/
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

/*
AtVec
Description:
*/
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

/*
Transpose
Description:

	This method creates the transpose of the current vector and returns it.
*/
func (vle VectorLinearExpr) Transpose() VectorExpression {
	return VectorLinearExpressionTranspose{
		L: vle.L,
		X: vle.X.Copy(),
		C: vle.C,
	}
}
