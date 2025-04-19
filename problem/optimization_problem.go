package problem

import (
	"fmt"

	"github.com/MatProGo-dev/MatProInterface.go/mpiErrors"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	getKVector "github.com/MatProGo-dev/SymbolicMath.go/get/KVector"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
)

// OptimizationProblem represents the overall constrained linear optimization model to be
// solved. OptimizationProblem contains all the variables associated with the optimization
// problem, constraints, objective, and parameters. New variables can only be
// created using an instantiated OptimizationProblem.
type OptimizationProblem struct {
	Name        string
	Variables   []symbolic.Variable
	Constraints []symbolic.Constraint
	Objective   Objective
}

// NewProblem returns a new model with some default arguments such as not to show
// the log and no time limit.
func NewProblem(name string) *OptimizationProblem {
	return &OptimizationProblem{Name: name}
}

/*
AddVariable
Description:

	This method adds an "unbounded" continuous variable to the model.
*/
func (op *OptimizationProblem) AddVariable() symbolic.Variable {
	return op.AddRealVariable()
}

/*
AddRealVariable
Description:

	Adds a Real variable to the model and returns said variable.
*/
func (op *OptimizationProblem) AddRealVariable() symbolic.Variable {
	return op.AddVariableClassic(-optim.INFINITY, optim.INFINITY, symbolic.Continuous)
}

// AddVariable adds a variable of a given variable type to the model given the lower
// and upper value limits. This variable is returned.
func (op *OptimizationProblem) AddVariableClassic(lower, upper float64, vtype symbolic.VarType) symbolic.Variable {
	id := uint64(len(op.Variables))
	newVar := symbolic.Variable{
		ID:    id,
		Lower: lower,
		Upper: upper,
		Type:  vtype,
		Name:  fmt.Sprintf("x_%v", id),
	}
	op.Variables = append(op.Variables, newVar)
	return newVar
}

// AddBinaryVar adds a binary variable to the model and returns said variable.
func (op *OptimizationProblem) AddBinaryVariable() symbolic.Variable {
	return op.AddVariableClassic(0, 1, symbolic.Binary)
}

/*
AddVariableVector
Description:

	Creates a VarVector object using a constructor that assumes you want an "unbounded" vector of real optimization
	variables.
*/
func (op *OptimizationProblem) AddVariableVector(dim int) symbolic.VariableVector {
	// Constants

	// Algorithm
	varSlice := make([]symbolic.Variable, dim)
	for eltIndex := 0; eltIndex < dim; eltIndex++ {
		varSlice[eltIndex] = op.AddVariable()
	}
	return varSlice
}

/*
AddVariableVectorClassic
Description:

	The classic version of AddVariableVector defined in the original goop.
*/
func (op *OptimizationProblem) AddVariableVectorClassic(
	num int, lower, upper float64, vtype symbolic.VarType,
) symbolic.VariableVector {
	stID := uint64(len(op.Variables))
	vs := make([]symbolic.Variable, num)
	for i := range vs {
		vs[i] = symbolic.Variable{
			ID:    stID + uint64(i),
			Lower: lower,
			Upper: upper,
			Type:  vtype,
			Name:  fmt.Sprintf("x_%v", stID+uint64(i)),
		}
	}

	op.Variables = append(op.Variables, vs...)
	return vs
}

// AddBinaryVariableVector adds a vector of binary variables to the model and
// returns the slice.
func (op *OptimizationProblem) AddBinaryVariableVector(num int) symbolic.VariableVector {
	return op.AddVariableVectorClassic(num, 0, 1, symbolic.Binary)
}

// AddVariableMatrix adds a matrix of variables of a given type to the model with
// lower and upper value limits and returns the resulting slice.
func (op *OptimizationProblem) AddVariableMatrix(
	rows, cols int, lower, upper float64, vtype symbolic.VarType,
) symbolic.VariableMatrix {
	// TODO: Add support for adding a variable matrix with a given
	// environment as well as upper and lower bounds.

	// Create variables
	vmOut := symbolic.NewVariableMatrix(rows, cols)

	// Add variables to the problem
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			op.Variables = append(op.Variables, vmOut[i][j])
		}
	}

	return vmOut
}

