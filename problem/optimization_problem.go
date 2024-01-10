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
func NewProblem(name string) *OptimizationProblem {
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
	newVar := symbolic.Variable{id, lower, upper, vtype}
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
	return symbolic.VariableVector{Elements: varSlice}
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
		vs[i] = symbolic.Variable{stID + uint64(i), lower, upper, vtype}
	}

	m.Variables = append(m.Variables, vs...)
	return symbolic.VariableVector{Elements: vs}
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
) [][]symbolic.Variable {
	vs := make([][]symbolic.Variable, rows)
	for i := range vs {
		tempVV := m.AddVariableVectorClassic(cols, lower, upper, vtype)
		vs[i] = tempVV.Elements
	}

	return vs
}

// AddBinaryVariableMatrix adds a matrix of binary variables to the model and returns
// the resulting slice.
func (m *OptimizationProblem) AddBinaryVariableMatrix(rows, cols int) [][]symbolic.Variable {
	return m.AddVariableMatrix(rows, cols, 0, 1, symbolic.Binary)
}

// AddConstr adds the given constraint to the model.
func (m *OptimizationProblem) AddConstraint(constr symbolic.Constraint, errors ...error) error {
	// Constants

	// Input Processing
	err := symbolic.CheckErrors(errors)
	if err != nil {
		return err
	}

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

//// Optimize optimizes the model using the given solver type and returns the
//// solution or an error.
//func (m *OptimizationProblem) Optimize(solver Solver) (*Solution, error) {
//	// Variables
//	var err error
//
//	// Input Processing
//	if len(m.Variables) == 0 {
//		return nil, errors.New("no variables in model")
//	}
//
//	solver.ShowLog(m.ShowLog)
//
//	if m.TimeLimit > 0 {
//		solver.SetTimeLimit(m.TimeLimit.Seconds())
//	}
//
//	solver.AddVariables(m.Variables)
//
//	for _, constr := range m.Constraints {
//		solver.AddConstraint(constr)
//	}
//
//	mipSol, err := solver.Optimize()
//	defer solver.DeleteSolver()
//
//	if err != nil {
//		return nil, fmt.Errorf("There was an issue while trying to optimize the model: %v", err)
//	}
//
//	if mipSol.Status != OptimizationStatus_OPTIMAL {
//		errorMessage, err := mipSol.Status.ToMessage()
//		if err != nil {
//			return nil, fmt.Errorf("There was an issue converting optimization status to a message: %v", err)
//		}
//		return nil, fmt.Errorf(
//			"[Code = %d] %s",
//			mipSol.Status,
//			errorMessage,
//		)
//	}
//
//	return &mipSol, nil
//}
