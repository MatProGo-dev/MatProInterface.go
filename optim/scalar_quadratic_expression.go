package optim

import (
	"fmt"
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

/*
QuadraticExpr
Description:

	A quadratic expression of optimization variables (given by their indices).
	The quadratic expression object defines a quadratic written as follows:
		x' * Q * x + L * x + C
*/
type ScalarQuadraticExpression struct {
	Q mat.Dense    // Quadratic Term
	L mat.VecDense // Linear Term
	C float64      // Constant Term
	X VarVector
}

// Member Functions
// ================

/*
NewQuadraticExpr_qb0
Description:

	NewQuadraticExpr_q0 returns a basic Quadratic expression with only the matrix Q being defined,
	all other values are assumed to be zero.
*/
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

/*
NewQuadraticExpr
Description:

	NewQuadraticExpr returns a basic Quadratic expression whuch is defined by QIn, qIn and bIn.
*/
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

/*
Check
Description:

	This function checks the dimensions of all of the members of the quadratic expression which are slices.
	They should have compatible dimensions.
*/
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

/*
Variables
Description:

	This function returns a slice containing all unique variables in the expression qe.
*/
func (qe ScalarQuadraticExpression) Variables() []Variable {
	return UniqueVars(qe.X.Elements)
}

/*
NumVars
Description:

	Returns the number of variables in the expression.
	To make this actually meaningful, we only count the unique vars.
*/
func (qe ScalarQuadraticExpression) NumVars() int {

	return len(qe.IDs())
}

/*
Vars
Description:

	Returns the ids of all of the variables in the quadratic expression.
*/
func (qe ScalarQuadraticExpression) IDs() []uint64 {
	return qe.X.IDs()
}

/*
Coeffs
Description:

	Returns the slice of all coefficient values for each pair of variable tuples.
	The coefficients of the quadratic expression are created in an ordering that comes from the following vector.

	Consider xI (the indices of the input expression e). The output coefficients will be c.
	The coefficients of the expression
		e = x' Q x + q' * x + b
	will be
		e = c' mx + b
	where
		mx = [ x[0]*x[0], x[0]*x[1], ... , x[0]*x[N-1], x[1]*x[1] , x[1]*x[2], ... , x[1]*x[N-1], x[2]*x[2], ... , x[N-1]*x[N-1], x[0], x[1], ... , x[N-1] ]
*/
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

/*
Constant
Description:

	Returns the constant value associated with a quadratic expression.
*/
func (qe ScalarQuadraticExpression) Constant() float64 {
	return qe.C
}

/*
Plus
Description:

	Adds a quadratic expression to either:
	- A Quadratic Expression,
	- A Linear Expression, or
	- A Constant
*/
func (qe ScalarQuadraticExpression) Plus(e interface{}, errors ...error) (ScalarExpression, error) {
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

/*
LessEq
Description:

	LessEq returns a less than or equal to (<=) constraint between the
	current expression and another
*/
func (qe ScalarQuadraticExpression) LessEq(rhsIn interface{}, errors ...error) (ScalarConstraint, error) {
	// Input Processing
	return qe.Comparison(rhsIn, SenseLessThanEqual, errors...)
}

/*
GreaterEq
Description:

	GreaterEq returns a greater than or equal to (>=) constraint between the
	current expression and another
*/
func (qe ScalarQuadraticExpression) GreaterEq(rhsIn interface{}, errors ...error) (ScalarConstraint, error) {
	return qe.Comparison(rhsIn, SenseGreaterThanEqual, errors...)
}

/*
Eq
Description:

	Form an equality constraint with this equality constraint and another
	Eq returns an equality (==) constraint between the current expression
	and another
*/
func (qe ScalarQuadraticExpression) Eq(rhsIn interface{}, errors ...error) (ScalarConstraint, error) {
	return qe.Comparison(rhsIn, SenseEqual, errors...)
}

/*
Comparison
Description:

	This method compares the receiver with expression rhs in the sense provided by sense.

Usage:

	constr, err := qe.Comparison(expr1,SenseGreaterThanEqual)
*/
func (qe ScalarQuadraticExpression) Comparison(rhsIn interface{}, sense ConstrSense, errors ...error) (ScalarConstraint, error) {
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

/*
RewriteInTermsOfIndices
Description:

	Rewrites the current quadratic expression in terms of the new variables.

Usage:

	rewrittenQE, err := orignalQE.RewriteInTermsOfIndices(newXIndices1)
*/
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

/*
Multiply
Description:

	Multiply() multiplies the current expression to another and returns the
	resulting expression
*/
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

func (qe ScalarQuadraticExpression) Dims() []uint64 {
	return []uint64{1, 1}
}