// AddBinaryVariableMatrix adds a matrix of binary variables to the model and returns
// the resulting slice.
func (op *OptimizationProblem) AddBinaryVariableMatrix(rows, cols int) [][]symbolic.Variable {
	return op.AddVariableMatrix(rows, cols, 0, 1, symbolic.Binary)
}

/*
SetObjective
Description:
	sets the objective of the model given an expression and
	objective sense.
Notes:
	To make this function easier to parse, we will assume an expression
	is given, even though objectives are normally scalars.
*/

func (op *OptimizationProblem) SetObjective(e symbolic.Expression, sense ObjSense) error {
	// Input Processing
	se, err := symbolic.ToScalarExpression(e)
	if err != nil {
		return fmt.Errorf("trouble parsing input expression: %v", err)
	}

	// Return
	op.Objective = *NewObjective(se, sense)
	return nil
}

/*
ToSymbolicConstraint
Description:

	Converts a constraint in the form of a optim.Constraint object into a symbolic.Constraint object.
*/
func ToSymbolicConstraint(inputConstraint optim.Constraint) (symbolic.Constraint, error) {
	// Input Processing

	// Input Processing
	err := inputConstraint.Check()
	if err != nil {
		return symbolic.ScalarConstraint{}, err
	}

	// Convert LHS to symbolic expression
	lhs, _ := inputConstraint.Left().ToSymbolic()

	// Convert RHS to symbolic expression
	rhs, _ := inputConstraint.Right().ToSymbolic()

	// Get Sense
	sense := inputConstraint.ConstrSense().ToSymbolic()

	// Convert
	switch {
	case symbolic.IsScalarExpression(lhs):
		return symbolic.ScalarConstraint{
			LeftHandSide:  lhs.(symbolic.ScalarExpression),
			RightHandSide: rhs.(symbolic.ScalarExpression),
			Sense:         sense,
		}, nil
	case symbolic.IsVectorExpression(lhs):
		return symbolic.VectorConstraint{
			LeftHandSide:  lhs.(symbolic.VectorExpression),
			RightHandSide: rhs.(symbolic.VectorExpression),
			Sense:         sense,
		}, nil
	default:
		return nil, fmt.Errorf("the left hand side of the input (%T) constraint is not recognized.", lhs)
	}

}

/*
From
Description:

	Converts the given input into an optimization problem.
*/
func From(inputModel optim.Model) (*OptimizationProblem, error) {
	// Create a new optimization problem
	newOptimProblem := NewProblem(inputModel.Name)

	// Input Processing
	err := inputModel.Check()
	if err != nil {
		return nil, err
	}

	if inputModel.Obj == nil {
		return nil, fmt.Errorf("the input model has no objective function!")
	}

	// Collect All Variables from Model and copy them into the new optimization
	// problem object.
	for ii, variable := range inputModel.Variables {
		newOptimProblem.Variables = append(newOptimProblem.Variables, symbolic.Variable{
			ID:    uint64(ii),
			Lower: variable.Lower,
			Upper: variable.Upper,
			Type:  symbolic.VarType(variable.Vtype),
		})
	}

	// Collect All Constraints from Model and copy them into the new optimization
	// problem object.
	for ii, constraint := range inputModel.Constraints {
		newConstraint, err := ToSymbolicConstraint(constraint)
		if err != nil {
			return nil, fmt.Errorf(
				"there was a problem creating the %v-th constraint: %v",
				ii,
				err,
			)
		}
		newOptimProblem.Constraints = append(newOptimProblem.Constraints, newConstraint)
	}

	// Convert Objective
	objectiveExpr, err := inputModel.Obj.ToSymbolic()
	if err != nil {
		return nil, err
	}

	err = newOptimProblem.SetObjective(
		objectiveExpr,
		ObjSense(inputModel.Obj.Sense),
	)
	if err != nil {
		return nil, err
	}

	// Done
	return newOptimProblem, nil

}

/*
Check
Description:

	Checks that the OptimizationProblem is valid.
*/
func (op *OptimizationProblem) Check() error {
	// Check Objective
	if op.Objective == (Objective{}) {
		return mpiErrors.NoObjectiveDefinedError{}
	}

	err := op.Objective.Check()
	if err != nil {
		return fmt.Errorf("the objective is not valid: %v", err)
	}

	// Check Variables
	for _, variable := range op.Variables {
		err = variable.Check()
		if err != nil {
			return fmt.Errorf("the variable is not valid: %v", err)
		}
	}

	// Check Constraints
	for _, constraint := range op.Constraints {
		err = constraint.Check()
		if err != nil {
			return fmt.Errorf("the constraint is not valid: %v", err)
		}
	}

	// All Checks Passed!
	return nil
}

