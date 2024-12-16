package problem

import (
	"fmt"

	"github.com/MatProGo-dev/MatProInterface.go/mpiErrors"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
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
