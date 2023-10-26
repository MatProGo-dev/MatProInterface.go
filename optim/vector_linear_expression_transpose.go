package optim

/*
vector_linear_expression_transpose.go
Description:

*/

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// VectorLinearExpressionTranspose represents a linear general expression of the form
//
//	x^T * L^T + C^T
//
// where L is an n x m matrix of coefficients that matches the dimension of x, the vector of variables
// and C is a constant vector
type VectorLinearExpressionTranspose struct {
	X VarVector
	L mat.Dense // Matrix of coefficients. Should match the dimensions of XIndices
	C mat.VecDense
}

/*
Check
Description:

	Checks to see if the VectorLinearExpressionTransposeession is well-defined.
*/
func (vle VectorLinearExpressionTranspose) Check() error {
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

	// If all other checks passed, then the VectorLinearExpressionTransposeession seems valid.
	return nil
}

/*
IDs
Description:

	Returns the MatProInterface ID of each variable in the current vector linear expression.
*/
func (vle VectorLinearExpressionTranspose) IDs() []uint64 {
	return vle.X.IDs()
}

/*
NumVars
Description:

	Returns the goop2 ID of each variable in the current vector linear expression.
*/
func (vle VectorLinearExpressionTranspose) NumVars() int {
	return len(vle.IDs())
}

/*
LinearCoeff
Description:

	Returns the matrix which is applied as a coefficient to the vector X in our expression.
*/
func (vle VectorLinearExpressionTranspose) LinearCoeff() mat.Dense {

	return vle.L
}

/*
Constant
Description:

	Returns the vector which is given as an offset vector in the linear expression represented by v
	(the c in the above expression).
*/
func (vle VectorLinearExpressionTranspose) Constant() mat.VecDense {

	return vle.C
}

/*
GreaterEq
Description:

	Creates a VectorConstraint that declares vle is greater than or equal to the value to the right hand side rhs.
*/
func (vle VectorLinearExpressionTranspose) GreaterEq(rhs interface{}) (VectorConstraint, error) {
	return vle.Comparison(rhs, SenseGreaterThanEqual)
}

/*
LessEq
Description:

	Creates a VectorConstraint that declares vle is less than or equal to the value to the right hand side rhs.
*/
func (vle VectorLinearExpressionTranspose) LessEq(rhs interface{}) (VectorConstraint, error) {
	return vle.Comparison(rhs, SenseLessThanEqual)
}

/*
Mult
Description:

	Returns an expression which scales every dimension of the vector linear expression by the input.
*/
func (vlet VectorLinearExpressionTranspose) Mult(c float64) (VectorExpression, error) {
	return vlet, fmt.Errorf("The multiplication method has not yet been implemented!")
}

/*
Multiply
Description:

	Multiplication of a VarVector with another expression.
*/
func (vlet VectorLinearExpressionTranspose) Multiply(e interface{}, errors ...error) (Expression, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return vlet, err
	}

	if IsVectorExpression(e) {
		// Check dimensions
		ve, _ := ToVectorExpression(e)
		if ve.Len() != vlet.Len() {
			return vlet, fmt.Errorf(
				"dimension of VectorLinearExpressionTranspose %v does not match the vector expression's dimension %v",
				vlet.Len(),
				ve.Len(),
			)
		}
	}

	if IsScalarExpression(e) {
		// Check to see if the dimension is 1
		se, _ := ToScalarExpression(e)
		if vlet.Len() != 1 {
			return vlet, DimensionError{
				Operation: "Multiply",
				Arg1:      vlet,
				Arg2:      se,
			}
		}
	}

	switch right := e.(type) {
	case float64:
		eAsK := K(right)
		return eAsK.Multiply(vlet)
	case K:
		return right.Multiply(vlet)
	case Variable:
		return right.Multiply(vlet)
	case VarVector:
		// Compile all of the unique variables
		newX := VarVector{
			UniqueVars(append(vlet.X.Elements, right.Elements...)),
		}

		// Rewrite vlet in terms of these new variables
		newVLET := vlet.RewriteInTermsOf(newX)
		nr_L, _ := newVLET.L.Dims()

		// Multiplication should now be easier?
		Q := ZerosMatrix(vlet.Len(), vlet.Len())
		for rowIndex := 0; rowIndex < nr_L; rowIndex++ {
			for colIndex := 0; colIndex < vlet.Len(); colIndex++ {
				Q.Set(
					rowIndex, colIndex,
					vlet.L.At(rowIndex, colIndex),
				)
			}
		}
		L := ZerosVector(vlet.Len())
		L.CopyVec(&newVLET.C)

		return ScalarQuadraticExpression{
			Q: Q, L: L, C: 0.0, X: newX,
		}, nil

	default:
		return vlet, fmt.Errorf(
			"The input to VarVector's Multiply() method (%v) has unexpected type: %T",
			right, e,
		)
	}
}

