package optim

import (
	"fmt"
)

// Model represents the overall constrained linear optimization model to be
// solved. It contains all variables associated with the optimization
// problem, constraints, objective, and parameters. New variables can only be
// created using an instantiated Model.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
type Model struct {
	Name        string
	Variables   []Variable
	Constraints []Constraint
	Obj         *Objective
}

// NewModel returns a new model with the given name.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func NewModel(name string) *Model {
	return &Model{Name: name}
}

// AddVariable adds an "unbounded" continuous variable to the model.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (m *Model) AddVariable() Variable {
	return m.AddRealVariable()
}

// AddRealVariable adds a real-valued variable to the model and returns said variable.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (m *Model) AddRealVariable() Variable {
	return m.AddVariableClassic(-INFINITY, INFINITY, Continuous)
}

// AddVariableClassic adds a variable of a given variable type to the model given the lower
// and upper value limits. This variable is returned.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (m *Model) AddVariableClassic(lower, upper float64, vtype VarType) Variable {
	id := uint64(len(m.Variables))
	newVar := Variable{id, lower, upper, vtype}
	m.Variables = append(m.Variables, newVar)
	return newVar
}

// AddBinaryVariable adds a binary variable to the model and returns said variable.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (m *Model) AddBinaryVariable() Variable {
	return m.AddVariableClassic(0, 1, Binary)
}

// AddVariableVector creates a VarVector of the given dimension containing
// unbounded real optimization variables.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (m *Model) AddVariableVector(dim int) VarVector {
	// Constants

	// Algorithm
	varSlice := make([]Variable, dim)
	for eltIndex := 0; eltIndex < dim; eltIndex++ {
		varSlice[eltIndex] = m.AddVariable()
	}
	return VarVector{varSlice}
}

// AddVariableVectorClassic adds a vector of num variables with the given lower
// bound, upper bound, and variable type to the model.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
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
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (m *Model) AddBinaryVariableVector(num int) VarVector {
	return m.AddVariableVectorClassic(num, 0, 1, Binary)
}

// AddVariableMatrix adds a matrix of variables of a given type to the model with
// lower and upper value limits and returns the resulting slice.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
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
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (m *Model) AddBinaryVariableMatrix(rows, cols int) [][]Variable {
	return m.AddVariableMatrix(rows, cols, 0, 1, Binary)
}

// AddConstraint adds the given constraint to the model.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
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

// SetObjective sets the objective of the model given an expression and
// objective sense.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
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

// Check checks the model for errors, ensuring that at least one variable exists.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (m *Model) Check() error {
	// Constants

	// Verifiy that there is at least one variable in the model.
	if len(m.Variables) == 0 {
		return fmt.Errorf("the model has no variables!")
	}

	// It's okay if there are no constraints.

	// It's okay if there is not an objective.

	// All Checks have passed.
	return nil
}
