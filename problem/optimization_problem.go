package problem

import (
	"fmt"
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
	Objective   symbolic.Expression
}

// NewProblem returns a new model with some default arguments such as not to show
// the log and no time limit.
func New(name string) *OptimizationProblem {
	return &OptimizationProblem{Name: name}
}

/*
AddVariable
Description:

	This method adds an "unbounded" continuous variable to the model.
*/
func (m *OptimizationProblem) AddVariable() symbolic.Variable {
	return m.AddRealVariable()
}

/*
AddRealVariable
Description:

	Adds a Real variable to the model and returns said variable.
*/
func (m *OptimizationProblem) AddRealVariable() symbolic.Variable {
	return m.AddVariableClassic(-optim.INFINITY, optim.INFINITY, symbolic.Continuous)
}

// AddVariable adds a variable of a given variable type to the model given the lower
// and upper value limits. This variable is returned.
func (m *OptimizationProblem) AddVariableClassic(lower, upper float64, vtype symbolic.VarType) symbolic.Variable {
	id := uint64(len(m.Variables))
	newVar := symbolic.Variable{
		ID:    id,
		Lower: lower,
		Upper: upper,
		Type:  vtype,
		Name:  fmt.Sprintf("x_%v", id),
	}
	m.Variables = append(m.Variables, newVar)
	return newVar
}

// AddBinaryVar adds a binary variable to the model and returns said variable.
func (m *OptimizationProblem) AddBinaryVariable() symbolic.Variable {
	return m.AddVariableClassic(0, 1, symbolic.Binary)
}

/*
AddVariableVector
Description:

	Creates a VarVector object using a constructor that assumes you want an "unbounded" vector of real optimization
	variables.
*/
func (m *OptimizationProblem) AddVariableVector(dim int) symbolic.VariableVector {
	// Constants

	// Algorithm
	varSlice := make([]symbolic.Variable, dim)
	for eltIndex := 0; eltIndex < dim; eltIndex++ {
		varSlice[eltIndex] = m.AddVariable()
	}
	return varSlice
}

/*
AddVariableVectorClassic
Description:

	The classic version of AddVariableVector defined in the original goop.
*/
func (m *OptimizationProblem) AddVariableVectorClassic(
	num int, lower, upper float64, vtype symbolic.VarType,
) symbolic.VariableVector {
	stID := uint64(len(m.Variables))
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

	m.Variables = append(m.Variables, vs...)
	return vs
}

// AddBinaryVariableVector adds a vector of binary variables to the model and
// returns the slice.
func (m *OptimizationProblem) AddBinaryVariableVector(num int) symbolic.VariableVector {
	return m.AddVariableVectorClassic(num, 0, 1, symbolic.Binary)
}

// AddVariableMatrix adds a matrix of variables of a given type to the model with
// lower and upper value limits and returns the resulting slice.
func (m *OptimizationProblem) AddVariableMatrix(
	rows, cols int, lower, upper float64, vtype symbolic.VarType,
) symbolic.VariableMatrix {
	// TODO: Add support for adding a variable matrix with a given
	// environment as well as upper and lower bounds.
	return symbolic.NewVariableMatrix(rows, cols)
}

// AddBinaryVariableMatrix adds a matrix of binary variables to the model and returns
// the resulting slice.
func (m *OptimizationProblem) AddBinaryVariableMatrix(rows, cols int) [][]symbolic.Variable {
	return m.AddVariableMatrix(rows, cols, 0, 1, symbolic.Binary)
}

// AddConstr adds the given constraint to the model.
func (m *OptimizationProblem) AddConstraint(constr symbolic.Constraint) error {
	// Constants

	// Input Processing
	//err := constr.Check()
	//if err != nil {
	//	return err
	//}

	// Algorithm
	m.Constraints = append(m.Constraints, constr)
	return nil
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

func (m *OptimizationProblem) SetObjective(e symbolic.Expression, sense ObjSense) error {
	// Input Processing
	se, err := symbolic.ToScalarExpression(e)
	if err != nil {
		return fmt.Errorf("trouble parsing input expression: %v", err)
	}

	// Return
	m.Objective = NewObjective(se, sense)
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
	lhs, err := inputConstraint.Left().ToSymbolic()
	if err != nil {
		return symbolic.ScalarConstraint{}, err
	}

	// Convert RHS to symbolic expression
	rhs, err := inputConstraint.Right().ToSymbolic()
	if err != nil {
		return symbolic.ScalarConstraint{}, err
	}

	// Get Sense
	sense := inputConstraint.ConstrSense().ToSymbolic()

	// Convert
	switch {
	case optim.IsScalarExpression(lhs):
		return symbolic.ScalarConstraint{
			LeftHandSide:  lhs.(symbolic.ScalarExpression),
			RightHandSide: rhs.(symbolic.ScalarExpression),
			Sense:         sense,
		}, nil
	case optim.IsVectorExpression(lhs):
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
ToOptimizationProblem
Description:

	Converts the given input into an optimization problem.
*/
func ToOptimizationProblem(inputModel optim.Model) (*OptimizationProblem, error) {
	// Create a new optimization problem
	newOptimProblem := NewProblem(inputModel.Name)

	// Input Processing
	err := inputModel.Check()
	if err != nil {
		return nil, err
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
	newOptimProblem.Objective, err = inputModel.Obj.ToSymbolic()
	if err != nil {
		return nil, err
	}

	// Done
	return newOptimProblem, nil

}
