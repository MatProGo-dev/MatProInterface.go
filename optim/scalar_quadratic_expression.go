package optim

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
)

/*
scalar_quadratic_expression.go
Description:
	Defines some of the functions necessary to define polynomial expressions in terms of the variables
	of an optimization problem.
*/

// Type Definitions
// ================

// ScalarQuadraticExpression represents a quadratic expression of optimization
// variables written as:
//
//	x' * Q * x + L * x + C
//
// where Q is the quadratic term matrix, L is the linear term vector, and C is
// the constant term.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
type ScalarQuadraticExpression struct {
	Q mat.Dense    // Quadratic Term
	L mat.VecDense // Linear Term
	C float64      // Constant Term
	X VarVector
}

// Member Functions
// ================

// NewQuadraticExpr_qb0 returns a basic ScalarQuadraticExpression with only
// the matrix Q being defined; all other values are assumed to be zero.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func NewQuadraticExpr_qb0(QIn mat.Dense, xIn VarVector) (ScalarQuadraticExpression, error) {
	// Constants
	numXIndices := xIn.Len()

	// Input Checking

	// Algorithm
	var qZero []float64
	for qInd := 0; qInd < numXIndices; qInd++ {
		qZero = append(qZero, 0.0)
	}
	q := mat.NewVecDense(numXIndices, qZero)

	return NewQuadraticExpr(QIn, *q, 0.0, xIn)
}

// NewQuadraticExpr returns a ScalarQuadraticExpression defined by QIn, qIn, bIn,
// and xIn.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func NewQuadraticExpr(QIn mat.Dense, qIn mat.VecDense, bIn float64, xIn VarVector) (ScalarQuadraticExpression, error) {
	// Constants

	// Input Checking
	tempExpr := ScalarQuadraticExpression{
		Q: QIn,
		L: qIn,
		C: bIn,
		X: xIn,
	}

	if err := tempExpr.Check(); err != nil {
		return tempExpr, err
	}

	// Algorithm

	return tempExpr, nil
}

// Check verifies the dimensions of all members of the quadratic expression,
// ensuring they have compatible dimensions.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) Check() error {
	// Make the number of elements in q be the dimension of the x in the expression.
	xLen := qe.X.Len()
	n_Q_rows, n_Q_cols := qe.Q.Dims()

	// Check Number of Rows in Q
	if n_Q_rows != xLen {
		return fmt.Errorf("The number of indices was %v which did not match the number of rows in QIn (%v).", xLen, n_Q_rows)
	}

	if n_Q_cols != xLen {
		return fmt.Errorf("The number of indices was %v which did not match the number of columns in QIn (%v).", xLen, n_Q_cols)
	}

	// Otherwise, return no errors.
	return nil
}

// Variables returns a slice containing all unique variables in the expression qe.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) Variables() []Variable {
	return UniqueVars(qe.X.Elements)
}

// NumVars returns the number of unique variables in the expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) NumVars() int {

	return len(qe.IDs())
}

// IDs returns the ids of all of the variables in the quadratic expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) IDs() []uint64 {
	return qe.X.IDs()
}

// Coeffs returns the slice of all coefficient values for each pair of variable
// tuples in the quadratic expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) Coeffs() []float64 {
	// Create container for all coefficients
	var coefficientList []float64

	// Consider all pairs of indices in x.
	var xPairs [][2]uint64
	for vIIndex, varIndex := range qe.X.IDs() {
		for vIIndex2, varIndex2 := range qe.X.IDs() {
			// Save pairs of indices and the associated coefficients
			xPairs = append(xPairs, [2]uint64{varIndex, varIndex2})

			coefficientList = append(coefficientList, qe.Q.At(vIIndex, vIIndex2))
		}
	}

	// Include the elements of L
	for vIIndex, _ := range qe.X.IDs() {
		coefficientList = append(coefficientList, qe.L.AtVec(vIIndex))
	}

	// Include the element of C
	coefficientList = append(coefficientList, qe.C)

	return coefficientList
}

// Constant returns the constant value associated with the quadratic expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) Constant() float64 {
	return qe.C
}

