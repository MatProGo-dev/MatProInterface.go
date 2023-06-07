package optim

import (
	"errors"
	"fmt"
	"time"
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
	ShowLog     bool
	TimeLimit   time.Duration
}

// NewModel returns a new model with some default arguments such as not to show
// the log and no time limit.
func NewModel(name string) *Model {
	return &Model{Name: name, ShowLog: false}
}

//// ShowLog instructs the solver to show the log or not.
//func (m *Model) ShowLog(shouldShow bool) {
//	m.ShowLog = shouldShow
//}

// SetTimeLimit sets the solver time limit for the model.
func (m *Model) SetTimeLimit(dur time.Duration) {
	m.TimeLimit = dur
}

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
func (m *Model) AddConstraint(constr Constraint, extras ...interface{}) error {
	// Constants
	nExtraArguments := len(extras)

	// Input Processing
	switch {
	case nExtraArguments > 1:
		// Do nothing, but report an error
		return fmt.Errorf(
			"The optimizer tried to add a constraint using a bad call to AddConstr! Skipping this constraint: %v , because of extra inputs %v",
			constr,
			extras,
		)
	case nExtraArguments == 1:
		e0AsError, tf := extras[0].(error)
		if !tf {
			return fmt.Errorf(
				"The second argument to AddConstraint must be of type error; received type \"%T\".",
				extras[0],
			)
		}

		if e0AsError != nil {
			return fmt.Errorf(
				"There was an error computing constraint %v: %v",
				constr, e0AsError,
			)
		}
	}
	// If no extras are given, then move on to last part.

	// Algorithm
	m.Constraints = append(m.Constraints, constr)
	return nil
}

// SetObjective sets the objective of the model given an expression and
// objective sense.
func (m *Model) SetObjective(e ScalarExpression, sense ObjSense) {
	m.Obj = NewObjective(e, sense)
}

// Optimize optimizes the model using the given solver type and returns the
// solution or an error.
func (m *Model) Optimize(solver Solver) (*Solution, error) {
	// Variables
	var err error

	// Input Processing
	if len(m.Variables) == 0 {
		return nil, errors.New("no variables in model")
	}

	// lbs := make([]float64, len(m.Variables))
	// ubs := make([]float64, len(m.Variables))
	// types := new(bytes.Buffer)
	// for i, v := range m.Variables {
	// 	lbs[i] = v.Lower
	// 	ubs[i] = v.Upper
	// 	types.WriteByte(byte(v.Vtype))
	// }

	solver.ShowLog(m.ShowLog)

	if m.TimeLimit > 0 {
		solver.SetTimeLimit(m.TimeLimit.Seconds())
	}

	solver.AddVariables(m.Variables)

	for _, constr := range m.Constraints {
		solver.AddConstraint(constr)
	}

	//if m.Obj != nil {
	//	logrus.WithField(
	//		"num_vars", m.Obj.NumVars(),
	//	).Info("Number of variables in objective")
	//	err = solver.SetObjective(*m.Obj)
	//	if err != nil {
	//		return nil, fmt.Errorf("There was an error setting the objective: %v", err)
	//	}
	//}

	mipSol, err := solver.Optimize()
	defer solver.DeleteSolver()

	if err != nil {
		return nil, fmt.Errorf("There was an issue while trying to optimize the model: %v", err)
	}

	if mipSol.Status != OptimizationStatus_OPTIMAL {
		errorMessage, err := mipSol.Status.ToMessage()
		if err != nil {
			return nil, fmt.Errorf("There was an issue converting optimization status to a message: %v", err)
		}
		return nil, fmt.Errorf(
			"[Code = %d] %s",
			mipSol.Status,
			errorMessage,
		)
	}

	return &mipSol, nil
}