/*
Plus
Description:

	Returns an expression which adds the expression e to the vector linear expression at hand.
*/
func (vle VectorLinearExpressionTranspose) Plus(e interface{}, errors ...error) (VectorExpression, error) {
	// Constants
	vleLen := vle.Len()

	// Input Processing

	// Algorithm
	switch eConverted := e.(type) {
	case KVector:
		return eConverted,
			fmt.Errorf(
				"Cannot add VectorLinearExpressioinTranspose with a normal vector %v (%T); Try transposing one or the other!",
				eConverted, eConverted,
			)
	case KVectorTranspose:
		// Check Length
		if eConverted.Len() != vleLen {
			return vle, fmt.Errorf(
				"The length of input KVector (%v) did not match the length of the VectorLinearExpressionTranspose (%v).",
				eConverted.Len(),
				vleLen,
			)
		}

		// Algorithm
		vleOut := vle
		tempSum, _ := KVectorTranspose(vle.C).Plus(eConverted)
		KSum, _ := tempSum.(KVectorTranspose)
		vleOut.C = mat.VecDense(KSum)

		// Return
		return vleOut, nil
	case VarVector:
		return eConverted,
			fmt.Errorf(
				"Cannot add VectorLinearExpressioinTranspose with a normal vector %v (%T); Try transposing one or the other!",
				eConverted, eConverted,
			)

	case VarVectorTranspose:
		eAsVLE := VectorLinearExpressionTranspose{
			L: Identity(eConverted.Len()),
			X: eConverted.Transpose().(VarVector),
			C: ZerosVector(eConverted.Len()),
		}

		return vle.Plus(eAsVLE)

	case VectorLinearExpr:
		return eConverted,
			fmt.Errorf(
				"Cannot add VectorLinearExpressioinTranspose with a normal vector %v (%T); Try transposing one or the other!",
				eConverted, eConverted,
			)

	case VectorLinearExpressionTranspose:
		// Check Lengths
		if eConverted.Len() != vleLen {
			return vle,
				fmt.Errorf(
					"The length of input VectorLinearExpressionTranspose (%v) did not match the length of the VectorLinearExpressionTranspose (%v).",
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
	default:
		return vle, fmt.Errorf(
			"The VectorLinearExpressionTranspose.Plus method has not yet been implemented for type %T!",
			eConverted,
		)
	}
}

/*
LessEq
Description:

	Returns a constraint between the current vector linear expression and the input given
	as the right hand side.
*/
//func (v VectorLinearExpressionTranspose) LessEq(rhsIn interface{}) (VectorConstraint, error) {
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
func (vlet VectorLinearExpressionTranspose) Eq(rhs interface{}) (VectorConstraint, error) {
	return vlet.Comparison(rhs, SenseEqual)
}

/*
Len
Description:

	The size of the constraint.
*/
func (vlet VectorLinearExpressionTranspose) Len() int {
	// Constants

	// Algorithm
	return vlet.C.Len()
}

/*
Comparison
Description:

	Compares the input vector linear expression with respect to the expression rhsIn and the sense
	senseIn.
*/
func (vlet VectorLinearExpressionTranspose) Comparison(rhs interface{}, sense ConstrSense) (VectorConstraint, error) {
	// Constants

	// Check Input
	err := vlet.Check()
	if err != nil {
		return VectorConstraint{}, fmt.Errorf(
			"There was an issue in the provided vector linear expression %v: %v",
			vlet, err,
		)
	}

	// Algorithm
	switch rhsConverted := rhs.(type) {
	case KVector:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare VectorLinearExpressionTranspose with a normal vector %v (%T); Try transposing one or the other!",
				rhsConverted, rhsConverted,
			)
	case KVectorTranspose:
		// Check length of input and output.
		if rhsConverted.Len() != vlet.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two vector inputs to Eq() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					vlet.Len(),
					rhsConverted.Len(),
				)
		}
		return VectorConstraint{vlet, rhsConverted, sense}, nil
	case mat.VecDense:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare VectorLinearExpressionTranspose with a normal vector %v (%T); Try transposing one or the other!",
				rhsConverted, rhsConverted,
			)
	case VarVector:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare VectorLinearExpressionTranspose with a normal vector %v (%T); Try transposing one or the other!",
				rhsConverted, rhsConverted,
			)
	case VarVectorTranspose:
		// Check length of input and output.
		if rhsConverted.Len() != vlet.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two vector inputs to Eq() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					vlet.Len(),
					rhsConverted.Len(),
				)
		}
		return VectorConstraint{vlet, rhsConverted, sense}, nil
	case VectorLinearExpr:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare VectorLinearExpressionTranspose with a normal vector %v (%T); Try transposing one or the other!",
				rhsConverted, rhsConverted,
			)
	case VectorLinearExpressionTranspose:
		// Check length of input and output.
		if rhsConverted.Len() != vlet.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two vector inputs to Eq() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					vlet.Len(),
					rhsConverted.Len(),
				)
		}
		return VectorConstraint{vlet, rhsConverted, sense}, nil

	default:
		return VectorConstraint{}, fmt.Errorf("The comparison of vector linear expression %v with object of type %T is not currently supported.", vlet, rhs)
	}
}