// Plus adds the quadratic expression to another expression and returns the result.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) Plus(e interface{}, errors ...error) (Expression, error) {
	// Constants

	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return qe, err
	}

	// Algorithm depends
	switch rhs := e.(type) {
	case float64:
		// Call the version of this function for K
		return qe.Plus(K(rhs), errors...)
	case K:
		// Get copy of qe
		var newQExpr ScalarQuadraticExpression = qe

		// Add to constant factor
		newQExpr.C += float64(rhs)

		return newQExpr, nil
	case Variable:
		return rhs.Plus(qe)

	case ScalarQuadraticExpression:
		var newQExpr ScalarQuadraticExpression = qe // get copy of e

		// Get Combined set of Variables
		newX := UniqueVars(append(newQExpr.X.Elements, rhs.X.Elements...))
		newQExprAligned, _ := newQExpr.RewriteInTermsOf(VarVector{newX})
		quadraticEInAligned, _ := rhs.RewriteInTermsOf(VarVector{newX})

		// Add matrices together
		var tempSum mat.Dense
		tempSum.Add(&newQExprAligned.Q, &quadraticEInAligned.Q)
		newQExprAligned.Q = tempSum

		// Add vectors together
		//var tempVecSum mat.VecDense
		//tempVecSum.AddVec(&newQExprAligned.L, &quadraticEInAligned.L)
		newQExprAligned.L.AddVec(&newQExprAligned.L, &quadraticEInAligned.L)

		// Add constants together
		newQExprAligned.C += quadraticEInAligned.C
		return newQExprAligned, nil

	case ScalarLinearExpr:
		// Collect Expressions
		var newQExpr ScalarQuadraticExpression = qe // get copy of e

		// Get Combined set of Variables
		newX := UniqueVars(append(newQExpr.X.Elements, rhs.X.Elements...))
		newQExprAligned, _ := newQExpr.RewriteInTermsOf(VarVector{newX})
		linearEInAligned, _ := rhs.RewriteInTermsOf(VarVector{newX})

		// Add linear vector together with the quadratic expression
		newQExprAligned.L.AddVec(&newQExprAligned.L, &linearEInAligned.L)

		// Add constants together
		newQExprAligned.C += rhs.C
		return newQExprAligned, nil
	default:
		return ScalarQuadraticExpression{}, fmt.Errorf("Unexpected type (%T) given as argument to Plus: %v.", e, e)
	}

}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) LessEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return qe.Comparison(rightIn, SenseLessThanEqual, errors...)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) GreaterEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return qe.Comparison(rightIn, SenseGreaterThanEqual, errors...)
}

// Eq returns an equality (==) constraint between the current expression
// and another.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) Eq(rightIn interface{}, errors ...error) (Constraint, error) {
	return qe.Comparison(rightIn, SenseEqual, errors...)
}

// Comparison compares the receiver with expression rhs in the sense provided by sense.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) Comparison(rhsIn interface{}, sense ConstrSense, errors ...error) (Constraint, error) {
	// Input Processing
	rhs, err := ToScalarExpression(rhsIn)
	if err != nil {
		return ScalarConstraint{}, err
	}

	err = CheckErrors(errors)
	if err != nil {
		return ScalarConstraint{}, err
	}

	return ScalarConstraint{qe, rhs, sense}, nil
}

// RewriteInTermsOf rewrites the current quadratic expression in terms of the
// new variables newX.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) RewriteInTermsOf(newX VarVector) (ScalarQuadraticExpression, error) {
	// Create new Quadratic Expression
	// ===============================

	// Find length of X indices
	dimX := newX.Len()

	// Create Q matrix of appropriate dimension.
	newQ := ZerosMatrix(dimX, dimX)

	// Create expression
	var newQE ScalarQuadraticExpression = ScalarQuadraticExpression{
		Q: newQ,
		X: newX,
		L: ZerosVector(dimX),
		C: 0.0,
	}

	// Populate Q
	for oi1Index, oldElt1 := range qe.X.Elements {
		for oi2Index, oldElt2 := range qe.X.Elements {
			// Identify what term is associated with the pair (oldIndex1, oldIndex2)
			oldQterm := qe.Q.At(oi1Index, oi2Index)

			// Get the new indices corresponding to oi1 and oi2
			ni1Index, _ := FindInSlice(oldElt1, newX.Elements)
			if ni1Index == -1 {
				return newQE, fmt.Errorf("The element %v was found in the old X indices, but it does not exist in the new ones!", oldElt1)
			}

			ni2Index, _ := FindInSlice(oldElt2, newX.Elements)
			if ni2Index == -1 {
				return newQE, fmt.Errorf("The element %v was found in the old X indices, but it does not exist in the new ones!", oldElt2)
			}

			// Plug the oldQterm into newQ
			newQE.Q.Set(ni1Index, ni2Index, oldQterm)
		}
	}

	// Create L matrix of appropriate dimension
	newL := ZerosVector(dimX)

	// Populate L
	for oi1Index, oldElt1 := range qe.X.Elements {
		// Identify what term is associated with the pair (oldIndex1, oldIndex2)
		oldLterm := qe.L.AtVec(oi1Index)

		// Get the new indices corresponding to oi1 and oi2
		ni1Index, _ := FindInSlice(oldElt1, newX.Elements) // No error handling or bad indexes should happen at this point.

		// Plug the oldQterm into newQ
		offset := ZerosVector(dimX)
		offset.SetVec(ni1Index, oldLterm)
		(&newL).AddVec(&newL, &offset)
	}
	newQE.L = newL

	// Populate C
	newQE.C = qe.C

	return newQE, nil

}

