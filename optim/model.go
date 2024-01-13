package optim

import (
	"fmt"
)

// Model represents the overall constrained linear optimization model to be
// solved. Model contains all the variables associated with the optimization
// problem, constraints, objective, and parameters. New variables can only be
// created using an instantiated Model.
type Model struct {
	Name        string
	Variables   []Variable
	Constraints []Constraint
	Obj         *Objective
}

// NewModel returns a new model with some default arguments such as not to show
// the log.
func NewModel(name string) *Model {
	return &Model{Name: name}
}

//// ShowLog instructs the solver to show the log or not.
//func (m *Model) ShowLog(shouldShow bool) {
//	m.ShowLog = shouldShow
//}

/*
AddVariable
Description:

	This method adds an "unbounded" continuous variable to the model.
*/
func (m *Model) AddVariable() Variable {
	return m.AddRealVariable()
}

/*
AddRealVariable
Description:

	Adds a Real variable to the model and returns said variable.
*/
func (m *Model) AddRealVariable() Variable {
	return m.AddVariableClassic(-INFINITY, INFINITY, Continuous)
}

// AddVariable adds a variable of a given variable type to the model given the lower
// and upper value limits. This variable is returned.
func (m *Model) AddVariableClassic(lower, upper float64, vtype VarType) Variable {
	id := uint64(len(m.Variables))
	newVar := Variable{id, lower, upper, vtype}
	m.Variables = append(m.Variables, newVar)
	return newVar
}

// AddBinaryVar adds a binary variable to the model and returns said variable.
func (m *Model) AddBinaryVariable() Variable {
	return m.AddVariableClassic(0, 1, Binary)
}

/*
AddVariableVector
Description:

	Creates a VarVector object using a constructor that assumes you want an "unbounded" vector of real optimization
	variables.
*/
func (m *Model) AddVariableVector(dim int) VarVector {
	// Constants

	// Algorithm
	varSlice := make([]Variable, dim)
	for eltIndex := 0; eltIndex < dim; eltIndex++ {
		varSlice[eltIndex] = m.AddVariable()
	}
	return VarVector{varSlice}
}

/*
AddVariableVectorClassic
Description:

	The classic version of AddVariableVector defined in the original goop.
*/
func (m *Model) AddVariableVectorClassic(
	num int, lower, upper float64, vtype VarType,
) VarVector {
	stID := uint64(len(m.Variables))
	vs := make([]Variable, num)
	for i := range vs {
		vs[i] = Variable{stID + uint64(i), lower, upper, vtype}
	}

	m.Variables = append(m.Variables, vs...)
	return VarVector{vs}
}

// AddBinaryVariableVector adds a vector of binary variables to the model and
// returns the slice.
func (m *Model) AddBinaryVariableVector(num int) VarVector {
	return m.AddVariableVectorClassic(num, 0, 1, Binary)
}

// AddVariableMatrix adds a matrix of variables of a given type to the model with
// lower and upper value limits and returns the resulting slice.
func (m *Model) AddVariableMatrix(
	rows, cols int, lower, upper float64, vtype VarType,
) [][]Variable {
	vs := make([][]Variable, rows)
	for i := range vs {
		tempVV := m.AddVariableVectorClassic(cols, lower, upper, vtype)
		vs[i] = tempVV.Elements
	}

	return vs
}

// AddBinaryVariableMatrix adds a matrix of binary variables to the model and returns
// the resulting slice.
func (m *Model) AddBinaryVariableMatrix(rows, cols int) [][]Variable {
	return m.AddVariableMatrix(rows, cols, 0, 1, Binary)
}

// AddConstr adds the given constraint to the model.
func (m *Model) AddConstraint(constr Constraint, errors ...error) error {
	// Constants

	// Input Processing
	err := CheckErrors(errors)
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

func (m *Model) SetObjective(e Expression, sense ObjSense) error {
	// Input Processing
	se, err := ToScalarExpression(e)
	if err != nil {
		return fmt.Errorf("trouble parsing input expression: %v", err)
	}

	// Return
	m.Obj = NewObjective(se, sense)
	return nil
}