/*
RewriteInTermsOf
Description:

	Rewrites the VectorLinearExpressionTransposeession in terms of a new set of variables vv

Assumes:

	vv contains all unique variables.
	All elements of vle.X are in vv.
*/
func (vlet VectorLinearExpressionTranspose) RewriteInTermsOf(vv VarVector) VectorLinearExpressionTranspose {
	// Constants

	// Create new empty vle
	vletOut := VectorLinearExpressionTranspose{
		L: ZerosMatrix(vlet.Len(), vv.Len()),
		X: vv,
		C: vlet.C,
	}

	// Create new L
	nR, _ := vletOut.L.Dims()
	for xIndex, tempVar := range vlet.X.Elements {
		// Identify new index of x
		xIndexInVV, _ := FindInSlice(tempVar, vletOut.X.Elements)

		// Change all columns
		for rowI := 0; rowI < nR; rowI++ {
			vletOut.L.Set(
				rowI, xIndexInVV,
				vlet.L.At(rowI, xIndex),
			)
		}
	}

	// Return new vle
	return vletOut

}

/*
AtVec
Description:
*/
func (vlet VectorLinearExpressionTranspose) AtVec(idx int) ScalarExpression {
	// Constants
	Li := vlet.L.RowView(idx)
	LiAsVecDense := Li.(*mat.VecDense)

	// Cast
	sleOut := ScalarLinearExpr{
		L: *LiAsVecDense,
		X: vlet.X,
		C: vlet.C.AtVec(idx),
	}

	return sleOut

}

/*
Transpose
Description:

	This method creates the transpose of the current vector and returns it.
*/
func (vlet VectorLinearExpressionTranspose) Transpose() VectorExpression {
	return VectorLinearExpr{
		L: vlet.L,
		X: vlet.X.Copy(),
		C: vlet.C,
	}
}

/*
Dims
Description:

	Returns the dimensions of the VectorLinearExpressionTranspose
	object.
*/
func (vlet VectorLinearExpressionTranspose) Dims() []uint {
	return []uint{1, uint(vlet.Len())}
}