// Multiply multiplies the current expression to another and returns the
// resulting expression.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) Multiply(val interface{}, errors ...error) (Expression, error) {
	// Input Processing
	if len(errors) > 0 {
		if errors[0] != nil {
			return qe, errors[0]
		}
	}

	// Create Output
	switch valAsType := val.(type) {
	case float64:
		// Algorithm
		var newQE ScalarQuadraticExpression = ScalarQuadraticExpression{
			X: (qe).X,
			C: qe.C,
		}

		// Iterate through all of the rows and columns of Q
		newQE.Q.Scale(valAsType, &qe.Q)

		// Iterate through the linear coefficients
		newQE.L.ScaleVec(valAsType, &qe.L)

		// Update through the constant
		newQE.C *= valAsType

		return newQE, nil
	case K:
		// Algorithm
		valAsFloat := float64(valAsType)
		return qe.Multiply(valAsFloat)

	case Variable:
		// Return error
		return qe, fmt.Errorf("Attempted to multiply Variable with ScalarQuadraticExpression which would result in degree 3 expression! MatProInterface can not currently handle such a high degree polynomial!")

	case ScalarLinearExpr:
		// Return error
		return qe, fmt.Errorf("Attempted to multiply ScalarLinearExpr with ScalarQuadraticExpression which would result in degree 3 expression! MatProInterface can not currently handle such a high degree polynomial!")

	case ScalarQuadraticExpression:
		// Return error
		return qe, fmt.Errorf("Attempted to multiply ScalarQuadraticExpression with ScalarQuadraticExpression which would result in degree 4 expression! MatProInterface can not currently handle such a high degree polynomial!")

	default:
		return qe, fmt.Errorf("Unexpected type of input to Multiply(): %T", val)
	}
}

// Dims returns the dimensions of the ScalarQuadraticExpression, which is always [1, 1].
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) Dims() []int {
	return []int{1, 1}
}

// Transpose returns the transpose of the ScalarQuadraticExpression, which is
// the expression itself.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) Transpose() Expression {
	return qe
}

// ToSymbolic converts the quadratic expression into a symbolic expression
// (i.e., one that uses the symbolic math toolbox).
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (qe ScalarQuadraticExpression) ToSymbolic() (symbolic.Expression, error) {
	// Input Checking
	err := qe.Check()
	if err != nil {
		return nil, err
	}

	// Convert Q, L and C to symbolic
	symQ := symbolic.DenseToKMatrix(qe.Q)
	symL := symbolic.VecDenseToKVector(qe.L)
	symC := symbolic.K(qe.C)

	// Comvert X to symbolic
	symXExpr, err := qe.X.ToSymbolic()
	if err != nil {
		return nil, err
	}

	symX, ok := symXExpr.(symbolic.VariableVector)
	if !ok {
		return nil, fmt.Errorf("Could not convert X to symbolic.VariableVector.")
	}

	// Perform Multplications in Symbolic
	quadraticTerm := symX.Transpose().Multiply(symQ).Multiply(symX)
	fmt.Println(quadraticTerm)
	fmt.Println(symX.Transpose().Multiply(symQ))
	linearTerm := symL.Transpose().Multiply(symX)
	fmt.Println(linearTerm)

	// Sum all terms together and return it
	return quadraticTerm.Plus(linearTerm).Plus(symC), nil
}