/*
IsLinear
Description:

	Checks if the optimization problem is linear.
	Per the definition of a linear optimization problem, the problem is linear if and only if:
	1. The objective function is linear (i.e., a constant or an affine combination of variables).
	2. All constraints are linear (i.e., an affine combination of variables in an inequality or equality).
*/
func (op *OptimizationProblem) IsLinear() bool {
	// Input Processing
	// Verify that the problem is well-formed
	err := op.Check()
	if err != nil {
		panic(fmt.Errorf("the optimization problem is not well-formed: %v", err))
	}

	// Check Objective
	if !op.Objective.IsLinear() {
		return false
	}

	// Check Constraints
	for _, constraint := range op.Constraints {
		if !constraint.IsLinear() {
			return false
		}
	}

	// All Checks Passed!
	return true
}

/*
LinearInequalityConstraintMatrices
Description:

	Returns the linear INEQUALITY constraint matrices and vectors.
	For all linear inequality constraints, we assemble them into the form:
		Ax <= b
	Where A is the matrix of coefficients, x is the vector of variables, and b is the vector of constants.
	We return A and b.
*/
func (op *OptimizationProblem) LinearInequalityConstraintMatrices() (symbolic.KMatrix, symbolic.KVector) {
	// Setup

	// Collect the Variables of this Problem
	x := op.Variables

	// Iterate through all constraints and collect the linear constraints
	// into a matrix and vector.
	scalar_constraints := make([]symbolic.ScalarConstraint, 0)
	vector_constraints := make([]symbolic.VectorConstraint, 0)
	for _, constraint := range op.Constraints {
		// Skip this constraint if it is not linear
		if !constraint.IsLinear() {
			continue
		}
		// Skip this constraint if it is not an inequality
		if constraint.ConstrSense() == symbolic.SenseEqual {
			continue
		}
		switch c := constraint.(type) {
		case symbolic.ScalarConstraint:
			scalar_constraints = append(scalar_constraints, c)
		case symbolic.VectorConstraint:
			vector_constraints = append(vector_constraints, c)
		}
	}

	// Create the matrix and vector elements from the scalar constraints
	A_components_scalar := make([]mat.VecDense, len(scalar_constraints))
	b_components_scalar := make([]float64, len(scalar_constraints))
	for ii, constraint := range scalar_constraints {
		A_components_scalar[ii], b_components_scalar[ii] = constraint.LinearInequalityConstraintRepresentation(x)
	}

	// Create the matrix and vector elements from the vector constraints
	A_components_vector := make([]mat.Dense, len(vector_constraints))
	b_components_vector := make([]mat.VecDense, len(vector_constraints))
	for ii, constraint := range vector_constraints {
		A_components_vector[ii], b_components_vector[ii] = constraint.LinearInequalityConstraintRepresentation(x)
	}

	// Assemble the matrix and vector components
	var AOut symbolic.Expression
	var bOut symbolic.Expression
	scalar_constraint_matrices_exist := len(A_components_scalar) > 0
	if scalar_constraint_matrices_exist {
		AOut = symbolic.VecDenseToKVector(A_components_scalar[0]).Transpose()
		for ii := 1; ii < len(A_components_scalar); ii++ {
			AOut = symbolic.VStack(
				AOut,
				symbolic.VecDenseToKVector(A_components_scalar[ii]).Transpose(),
			)
		}
		bOut = getKVector.From(b_components_scalar)
	}

	vector_constraint_matrices_exist := len(A_components_vector) > 0
	if vector_constraint_matrices_exist {
		// Create the matrix, if it doesn't already exist
		if !scalar_constraint_matrices_exist {
			AOut = symbolic.DenseToKMatrix(A_components_vector[0])
			bOut = symbolic.VecDenseToKVector(b_components_vector[0])
		} else {
			AOut = symbolic.VStack(
				AOut,
				symbolic.DenseToKMatrix(A_components_vector[0]),
			)
			bOut = symbolic.VStack(
				bOut,
				symbolic.VecDenseToKVector(b_components_vector[0]),
			)
		}
		for ii := 1; ii < len(A_components_vector); ii++ {
			AOut = symbolic.VStack(
				AOut,
				symbolic.DenseToKMatrix(A_components_vector[ii]),
			)
			bOut = symbolic.VStack(
				bOut,
				symbolic.VecDenseToKVector(b_components_vector[ii]),
			)
		}
	}

	return AOut.(symbolic.KMatrix), bOut.(symbolic.KVector)
}

