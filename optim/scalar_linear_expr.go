package optim

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// ScalarLinearExpr represents a linear general expression of the form
//
//	L' * x + C
//
// where L is a vector of coefficients that matches the dimension of x, the vector of variables
// variables and C is a constant
type ScalarLinearExpr struct {
	X VarVector
	L mat.VecDense // Vector of coefficients. Should match the dimensions of XIndices
	C float64
}

// NewLinearExpr returns a new expression with a single additive constant
// value, c, and no variables.
func NewLinearExpr(c float64) ScalarExpression {
	return ScalarLinearExpr{C: c}
}

/*
Variables
Description:

	This function returns a slice containing all unique variables in the linear expression le.
*/
func (sle ScalarLinearExpr) Variables() []Variable {
	return UniqueVars(sle.X.Elements)
}

// NumVars returns the number of variables in the expression
func (sle ScalarLinearExpr) NumVars() int {
	return sle.X.Len()
}

// Vars returns a slice of the Var ids in the expression
func (sle ScalarLinearExpr) IDs() []uint64 {
	return sle.X.IDs()
}

// Coeffs returns a slice of the coefficients in the expression
func (sle ScalarLinearExpr) Coeffs() []float64 {
	var coeffsOut []float64
	for i := 0; i < sle.L.Len(); i++ {
		coeffsOut = append(coeffsOut, sle.L.AtVec(i))
	}
	return coeffsOut
}

// Constant returns the constant additive value in the expression
func (sle ScalarLinearExpr) Constant() float64 {
	return sle.C
}

// Plus adds the current expression to another and returns the resulting
// expression
func (sle ScalarLinearExpr) Plus(e interface{}, extras ...interface{}) (ScalarExpression, error) {
	// Algorithm depends on the type of eIn.
	switch e.(type) {
	case K:
		// Collect Expression
		KIn := e.(K)

		// Create new expression and add to its constant term
		sleOut := sle
		sleOut.C += float64(KIn)

		return sleOut, nil
	case Variable:
		// Collect Expression
		vIn := e.(Variable)
		return vIn.Plus(sle)

	case ScalarLinearExpr:
		// Collect Expressions
		linearEIn := e.(ScalarLinearExpr)

		// Get Combined set of Variables
		newX := UniqueVars(append(sle.X.Elements, linearEIn.X.Elements...))
		newSLEAligned, _ := sle.RewriteInTermsOf(VarVector{newX})
		linearEInAligned, _ := linearEIn.RewriteInTermsOf(VarVector{newX})

		// Create new vector
		var newSLE ScalarLinearExpr = newSLEAligned // get copy of e
		// Add linear vector together with the quadratic expression
		//var vectorSum mat.VecDense
		//vectorSum.AddVec(newQExprAligned.L, linearEInAligned.L)
		(&newSLE.L).AddVec(&newSLE.L, &linearEInAligned.L)

		// Add constants together
		newSLE.C += linearEIn.C
		return newSLE, nil

	case ScalarQuadraticExpression:

		//var newQExpr QuadraticExpr = *qe // get copy of e
		quadraticEIn := e.(ScalarQuadraticExpression)
		//
		//// Get Combined set of Variables
		//newX := UniqueVars(append(newQExpr.X.Elements, quadraticEIn.X.Elements...))
		//newQExprAligned, _ := newQExpr.RewriteInTermsOf(VarVector{newX})
		//quadraticEInAligned, _ := quadraticEIn.RewriteInTermsOf(VarVector{newX})
		//
		//// Add matrices together
		//var tempSum mat.Dense
		//tempSum.Add(&newQExprAligned.Q, &quadraticEInAligned.Q)
		//newQExprAligned.Q = tempSum
		//
		//// Add vectors together
		////var tempVecSum mat.VecDense
		////tempVecSum.AddVec(&newQExprAligned.L, &quadraticEInAligned.L)
		//newQExprAligned.L.AddVec(&newQExprAligned.L, &quadraticEInAligned.L)
		//
		//// Add constants together
		//newQExprAligned.C += quadraticEInAligned.C
		//return newQExprAligned, nil
		return quadraticEIn.Plus(sle)

	default:
		fmt.Println("Unexpected type given to Plus().")

		return ScalarQuadraticExpression{}, fmt.Errorf("Unexpected type (%T) given as first argument to Plus as %v.", e, e)
	}
}

// Mult multiplies the current expression to another and returns the
// resulting expression
func (sle ScalarLinearExpr) Mult(c float64) (ScalarExpression, error) {
	sle.L.ScaleVec(c, &sle.L)
	sle.C *= c

	return sle, nil
}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (sle ScalarLinearExpr) LessEq(other ScalarExpression) (ScalarConstraint, error) {
	return sle.Comparison(other, SenseLessThanEqual)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (sle ScalarLinearExpr) GreaterEq(other ScalarExpression) (ScalarConstraint, error) {
	return sle.Comparison(other, SenseGreaterThanEqual)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (sle ScalarLinearExpr) Eq(other ScalarExpression) (ScalarConstraint, error) {
	return sle.Comparison(other, SenseEqual)
}

/*
Comparison
Description:

	This method compares the receiver with expression rhs in the sense provided by sense.

Usage:

	constr, err := e.Comparison(expr1,SenseGreaterThanEqual)
*/
func (sle ScalarLinearExpr) Comparison(rhs ScalarExpression, sense ConstrSense) (ScalarConstraint, error) {
	return ScalarConstraint{sle, rhs, sense}, nil
}

/*
RewriteInTermsOf
Description:

	Rewrites the current linear expression in terms of the new variables.

Usage:

	rewrittenLE, err := orignalLE.RewriteInTermsOfIndices(newXIndices1)
*/
func (sle ScalarLinearExpr) RewriteInTermsOf(newX VarVector) (ScalarLinearExpr, error) {
	// Create new Linear Express
	var newLE ScalarLinearExpr = ScalarLinearExpr{
		X: newX,
	}

	// Find length of X indices
	dimX := newX.Len()

	// Create L matrix of appropriate dimension
	var newLfloat []float64
	for rowIndex := 0; rowIndex < dimX; rowIndex++ {
		newLfloat = append(newLfloat, 0.0)
	}
	newL := mat.NewVecDense(dimX, newLfloat)

	// Populate L
	for oi1Index, oldElt1 := range sle.X.Elements {
		// Identify what term is associated with the pair (oldIndex1, oldIndex2)
		oldLterm := sle.L.AtVec(oi1Index)

		// Get the new indices corresponding to oi1 and oi2
		ni1Index, err := FindInSlice(oldElt1, newX.Elements)
		if err != nil {
			return newLE, fmt.Errorf("The element %v was found in the old X indices, but it does not exist in the new ones!", oldElt1)
		}
		//newElt1 := newX.Elements[ni1Index]

		// Plug the old Linearterm into newLinear expression
		offset := ZerosVector(dimX)
		offset.SetVec(ni1Index, oldLterm)

		newL.AddVec(newL, &offset)
	}
	newLE.L = *newL

	// Populate C
	newLE.C = sle.C

	return newLE, nil

}