/*
LinearEqualityConstraintMatrices
Description:

	Returns the linear EQUALITY constraint matrices and vectors.
	For all linear equality constraints, we assemble them into the form:
		Cx = d
	Where C is the matrix of coefficients, x is the vector of variables, and d is the vector of constants.
	We return C and d.
*/
func (op *OptimizationProblem) LinearEqualityConstraintMatrices() (symbolic.KMatrix, symbolic.KVector) {
	// Setup

	// Collect the Variables of this Problem
	x := op.Variables

	// Iterate through all constraints and collect the linear constraints
	// into a matrix and vector.
	scalar_constraints := make([]symbolic.ScalarConstraint, 0)
	vector_constraints := make([]symbolic.VectorConstraint, 0)
	for _, constraint := range op.Constraints {
		// Skip this constraint if it is not linear
		if !constraint.IsLinear() {
			continue
		}
		// Skip this constraint if it is not an equality
		if constraint.ConstrSense() != symbolic.SenseEqual {
			continue
		}
		switch c := constraint.(type) {
		case symbolic.ScalarConstraint:
			scalar_constraints = append(scalar_constraints, c)
		case symbolic.VectorConstraint:
			vector_constraints = append(vector_constraints, c)
		}
	}

	// Create the matrix and vector elements from the scalar constraints
	C_components_scalar := make([]mat.VecDense, len(scalar_constraints))
	d_components_scalar := make([]float64, len(scalar_constraints))
	for ii, constraint := range scalar_constraints {
		C_components_scalar[ii], d_components_scalar[ii] = constraint.LinearEqualityConstraintRepresentation(x)
	}

	// Create the matrix and vector elements from the vector constraints
	C_components_vector := make([]mat.Dense, len(vector_constraints))
	d_components_vector := make([]mat.VecDense, len(vector_constraints))
	for ii, constraint := range vector_constraints {
		C_components_vector[ii], d_components_vector[ii] = constraint.LinearEqualityConstraintRepresentation(x)
	}

	// Assemble the matrix and vector components
	var COut symbolic.Expression
	var dOut symbolic.Expression
	scalar_constraint_matrices_exist := len(C_components_scalar) > 0
	if scalar_constraint_matrices_exist {
		COut = symbolic.VecDenseToKVector(C_components_scalar[0]).Transpose()
		for ii := 1; ii < len(C_components_scalar); ii++ {
			COut = symbolic.VStack(
				COut,
				symbolic.VecDenseToKVector(C_components_scalar[ii]).Transpose(),
			)
		}
		dOut = getKVector.From(d_components_scalar)
	}
	vector_constraint_matrices_exist := len(C_components_vector) > 0

	if vector_constraint_matrices_exist {
		// Create the matrix, if it doesn't already exist
		if !scalar_constraint_matrices_exist {
			COut = symbolic.DenseToKMatrix(C_components_vector[0])
			dOut = symbolic.VecDenseToKVector(d_components_vector[0])
		} else {
			COut = symbolic.VStack(
				COut,
				symbolic.DenseToKMatrix(C_components_vector[0]),
			)
			dOut = symbolic.VStack(
				dOut,
				symbolic.VecDenseToKVector(d_components_vector[0]),
			)
		}
		for ii := 1; ii < len(C_components_vector); ii++ {
			COut = symbolic.VStack(
				COut,
				symbolic.DenseToKMatrix(C_components_vector[ii]),
			)
			dOut = symbolic.VStack(
				dOut,
				symbolic.VecDenseToKVector(d_components_vector[ii]),
			)
		}
	}
	return COut.(symbolic.KMatrix), dOut.(symbolic.KVector)
}
